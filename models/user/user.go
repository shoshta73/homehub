package user

import (
	"crypto/rand"
	"crypto/sha512"
	"database/sql"
	"fmt"
	"strconv"

	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"xorm.io/xorm"

	"github.com/charmbracelet/log"
	"github.com/shoshta73/homehub/constants"
	"github.com/shoshta73/homehub/server/metadata"

	_ "github.com/mattn/go-sqlite3"
)

const secretsDir = "secrets"

const tokenFile = "token.txt"

var orm *xorm.Engine

const (
	adminPermission uint8 = 1 << iota
	userPermission
)

type User struct {
	Id          int64
	Username    string `xorm:"unique not null"`
	Name        string
	Email       string `xorm:"unique not null" valid:"Email"`
	Password    string
	Permissions uint8
	Avatar      string `xorm:"varchar(2048) not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserClaims struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Permissions uint8  `json:"permissions"`
	jwt.RegisteredClaims
}

func (uc UserClaims) GenerateToken() (string, error) {
	data, err := os.ReadFile(filepath.Join(secretsDir, tokenFile))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString([]byte(data))
}

func (u User) TableName() string {
	return "users"
}

func (u User) GetAvatarURL() string {
	return filepath.Join("avatars", u.Username+".png")
}

func (u User) GetAvatar() string {
	return filepath.Join(constants.DATA_DIR, "identicons", u.Username+".png")
}

func (u User) GetClaims() *UserClaims {
	tn := time.Now()
	return &UserClaims{
		Id:          u.Id,
		Username:    u.Username,
		Permissions: u.Permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "homehub",
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 24 * 3),
			},
			IssuedAt: &jwt.NumericDate{
				Time: tn,
			},
			NotBefore: &jwt.NumericDate{
				Time: tn,
			},
		},
	}
}

func (u *User) BeforeInsert() {
	tn := time.Now()

	u.CreatedAt = tn
	u.UpdatedAt = tn

	usernameHash := sha512.Sum512([]byte(u.Username))
	emailHash := sha512.Sum512([]byte(u.Email))
	var nameHash [64]byte
	if u.Name != "" {
		nameHash = sha512.Sum512([]byte(u.Name))
	} else {
		nameHash = sha512.Sum512([]byte(u.Username))
	}
	createdAtHash := sha512.Sum512([]byte(strconv.FormatInt(u.CreatedAt.Unix(), 10)))

	meta := metadata.GetMetadata()

	if meta.HasAdmin {
		u.Permissions = userPermission
	} else {
		log.Info("Creating first user")
		u.Permissions = adminPermission | userPermission
		meta.UpdateHasAdmin(true)
		go meta.Write()
	}

	combinedHash := make([]byte, 0, 64*4)
	combinedHash = append(combinedHash, usernameHash[:]...)
	combinedHash = append(combinedHash, emailHash[:]...)
	combinedHash = append(combinedHash, nameHash[:]...)
	combinedHash = append(combinedHash, createdAtHash[:]...)

	u.Avatar = string(combinedHash)
}

func (u *User) encryptPassword(pass string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func IsExistingByEmail(email string) (bool, error) {
	return orm.Get(&User{Email: email})
}

func IsExistingByUsername(username string) (bool, error) {
	return orm.Get(&User{Username: username})
}

func init() {
	_, err := os.Stat(secretsDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(secretsDir, 0755)
		} else {
			log.Fatal(err)
		}
	}

	created := false
	_, err = os.Stat(filepath.Join(constants.DATA_DIR, constants.DATABASE_FILE))
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(filepath.Join(constants.DATA_DIR, constants.DATABASE_FILE))
			if err != nil {
				log.Fatal(err)
			}
			f.Close()
			created = true
		} else {
			log.Fatal(err)
		}
	}

	if created {
		db, err := sql.Open("sqlite3", filepath.Join(constants.DATA_DIR, constants.DATABASE_FILE))
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("PRAGMA foreign_keys = ON;")
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("PRAGMA journal_mode=WAL;")
		if err != nil {
			log.Fatal(err)
		}
	}

	engine, err := xorm.NewEngine("sqlite3", filepath.Join(constants.DATA_DIR, constants.DATABASE_FILE))
	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync(new(User))
	if err != nil {
		log.Fatal(err)
	}

	orm = engine

	created = false
	_, err = os.Stat(filepath.Join(secretsDir, tokenFile))
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(filepath.Join(secretsDir, tokenFile))
			if err != nil {
				log.Fatal(err)
			}
			f.Close()
			created = true
		} else {
			log.Fatal(err)
		}
	}

	if created {
		token := make([]byte, 64)
		rand.Read(token)
		err = os.WriteFile(filepath.Join(secretsDir, tokenFile), token, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CreateUser(username, name, email, pass string) (*User, error) {
	var user User
	if name == "" {
		user = User{
			Username: username,
			Name:     username,
			Email:    email,
		}
	} else {
		user = User{
			Username: username,
			Name:     name,
			Email:    email,
		}
	}

	err := user.encryptPassword(pass)
	if err != nil {
		return nil, err
	}

	user.BeforeInsert()
	_, err = orm.Insert(&user)
	if err != nil {
		return nil, err
	}

	log.Info("User created", "user", map[string]string{"Id": strconv.Itoa(int(user.Id)), "username": username, "name": name, "email": email})

	return &user, nil
}

func VerifyUser(username, email, pass string) bool {
	user := &User{}

	if username == "" && email == "" || pass == "" {
		return false
	}

	if email != "" {
		_, err := orm.Where("email = ?", email).Get(user)
		if err != nil {
			return false
		}
	}

	if username != "" {
		_, err := orm.Where("username = ?", username).Get(user)
		if err != nil {
			return false
		}
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)) == nil
}

func VerifyUserByEmail(email, pass string) bool {
	user := &User{}

	_, err := orm.Where("email = ?", email).Get(user)
	if err != nil {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)) == nil
}

func VerifyUserByUsername(username, pass string) bool {
	user := &User{}

	_, err := orm.Where("username = ?", username).Get(user)
	if err != nil {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)) == nil
}

func GetUserById(id int64) (*User, error) {
	user := &User{}

	_, err := orm.Where("id = ?", id).Get(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByUsername(username string) (*User, error) {
	user := &User{}

	_, err := orm.Where("username = ?", username).Get(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(email string) (*User, error) {
	user := &User{}

	_, err := orm.Where("email = ?", email).Get(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByToken(token string) (*User, error) {
	user := &User{}

	log.Info("Getting user by token", "token", token)

	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return os.ReadFile(filepath.Join(secretsDir, tokenFile))
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	log.Info("Getting user by id", "id", claims.Id)

	_, err = orm.Where("id = ?", claims.Id).Get(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

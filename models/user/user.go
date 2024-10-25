package user

import (
	"crypto/sha512"
	"database/sql"
	"strconv"

	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/bcrypt"
	"xorm.io/xorm"

	"github.com/shoshta73/homehub/log"

	_ "github.com/mattn/go-sqlite3"
)

const dataDir = "data"
const databaseFile = "homehub.db"

var orm *xorm.Engine

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

func (User) TableName() string {
	return "users"
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
	_, err := os.Stat(dataDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dataDir, 0755)
		} else {
			log.Fatal(err)
		}
	}

	created := false
	_, err = os.Stat(filepath.Join(dataDir, databaseFile))
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(filepath.Join(dataDir, databaseFile))
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
		db, err := sql.Open("sqlite3", filepath.Join(dataDir, databaseFile))
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

	engine, err := xorm.NewEngine("sqlite3", filepath.Join(dataDir, databaseFile))
	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync(new(User))
	if err != nil {
		log.Fatal(err)
	}

	orm = engine
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

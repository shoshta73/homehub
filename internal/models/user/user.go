package user

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/shoshta73/homehub/internal/storage/database"
	"golang.org/x/crypto/bcrypt"
)

var logger = log.New(os.Stderr)

func init() {
	logger.SetPrefix("user")

	logger.Info("Syncing user mode")
	err := database.GetEngine().Sync(&User{})
	if err != nil {
		logger.Fatal("Failed to sync user model", err)
	}
	logger.Info("User model synced")
}

type User struct {
	ID        string `xorm:"pk 'id'"`
	Username  string `xorm:"unique"`
	Email     string `xorm:"unique"`
	Password  string
	Name      string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) VerifyPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)) == nil
}

func CreateUser(username, email, password string, optionals map[string]string) *User {
	logger.Info("Creating user")
	user := &User{}

	tn := time.Now()

	user.ID = uuid.New().String()

	user.Username = username
	user.Email = email
	log.Info("Encrypting password")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Failed to hash password", err)
		return nil
	}
	user.Password = string(hash)

	name, e := optionals["name"]
	if !e {
		logger.Info("Name not found, using username")
		user.Name = username
	} else {
		logger.Info("Name found", "name", name)
		user.Name = name
	}

	user.CreatedAt = tn
	user.UpdatedAt = tn

	logger.Info("User created")
	return user
}

func InsertUser(user *User) error {
	logger.Info("Inserting user")
	_, err := database.GetEngine().Insert(user)
	if err != nil {
		return err
	}
	logger.Info("User inserted")
	return nil
}

func UsernameExists(username string) (bool, error) {
	return database.GetEngine().Get(&User{Username: username})
}

func EmailExists(email string) (bool, error) {
	return database.GetEngine().Get(&User{Email: email})
}

func GetUserByEmail(email string) (*User, error) {
	user := &User{}
	_, err := database.GetEngine().Where("email = ?", email).Get(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

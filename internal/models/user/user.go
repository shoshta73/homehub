package user

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/shoshta73/homehub/internal/storage/database"
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

func CreateUser(username, email, password string, optionals map[string]string) *User {
	user := &User{}

	tn := time.Now()

	user.ID = uuid.New().String()

	user.Username = username
	user.Email = email
	user.Password = password

	name, e := optionals["name"]
	if !e {
		user.Name = username
	} else {
		user.Name = name
	}

	user.CreatedAt = tn
	user.UpdatedAt = tn

	return user
}

func InsertUser(user *User) error {
	_, err := database.GetEngine().Insert(user)
	if err != nil {
		logger.Error("Failed to insert user", err)
		return err
	}
	return nil
}

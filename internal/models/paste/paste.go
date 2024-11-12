package paste

import (
	"errors"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"

	"github.com/shoshta73/homehub/internal/storage/database"
)

var logger = log.New(os.Stderr)

func init() {
	logger.SetPrefix("paste")

	logger.Info("Syncing paste model")
	err := database.GetEngine().Sync(&Paste{})
	if err != nil {
		logger.Fatal("Failed to sync paste model", err)
	}
	logger.Info("Paste model synced")
}

type Paste struct {
	ID         string `xorm:"pk 'id'"`
	OwnerID    string `xorm:"'owner_id'"`
	Title      string
	Content    string
	Length     int
	Compressed bool
	SharedWith map[string]uint8 `xorm:"json 'shared_with'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (paste *Paste) SetOwnerId(id string) {
	paste.OwnerID = id
}

func HasTitle(userId, title string) bool {
	var pastes []Paste
	var pastesMap map[string]string

	pastes = []Paste{}
	pastesMap = map[string]string{}

	err := database.GetEngine().Where("owner_id = ?", userId).Find(&pastes)
	if err != nil {
		logger.Error("Error getting pastes", err)
		return false
	}

	for _, paste := range pastes {
		pastesMap[paste.Title] = paste.ID
	}

	_, exists := pastesMap[title]

	return exists
}

func Create(title, content string) (*Paste, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	if content == "" {
		return nil, errors.New("empty pastes are not allowed")
	}

	paste := &Paste{}

	tn := time.Now()

	paste.ID = uuid.New().String()
	paste.CreatedAt = tn
	paste.UpdatedAt = tn

	paste.Title = title
	paste.Content = content
	paste.Length = len(content)
	paste.SharedWith = map[string]uint8{}
	paste.Compressed = false

	return paste, nil
}

func Insert(paste *Paste) error {
	logger.Info("Inserting paste")
	_, err := database.GetEngine().Insert(paste)
	if err != nil {
		return err
	}
	logger.Info("Inserted paste")
	return nil
}

func CreatedCount(userId string) (int64, error) {
	logger.Info("Getting created pastes for user", "id", userId)

	return database.GetEngine().Where("owner_id = ?", userId).Count(&Paste{})
}

func GetCreatedPastes(userId string) ([]Paste, error) {
	logger.Info("Getting pastes created by user", "id", userId)

	var pastes []Paste
	pastes = []Paste{}

	err := database.GetEngine().Where("owner_id = ?", userId).Find(&pastes)
	if err != nil {
		logger.Error("Error getting pastes", err)
		return nil, err
	}

	return pastes, nil
}

func GetById(id string) (*Paste, error) {
	paste := &Paste{}

	_, err := database.GetEngine().Where("id = ?", id).Get(paste)
	if err != nil {
		return nil, err
	}

	return paste, nil
}

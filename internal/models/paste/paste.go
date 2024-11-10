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
	logger.SetPrefix("user")

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

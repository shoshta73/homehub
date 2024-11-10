package paste

import (
	"github.com/charmbracelet/log"
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

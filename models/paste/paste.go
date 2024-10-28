package paste

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/shoshta73/homehub/constants"
	"github.com/shoshta73/homehub/models/stats"
	"xorm.io/xorm"
)

var orm *xorm.Engine = nil

func init() {
	created := false
	_, err := os.Stat(filepath.Join(constants.DATA_DIR, constants.DATABASE_FILE))
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

	err = engine.Sync(new(Paste))
	if err != nil {
		log.Fatal(err)
	}

	orm = engine
}

type Paste struct {
	Id        int64 `xorm:"unique pk"`
	Title     string
	Content   string
	CreatedAt int64
	UpdatedAt int64
	OwnerId   int64
}

func (p Paste) TableName() string {
	return "pastes"
}

func CreatePaste(title, content string, ownerId int64) (*Paste, error) {
	paste := Paste{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		OwnerId:   ownerId,
	}

	_, err := orm.Insert(&paste)
	if err != nil {
		return nil, err
	}

	go stats.IncrementPasteCreated(ownerId)

	return &paste, nil
}

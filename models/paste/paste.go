package paste

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/shoshta73/homehub/constants"
	"github.com/shoshta73/homehub/log"
	"github.com/shoshta73/homehub/models/stats"
	"xorm.io/xorm"
)

var pastedir string = filepath.Join(constants.DATA_DIR, "pastes")
var orm *xorm.Engine = nil

func init() {
	log.Info("Initializing pastes")
	created := false
	_, err := os.Stat(filepath.Join(constants.DATA_DIR, constants.DATABASE_FILE))
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("Database file not found, creating")
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
		log.Info("Database file created, initializing")
		db, err := sql.Open("sqlite3", filepath.Join(constants.DATA_DIR, constants.DATABASE_FILE))
		if err != nil {
			log.Fatal(err)
		}

		log.Info("Setting foreign keys")
		_, err = db.Exec("PRAGMA foreign_keys = ON;")
		if err != nil {
			log.Fatal(err)
		}

		log.Info("Setting journal mode")
		_, err = db.Exec("PRAGMA journal_mode=WAL;")
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Info("Opening database")
	engine, err := xorm.NewEngine("sqlite3", filepath.Join(constants.DATA_DIR, constants.DATABASE_FILE))
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Syncing paste model to database")
	err = engine.Sync(new(Paste))
	if err != nil {
		log.Fatal(err)
	}

	orm = engine

	log.Info("Checking paste directory")
	_, err = os.Stat(pastedir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("Creating paste directory")
			os.Mkdir(pastedir, 0755)
			log.Info("Paste directory created")
		} else {
			log.Fatal(err)
		}
	}
}

type Paste struct {
	Id         int64 `xorm:"unique pk autoincr"`
	Title      string
	Content    string
	CreatedAt  int64
	UpdatedAt  int64
	OwnerId    int64
	Compressed bool
}

func (p Paste) TableName() string {
	return "pastes"
}

func CreatePaste(title, content string, ownerId int64) (*Paste, error) {
	var cnt string
	comp := false
	if len(content) > 512 {
		comp = true
		log.Info("Content is too long, compressing")

		var buf bytes.Buffer
		hash := sha256.Sum256([]byte(fmt.Sprintf("%s %s:%s", time.Now().String(), title, content)))
		str := hex.EncodeToString(hash[:])
		filePath := filepath.Join(pastedir, str)

		log.Info("Compressing to", filePath)
		zipWr := zip.NewWriter(&buf)
		wr, err := zipWr.Create(filePath)
		if err != nil {
			return nil, err
		}

		log.Info("Writing to zip")
		_, err = wr.Write([]byte(content))
		if err != nil {
			return nil, err
		}

		log.Info("Closing zip")
		err = zipWr.Close()
		if err != nil {
			return nil, err
		}

		log.Info("Closing file")
		err = os.WriteFile(filePath, buf.Bytes(), 0644)
		if err != nil {
			return nil, err
		}

		cnt = str
	} else {
		cnt = content
	}

	paste := Paste{
		Title:      title,
		Content:    cnt,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
		OwnerId:    ownerId,
		Compressed: comp,
	}

	log.Info("Inserting paste")
	_, err := orm.Insert(&paste)
	if err != nil {
		return nil, err
	}

	log.Info("Incrementing paste created")
	go stats.IncrementPasteCreated(ownerId)

	return &paste, nil
}

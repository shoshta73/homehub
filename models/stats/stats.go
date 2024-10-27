package stats

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/shoshta73/homehub/constants"
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

	err = engine.Sync(new(PastebinStats))
	if err != nil {
		log.Fatal(err)
	}

	orm = engine
}

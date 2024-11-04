package database

import (
	"database/sql"
	"os"

	"github.com/charmbracelet/log"
	"github.com/shoshta73/homehub/internal/fs"
	"xorm.io/xorm"

	_ "github.com/mattn/go-sqlite3"
)

var (
	logger = log.New(os.Stderr)
	engine *xorm.Engine
)

func init() {

	logger.SetPrefix("db")

	dbpath := fs.DataDir + "/homehub.db"

	logger.Info("Initializing homehub database...")
	exists := fs.FileExists(dbpath)
	if !exists {
		logger.Warn("Database does not exist, creating...")

		file, err := os.Create(dbpath)
		if err != nil {
			logger.Fatal("Failed to create database", err)
		}
		defer file.Close()
		logger.Info("Database created")

		db, err := sql.Open("sqlite3", dbpath)
		if err != nil {
			logger.Fatal("Failed to open database", err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			logger.Fatal("Failed to ping database", err)
		}
		logger.Info("Database opened")

		_, err = db.Exec("PRAGMA journal_mode = WAL;")
		if err != nil {
			logger.Fatal("Failed to set journal mode", err)
		}
		logger.Info("Journal mode set")
	}

	orm, err := xorm.NewEngine("sqlite3", dbpath)
	if err != nil {
		logger.Fatal("Failed to create database", err)
	}
	engine = orm
}

func GetEngine() *xorm.Engine {
	return engine
}

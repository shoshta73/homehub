package fs

import (
	"os"

	"github.com/charmbracelet/log"
)

var logger = log.New(os.Stderr)

var DataDir = "data"
var SecretsDir = "secrets"

func init() {
	logger.SetPrefix("fs")

	logger.Info("Initializing homehub filesystem...")
	logger.Info("Checking for data directory...")
	_, err := os.Stat(DataDir)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Warn("Data directory does not exist, creating...")

			err := os.Mkdir(DataDir, 0755)
			if err != nil {
				logger.Fatal("Failed to create data directory", err)
			}
			logger.Info("Data directory created")
		} else {
			logger.Fatal("Failed to check for data directory", err)
		}
	}

	logger.Info("Checking for secrets directory...")
	_, err = os.Stat(SecretsDir)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Warn("Secrets directory does not exist, creating...")

			err := CreateDir(SecretsDir)
			if err != nil {
				logger.Fatal("Failed to create secrets directory", err)
			}
			logger.Info("Secrets directory created")
		} else {
			logger.Fatal("Failed to check for secrets directory", err)
		}
	}
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			logger.Error("Failed to check for file", err)
			return false
		}
	}
	return !info.IsDir()
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			logger.Error("Failed to check for directory", err)
			return false
		}
	}
	return info.IsDir()
}

func CreateFile(path string) error {
	logger.Info("Creating file", "path", path)
	file, err := os.Create(path)
	if err != nil {
		logger.Error("Failed to create file", err)
		return err
	}
	defer file.Close()
	return nil
}

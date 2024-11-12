package fs

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

var logger = log.New(os.Stderr)

var DataDir = "data"
var SecretsDir = "secrets"
var PastesDir = filepath.Join(DataDir, "pastes")
var ConfigDir = "config"

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

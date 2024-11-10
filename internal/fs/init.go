//go:build !dev
// +build !dev

package fs

import "os"

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

			err := os.Mkdir(SecretsDir, 0755)
			if err != nil {
				logger.Fatal("Failed to create secrets directory", err)
			}
			logger.Info("Secrets directory created")
		} else {
			logger.Fatal("Failed to check for secrets directory", err)
		}
	}

	logger.Info("Checking for pastes directory")
	_, err = os.Stat(PastesDir)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Warn("Pastes directory does not exist, creating...")
			err := os.Mkdir(PastesDir, 0755)
			if err != nil {
				logger.Fatal("Failed to create pastes directory", err)
			}
			logger.Info("Pastes directory created")
		} else {
			logger.Fatal("Failed to check for pastes directory", err)
		}
	}
}

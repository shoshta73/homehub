//go:build dev
// +build dev

package fs

import "os"

func init() {
	logger.SetPrefix("fs")

	logger.Info("Changing directory to bin")
	err := os.Chdir("bin")
	if err != nil {
		logger.Fatal("Failed to change directory", err)
	}
	logger.Info("Directory changed")

	logger.Info("Initializing homehub filesystem...")
	logger.Info("Checking for data directory...")
	_, err = os.Stat(DataDir)
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
}

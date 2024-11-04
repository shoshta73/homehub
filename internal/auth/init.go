package auth

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"

	"github.com/charmbracelet/log"
	"github.com/shoshta73/homehub/internal/fs"
)

var logger = log.New(os.Stderr)

var jwtKey []byte

func init() {
	logger.SetPrefix("auth")

	logger.Info("Checking for JWT secret...")

	fp := fs.SecretsDir + "/jwt.key"

	exists := fs.FileExists(fp)
	if !exists {
		logger.Warn("JWT secret does not exist, creating...")

		file, err := os.Create(fp)
		if err != nil {
			logger.Fatal("Failed to create JWT secret", err)
		}
		logger.Info("JWT secret created")

		logger.Info("Generating JWT secret...")

		bytes := make([]byte, 32)

		_, err = rand.Read(bytes)
		if err != nil {
			logger.Fatal("Failed to generate JWT secret", err)
		}

		key := base64.StdEncoding.EncodeToString(bytes)
		logger.Info("JWT secret generated")

		jwtKey = []byte(key)

		defer file.Close()

		_, err = file.WriteString(key)
		if err != nil {
			logger.Fatal("Failed to write JWT secret to file", err)
		}
		logger.Info("JWT secret written to file")

		return
	}

	logger.Info("JWT secret exists")

	file, err := os.Open(fp)
	if err != nil {
		logger.Fatal("Failed to open JWT secret file", err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		logger.Fatal("Failed to read JWT secret from file", err)
	}

	jwtKey = bytes
}

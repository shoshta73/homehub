package auth

import (
	"os"

	"github.com/charmbracelet/log"
)

var logger = log.New(os.Stderr)

func init() {
	logger.SetPrefix("auth")
}

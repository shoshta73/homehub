package pastebin

import (
	"github.com/charmbracelet/log"
	"os"
)

var logger = log.New(os.Stderr)

func init() {
	logger.SetPrefix("pastebin")
}

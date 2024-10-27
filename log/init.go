package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const logDir = "logs"
const globalLogFile = "log.log"

var globalLogger *multiLogger

func init() {
	_, err := os.Stat(logDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(logDir, 0755)
		} else {
			panic(err)
		}
	}

	lfp := filepath.Join(logDir, globalLogFile)
	var logFile *os.File
	_, err = os.Stat(lfp)
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(lfp)
			if err != nil {
				panic(err)
			}
			logFile = f
		} else {
			panic(err)
		}
	} else {
		de, err := os.ReadDir(logDir)
		if err != nil {
			panic(err)
		}

		numLogFiles := 0

		for _, d := range de {
			if d.IsDir() {
				continue
			}

			if strings.HasPrefix(d.Name(), "log") {
				numLogFiles++
			}
		}

		os.Rename(lfp, filepath.Join(logDir, fmt.Sprintf("log_%d.log", numLogFiles)))
		f, err := os.Create(lfp)
		if err != nil {
			panic(err)
		}
		logFile = f
	}

	globalLogger = newMultiLogger([]io.Writer{os.Stderr, logFile})

	Info("Logger has been initialized")
}

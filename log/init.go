package log

import (
	"io"
	"os"
	"path/filepath"
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

	var logFile *os.File
	os.Stat(filepath.Join(logDir, globalLogFile))
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(filepath.Join(logDir, globalLogFile))
			if err != nil {
				panic(err)
			}
			logFile = f
		} else {
			panic(err)
		}
	}

	globalLogger = newMultiLogger([]io.Writer{os.Stderr, logFile})

	Info("Logger has been initialized")
}

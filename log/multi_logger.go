package log

import (
	"fmt"
	"io"
	"math"
	"os"

	charmLogger "github.com/charmbracelet/log"
)

const (
	noLevel    charmLogger.Level = math.MaxInt32
	timeFormat                   = "06-01-02 15:04:05.000000"
)

type multiLogger struct {
	loggers []*charmLogger.Logger
	writers []io.Writer
}

func newMultiLogger(writers []io.Writer) *multiLogger {
	ml := new(multiLogger)
	ml.writers = make([]io.Writer, len(writers))
	ml.loggers = make([]*charmLogger.Logger, len(writers))

	for i, w := range writers {
		ml.writers[i] = w
		ml.loggers[i] = charmLogger.New(w)
		ml.loggers[i].SetReportTimestamp(true)
		ml.loggers[i].SetTimeFormat(timeFormat)
	}

	return ml
}

func Debug(msg interface{}, keyvals ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(charmLogger.DebugLevel, msg, keyvals...)
	}
}

func Info(msg interface{}, keyvals ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(charmLogger.InfoLevel, msg, keyvals...)
	}
}

func Warn(msg interface{}, keyvals ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(charmLogger.WarnLevel, msg, keyvals...)
	}
}

func Error(msg interface{}, keyvals ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(charmLogger.ErrorLevel, msg, keyvals...)
	}
}

func Fatal(msg interface{}, keyvals ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(charmLogger.FatalLevel, msg, keyvals...)
	}
	os.Exit(1)
}

func Print(msg interface{}, keyvals ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(noLevel, msg, keyvals...)
	}
}

func Debugf(format string, args ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(charmLogger.DebugLevel, fmt.Sprintf(format, args...))
	}
}

func Infof(format string, args ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(charmLogger.InfoLevel, fmt.Sprintf(format, args...))
	}
}

func Warnf(format string, args ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(charmLogger.WarnLevel, fmt.Sprintf(format, args...))
	}
}

func Errorf(format string, args ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(charmLogger.ErrorLevel, fmt.Sprintf(format, args...))
	}
}

func Fatalf(format string, args ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(charmLogger.FatalLevel, fmt.Sprintf(format, args...))
	}
	os.Exit(1)
}

func Printf(format string, args ...interface{}) {
	for _, l := range globalLogger.loggers {
		l.Log(noLevel, fmt.Sprintf(format, args...))
	}
}

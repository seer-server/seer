package log

import (
	"log"
	"os"
	"strings"
)

type LogLevel int

func (l LogLevel) String() string {
	switch l {
	case LogError:
		return "error"
	case LogInfo:
		return "info"
	case LogVerbose:
		return "verbose"
	case LogDebug:
		return "debug"
	default:
		return "INVALID"
	}
}

const (
	stdoutFilename = ":stdout:"
	stderrFilename = ":stderr:"
)

const (
	LogError LogLevel = iota
	LogInfo
	LogVerbose
	LogDebug
)

type Logger struct {
	*log.Logger
	maxLevel LogLevel
}

var (
	logMap   = make(map[string]*Logger)
	logFlags = log.Ldate | log.Ltime | log.Lshortfile
)

func Make(name string, filename string, maxLevel LogLevel) *Logger {
	var (
		file *os.File
		err  error
	)
	switch filename {
	case stdoutFilename:
		file = os.Stdout
	case stderrFilename:
		file = os.Stderr
	default:
		file, err = os.Open(filename)
		if err != nil {
			file = os.Stdout
		}
	}

	if logger, ok := logMap[name]; ok {
		return logger
	} else {
		prefix := []string{"[", strings.Title(name), "] "}
		logger := &Logger{log.New(file, strings.Join(prefix, ""), logFlags), maxLevel}
		logMap[name] = logger

		return logger
	}
}

func Get(name string) (*Logger, bool) {
	l, ok := logMap[name]

	return l, ok
}

func (l *Logger) Fatal(err error, exitCode int) {
	l.Log(LogError, err.Error())
	os.Exit(exitCode)
}

func (l *Logger) Log(level LogLevel, msg string) {
	if level <= l.maxLevel {
		l.Logger.Printf("%10s   %s", level, msg)
	}
}

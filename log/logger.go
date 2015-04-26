package log

import (
	"log"
	"os"
	"strings"
)

var (
	logMap   = make(map[string]*log.Logger)
	logFlags = log.Ldate | log.Ltime | log.Lshortfile
)

func MakeLogFile(name, fname string) (*log.Logger, error) {
	if logger, ok := logMap[name]; ok {
		return logger, nil
	} else {
		file, err := os.Open(fname)
		if err != nil {
			return nil, err
		}
		logger := log.New(file, strings.Title(name), logFlags)
		logMap[name] = logger

		return logger, nil
	}
}

func MakeLogConsole(name string, file *os.File) *log.Logger {
	if logger, ok := logMap[name]; ok {
		return logger
	} else {
		logger := log.New(file, strings.Title(name), logFlags)
		logMap[name] = logger

		return logger
	}
}

func Get(name string) (*log.Logger, bool) {
	return logMap[name]
}

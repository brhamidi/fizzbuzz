package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(...interface{})
	Error(...interface{})
	Fatal(...interface{})
}

func NewLogger(appName string) Logger {
	var log = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
		Level:     logrus.InfoLevel,
	}

	entry := log.WithFields(logrus.Fields{
		"appname": appName,
	})

	return entry
}

package utils

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

var log *logrus.Logger

func InitLogger(logger *logrus.Logger) {
	log = logger
}

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	Log = logrus.New()

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "../log/application.log",
		logrus.WarnLevel:  "../log/application.log",
		logrus.ErrorLevel: "../log/error.log",
		logrus.FatalLevel: "../log/error.log",
	}

	Log.SetLevel(logrus.TraceLevel)

	var hook *lfshook.LfsHook = lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{PrettyPrint: false},
	)
	Log.Hooks.Add(hook)
	return Log
}

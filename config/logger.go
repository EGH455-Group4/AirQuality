package config

import (
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

func LoggerSetup() bool {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)

			return "", fileName
		},
	})

	logrus.Info("Created logger")

	return true
}

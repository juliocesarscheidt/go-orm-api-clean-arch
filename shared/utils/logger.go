package utils

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger = logrus.New()

func init() {
	Logger.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.DebugLevel)
}

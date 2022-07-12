package helper

import (
	"os"

	"github.com/sirupsen/logrus"
)

func ErrorLogging(err error) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Panic(err)
}
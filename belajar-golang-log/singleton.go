package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Hello World")
	logrus.Warn("Warn level")

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Hello World")
	logrus.Warn("Warn level")
}
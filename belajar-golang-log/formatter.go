package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.TraceLevel)
	logger.Trace("Trace level")
	logger.Debug("Debug level")
	logger.Info("Hello World")
	logger.Warn("Warn level")
	logger.Error("Error level")
	logger.Fatal("Fatal level")
	logger.Panic("Panic level")
}
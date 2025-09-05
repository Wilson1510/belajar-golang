package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	logger.SetLevel(logrus.TraceLevel)
	logger.Trace("Trace level")
	logger.Debug("Debug level")
	logger.Info("Hello World")
	logger.Warn("Warn level")
	logger.Error("Error level")
	logger.Fatal("Fatal level")
	logger.Panic("Panic level")
}
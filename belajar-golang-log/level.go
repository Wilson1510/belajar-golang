package main

import "github.com/sirupsen/logrus"

func main() {
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)
	logger.Trace("Level Trace")
	logger.Debug("Level Debug")
	logger.Info("Level Info")
	logger.Warn("Level Warn")
	logger.Error("Level Error")
	logger.Fatal("Level Fatal")
	logger.Panic("Level Panic")
}
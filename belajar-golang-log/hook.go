package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

func (hook *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel}
}

func (hook *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook", entry.Message, entry.Level)
	return nil
}

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.AddHook(&SampleHook{})
	logger.Info("Info level")
	logger.Warn("Warn level")
	logger.Error("Error level")
	logger.Debug("Debug level")
}
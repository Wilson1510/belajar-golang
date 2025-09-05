package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.WithField("username", "john").Trace("Trace level")
	logger.WithField("username", "john").WithField("email", "john@example.com").Debug("Debug level")

	logger.WithFields(logrus.Fields{
		"username": "john",
		"email": "john@example.com",
	}).Info("Info level")
}
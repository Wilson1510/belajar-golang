package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	entry := logrus.NewEntry(logrus.New())
	entry.Info("Hello World")
	entry.WithFields(logrus.Fields{
		"username": "john",
		"email": "john@example.com",
	}).Info("Info level")
}
package main

import (
	"github.com/sirupsen/logrus"
	"fmt"
)

func main() {
	logger := logrus.New()
	logger.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println(logger)
}
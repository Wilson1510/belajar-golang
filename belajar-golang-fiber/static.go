package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/static", "./source")
	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
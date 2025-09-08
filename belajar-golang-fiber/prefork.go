package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	if fiber.IsChild() {
		fmt.Println("In child process")
	} else {
		fmt.Println("In parent process")
	}
	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
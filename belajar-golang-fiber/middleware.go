package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Request received")
		err := c.Next()
		if err != nil {
			return err
		}
		fmt.Println("Request processed")
		return err
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, api!")
	})
	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
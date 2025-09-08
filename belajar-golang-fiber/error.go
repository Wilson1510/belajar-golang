package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"errors"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Status(fiber.StatusInternalServerError)
			return c.SendString("Error: " + err.Error())
		},
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/error", func(c *fiber.Ctx) error {
		return errors.New("terjadi error")
	})
	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
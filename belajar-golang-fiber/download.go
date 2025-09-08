package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Download("source/bar.jpg")
	})
	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
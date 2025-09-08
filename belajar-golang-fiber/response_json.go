package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"id": 1,
			"name": "John",
			"email": "john@example.com",
		})
	})
	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
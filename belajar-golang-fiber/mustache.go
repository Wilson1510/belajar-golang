package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: mustache.New("./template", ".mustache"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Status(fiber.StatusInternalServerError)
			return c.SendString("Error: " + err.Error())
		},
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/template", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Hello Title",
			"header": "Hello Header",
			"description": "Hello Description",
			"body": "Hello Body",
		})
	})
	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
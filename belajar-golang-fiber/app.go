package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")
		return c.SendString("Hello, " + name)
	})

	app.Get("/hello/:name/:age", func(c *fiber.Ctx) error {
		name := c.Params("name")
		age := c.Params("age")
		return c.SendString("Hello, " + name + " and your age is " + age)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		age := c.FormValue("age")
		return c.SendString("Hello, " + name + " and your age is " + age)
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running on port 3000")
}
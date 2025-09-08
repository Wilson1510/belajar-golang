package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func a(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	app := fiber.New()
	api := app.Group("/api")
	api.Get("/hello", a)
	api.Get("/hello2", a)
	api.Get("/", a)

	api2 := app.Group("/api2")
	api2.Get("/hello3", a)
	api2.Get("/hello4", a)

	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
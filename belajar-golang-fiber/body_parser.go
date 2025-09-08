package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		person := Person{}
		err := c.BodyParser(&person)
		if err != nil {
			return err
		}
		fmt.Println(person)
		return c.SendString("Hello, " + person.Name + " and your age is " + strconv.Itoa(person.Age))
	})
	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
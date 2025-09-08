package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		ctx := c.Body()
		request := new(Person)
		err := json.Unmarshal(ctx, &request)
		if err != nil {
			return err
		}
		fmt.Println(request)
		return c.SendString("Hello, " + request.Name + " and your age is " + strconv.Itoa(request.Age))
	})
	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
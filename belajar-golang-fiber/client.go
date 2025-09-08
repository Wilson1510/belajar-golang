package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	agent := client.Get("https://example.com")
	status, response, err := agent.String()
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
	fmt.Println(status)
	fmt.Println(response)
}
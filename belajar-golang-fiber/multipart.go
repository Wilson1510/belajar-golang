package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10MB
	})
	
	// Route untuk halaman utama - menampilkan form upload
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./index.html")
	})
	
	// Route untuk halaman upload form (alternatif)
	app.Get("/upload", func(c *fiber.Ctx) error {
		return c.SendFile("./index.html")
	})
	
	// Route untuk memproses upload file
	app.Post("/upload", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		avatar, err := c.FormFile("avatar")
		if err != nil {
			return err
		}

		err = c.SaveFile(avatar, "target/"+avatar.Filename)
		if err != nil {
			return err
		}

		return c.SendString("Hello, " + username + " and your avatar is " + avatar.Filename)
	})
	
	fmt.Println("Server is running on port 3000")
	fmt.Println("Open http://localhost:3000 to access the upload form")
	app.Listen(":3000")
}

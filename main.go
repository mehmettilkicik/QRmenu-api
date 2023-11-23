package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("api", welcome)
	log.Fatal(app.Listen(":3000"))

}

func welcome(c *fiber.Ctx) error {
	return c.Status(200).JSON("welcome there ")
}

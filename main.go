package main

import (
	"log"
	"qr-menu-api/config"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	//Table endpoints

	//Item endpoints

	//Order endpoints
}

func main() {
	config.ConnectDb()

	app := fiber.New()
	app.Get("api", welcome)
	log.Fatal(app.Listen(":3000"))

}

func welcome(c *fiber.Ctx) error {
	return c.Status(200).JSON("welcome there ")
}

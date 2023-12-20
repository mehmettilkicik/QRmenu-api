package main

import (
	"log"
	"qr-menu-api/config"
	"qr-menu-api/routes"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	//Table endpoints
	app.Post("/api/tables/", routes.CreateTable)
	app.Get("/api/tables/", routes.GetTables)
	app.Get("/api/tables/:id", routes.GetTable)
	app.Put("/api/tables/:id", routes.UpdateTable)
	app.Delete("/api/tables/:id", routes.DeleteTable)

	//Category endpoints
	app.Post("/api/categories", routes.CreateCategory)
	app.Get("/api/categories", routes.GetCategories)
	app.Get("/api/categories/:id", routes.GetCategory)
	app.Put("/api/categories/:id", routes.UpdateCategory)
	app.Delete("/api/categories/:id", routes.DeleteCategory)
	//Item endpoints
	app.Post("/api/items", routes.CreateItem)
	app.Get("/api/items", routes.GetItems)
	app.Get("/api/items/:category_refer", routes.GetItemsByCategory)
	app.Get("api/items/:category_refer/:id", routes.GetItem)
	app.Put("api/items/:category_refer/:id", routes.UpdateItem)
	app.Delete("api/items/:category_refer/:id", routes.DeleteItem)
	//Order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("api/orders/:table_refer", routes.GetAllOrders)
}

func main() {
	config.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	app.Get("api", welcome)
	log.Fatal(app.Listen(":3000"))

}

func welcome(c *fiber.Ctx) error {
	return c.Status(200).JSON("welcome there ")
}

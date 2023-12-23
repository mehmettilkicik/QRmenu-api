package main

import (
	"log"
	"qr-menu-api/config"
	"qr-menu-api/routes"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("api/")

	table_api := api.Group("tables/")

	//Table endpoints
	table_api.Post("/", routes.CreateTable)
	table_api.Get("/", routes.GetTables)
	table_api.Get("/:id", routes.GetTable)
	table_api.Put("/:id", routes.UpdateTable)
	table_api.Delete("/:id", routes.DeleteTable)

	category_api := api.Group("categories/")
	//Category endpoints
	category_api.Post("", routes.CreateCategory)
	category_api.Get("", routes.GetCategories)
	category_api.Get(":id", routes.GetCategory)
	category_api.Put(":id", routes.UpdateCategory)
	category_api.Delete(":id", routes.DeleteCategory)

	item_api := api.Group("items/")
	//Item endpoints
	item_api.Post("", routes.CreateItem)
	item_api.Get("", routes.GetItems)
	item_api.Get(":category_refer", routes.GetItemsByCategory)
	item_api.Get(":category_refer/:id", routes.GetItem)
	item_api.Put(":category_refer/:id", routes.UpdateItem)
	item_api.Delete(":category_refer/:id", routes.DeleteItem)

	order_api := api.Group("orders/")
	//Order endpoints
	order_api.Post("", routes.CreateOrder)
	order_api.Get("", routes.GetOrders)
	order_api.Get(":is_paid", routes.GetActiveOrInactiveOrders)
	order_api.Get(":is_paid/:table_refer", routes.GetOrdersByTable)
	order_api.Get(":is_paid/:table_refer/:id", routes.GetSpecificOrder)
	order_api.Post(":is_paid/:table_refer/:id", routes.UpdateOrderDetail)
	order_api.Put(":is_paid/:table_refer/:id", routes.UpdateOrder)
}

func main() {
	config.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}

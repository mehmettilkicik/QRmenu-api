package routes

import (
	"qr-menu-api/config"
	"qr-menu-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type OrderDetail struct {
	OrderID  uint `json:"order_id"`
	Item     Item
	Quantity uint `json:"quantity"`
}

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Table        Table
	OrderDetails []OrderDetail `gorm:"foreignKey:order_id"`
}

func CreateResponseOrder(order models.Order, table Table, orderdetails []OrderDetail) Order {
	return Order{ID: order.ID, CreatedAt: order.CreatedAt, Table: table, OrderDetails: orderdetails}
}
func CreateResponseOrderDetail(orderdetail models.OrderDetail, order models.Order, Item Item) OrderDetail {
	return OrderDetail{OrderID: order.ID, Item: Item, Quantity: orderdetail.Quantity}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var orderDetails []models.OrderDetail
	for _, detail := range order.OrderDetails {
		orderDetails = append(orderDetails, detail)
	}
	config.Database.Db.Create(&orderDetails)

	var table models.Table
	if err := findTable(order.TableRefer, &table); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	config.Database.Db.Create(&order)

	responseDetails := []OrderDetail{}

	for _, v := range orderDetails {
		var item models.Item
		var category models.Category
		if err := findItemOrder(v.ItemRefer, &item); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		if err := findCategory(item.CategoryRefer, &category); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		responseCategory := CreateResponseCategory(category)
		responseItem := CreateResponseItem(item, responseCategory)
		responseDetail := CreateResponseOrderDetail(v, order, responseItem)
		responseDetails = append(responseDetails, responseDetail)
	}
	responseTable := CreateResponseTable(table)
	responseOrder := CreateResponseOrder(order, responseTable, responseDetails)

	return c.Status(200).JSON(responseOrder)
}

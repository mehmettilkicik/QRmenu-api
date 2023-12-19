package routes

import (
	"qr-menu-api/config"
	"qr-menu-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Table        Table         `gorm:"foreignKey:TableRefer"`
	OrderDetails []OrderDetail `json:"order_details" gorm:"foreignKey:OrderID"`
	IsPaid       bool          `json:"is_paid"`
}

type OrderDetail struct {
	OrderID  uint `json:"order_id" gorm:"index"`
	Item     Item
	Quantity uint `json:"quantity"`
}

func CreateResponseOrder(order models.Order, table Table, ordedetails []OrderDetail) Order {
	return Order{ID: order.ID, CreatedAt: order.CreatedAt, Table: table, OrderDetails: ordedetails}
}

func CreateResponseOrderDetail(ordedetail models.OrderDetail, item Item) OrderDetail {
	return OrderDetail{OrderID: ordedetail.OrderID, Item: item, Quantity: ordedetail.Quantity}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var table models.Table

	if err := findTable(order.TableRefer, &table); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseDetails := []OrderDetail{}
	config.Database.Db.Create(&order)
	for _, v := range order.OrderDetails {
		v.OrderID = order.ID
		config.Database.Db.Create(&v)
		var item models.Item
		if err := findItemByID(v.ItemRefer, &item); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		var category models.Category
		if err := findCategory(item.CategoryRefer, &category); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		responseCategory := CreateResponseCategory(category)
		responseItem := CreateResponseItem(item, responseCategory)
		responseDetails = append(responseDetails, CreateResponseOrderDetail(v, responseItem))
	}
	responseTable := CreateResponseTable(table)
	responseOrder := CreateResponseOrder(order, responseTable, responseDetails)

	return c.Status(200).JSON(responseOrder)
}

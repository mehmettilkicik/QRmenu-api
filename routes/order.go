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
}

type OrderDetail struct {
	OrderID  uint `json:"order_id" gorm:"index"`
	ItemID   uint `json:"item_id"`
	Quantity uint `json:"quantity"`
}

func CreateResponseOrder(order models.Order, table Table, ordedetails []OrderDetail) Order {
	return Order{ID: order.ID, CreatedAt: order.CreatedAt, Table: table, OrderDetails: ordedetails}
}

func CreateResponseOrderDetail(order models.Order, ordedetail models.OrderDetail) OrderDetail {
	return OrderDetail{OrderID: order.ID, ItemID: ordedetail.ItemID, Quantity: ordedetail.Quantity}
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

	orderdetails := []models.OrderDetail{}
	config.Database.Db.Create(&order)
	for _, v := range order.OrderDetails {
		v.OrderID = order.ID
		orderdetails = append(orderdetails, v)
	}

	return c.Status(200).JSON(orderdetails)
}

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
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	OrderID  uint `gorm:"foreignKey:OrderID"`
	ItemID   uint `json:"item_id"`
	Quantity uint `json:"quantity"`
}

func CreateResponseOrderDetail(orderdetail models.OrderDetail) OrderDetail {
	return OrderDetail{OrderID: orderdetail.OrderID, ItemID: orderdetail.ItemID, Quantity: orderdetail.Quantity}
}

func CreateResponseOrder(order models.Order, table Table, orderdetails []OrderDetail) Order {
	return Order{ID: order.ID, CreatedAt: order.CreatedAt, Table: table, OrderDetails: orderdetails}
}

func CreateOrderDetail(order models.Order, orderdetail models.OrderDetail) error {
	orderdetail.OrderID = order.ID
	config.Database.Db.Create(&orderdetail)
	return nil
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

	var responseOrderDetails []OrderDetail

	for _, v := range order.OrderDetails {
		CreateOrderDetail(order, v)
	}
	config.Database.Db.Create(&order)
	for _, v := range order.OrderDetails {
		responseOrderDetail := CreateResponseOrderDetail(v)
		responseOrderDetails = append(responseOrderDetails, responseOrderDetail)
	}

	responseTable := CreateResponseTable(table)
	responseOrder := CreateResponseOrder(order, responseTable, responseOrderDetails)
	return c.Status(200).JSON(responseOrder)
}

package routes

import (
	"qr-menu-api/models"
	"time"
)

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Table     Table
}

func CreateResponseOrder(order models.Order, table Table) Order {
	return Order{ID: order.ID, CreatedAt: order.CreatedAt, Table: table}
}

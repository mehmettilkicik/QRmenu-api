package routes

import (
	"qr-menu-api/models"
	"time"
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

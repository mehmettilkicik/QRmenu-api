package routes

import "time"

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

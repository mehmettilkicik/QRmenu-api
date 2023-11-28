package models

import "time"

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	TableRefer   int           `json:"table_id"`
	Table        Table         `gorm:"foreignKey:TableRefer"`
	OrderDetails []OrderDetail `json:"order_details" gorm:"foreignKey:OrderID"`
}

package models

import "time"

type Order struct {
	ID           uint `json:"primaryKey"`
	CreatedAt    time.Time
	TableRefer   int           `json:"table_id"`
	Table        Table         `gorm:"foreignKey:TableRefer"`
	OrderDetails []OrderDetail `gorm:"foreignKey:order_id"`
}

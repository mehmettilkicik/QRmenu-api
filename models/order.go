package models

import "time"

type Order struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time
	ItemRefer  map[int]int  `json:"quantity_and_item_id"`
	Item       map[int]Item `gorm:"foreignKey:ItemRefer"`
	TableRefer int          `json:"table_id"`
	Table      Table        `gorm:"foreignKey:TableRefer"`
}

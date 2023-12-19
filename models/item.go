package models

import "time"

type Item struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time
	Name          string   `json:"name"`
	CategoryRefer int      `json:"category_id"`
	Category      Category `gorm:"foreignKey:CategoryRefer"`
	Price         string   `json:"price"`
}

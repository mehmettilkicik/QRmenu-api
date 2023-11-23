package models

import "time"

type Item struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time
	Name          string `json:"name"`
	CategoryRefer uint   `json:"category_id"`
	Category      string `gorm:"foreignKey:CategoryRefer"`
}

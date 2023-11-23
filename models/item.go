package models

import "time"

type Item struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time
	Name       string `json:"name"`
	Categories string `json:"categories"`
}

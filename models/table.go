package models

import "time"

type Table struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Number    uint `json:"number"`
}

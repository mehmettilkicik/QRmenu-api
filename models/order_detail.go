package models

type OrderDetail struct {
	OrderID  uint `json:"order_id" gorm:"index"`
	ItemID   uint `json:"item_id"`
	Quantity uint `json:"quantity"`
}

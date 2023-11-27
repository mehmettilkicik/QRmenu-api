package models

type OrderDetail struct {
	OrderID  uint `json:"order_id"`
	ItemID   uint `json:"item_id"`
	Quantity uint `json:"quantity"`
}

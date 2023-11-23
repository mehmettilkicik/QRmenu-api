package models

//Pivot Table

type OrderDetail struct {
	OrderID   uint `json:"order_id"`
	ItemRefer uint `json:"item_id"`
	Item      Item `gorm:"foreignKey:ItemRefer"`
	Quantity  uint `json:"quantity"`
}

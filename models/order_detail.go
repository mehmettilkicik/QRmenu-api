package models

type OrderDetail struct {
	OrderID   uint `json:"order_id" gorm:"index"`
	ItemRefer int  `json:"item_id"`
	Item      Item `gorm:"foreignKey:ItemRefer"`
	Quantity  uint `json:"quantity"`
}

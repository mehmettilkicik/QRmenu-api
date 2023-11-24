package routes

type Item struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
}

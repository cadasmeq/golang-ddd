package model

type Payment struct {
	OrderID string `gorm:"type:uuid,primaryKey"`
	CardID  string `gorm:"type:uuid"`
	Card    Card
}

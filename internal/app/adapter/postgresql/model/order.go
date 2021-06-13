package model

type Order struct {
	ID       string `gorm:"column:order_id;type:uuid"`
	personID string `gorm>"type:uuid"`
	Person   Person
	Payment  Payment
}

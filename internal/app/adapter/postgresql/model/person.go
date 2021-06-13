package model

type Person struct {
	ID   string `gorm:"column:person_id;type:uuid"`
	Name string `gorm:"type:text";not null`
}

func (Person) TableName() string {
	return "persons"
}

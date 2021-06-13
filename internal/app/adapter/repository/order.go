package repository

import (
	"errors"

	"github.com/cadasmeq/golang-ddd/internal/app/adapter/postgresql"
	"github.com/cadasmeq/golang-ddd/internal/app/adapter/postgresql/model"
	"github.com/cadasmeq/golang-ddd/internal/app/domain"
	"github.com/cadasmeq/golang-ddd/internal/app/domain/factory"
	"gorm.io/gorm"
)

type Order struct{}

func (o Order) Get() domain.Order {
	db := postgresql.Connection()
	var order model.Order
	result := db.Preload("Person").Preload("Payment.Card.CardBrand").Find(&order)
	if result.Error != nil {
		panic(result.Error)
	}
	orderFactory := factory.Order{}
	return orderFactory.Generate(
		order.Person.ID,
		order.Person.Name,
		order.Payment.Card.ID,
		order.Payment.Card.CardBrand.Brand,
		order.ID,
	)
}

func (o Order) Update(order domain.Order) {
	db := postgresql.Connection()
	card := model.Card{
		ID:    order.Payment.Card.ID,
		Brand: string(order.Payment.Card.Brand),
	}
	payment := model.Payment{
		OrderID: order.ID,
		CardID:  card.ID,
		Card:    card,
	}

	person := model.Person{
		ID:   order.Person.ID,
		Name: order.Person.Name,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		var err error
		err = tx.Exec("update persons set name = ? where person_id = ?", person.Name, person.ID).Error
		if err != nil {
			return errors.New("rollback")
		}

		err = tx.Exec("insert into cards values (?, ?)", card.ID, card.Brand).Error
		if err != nil {
			return errors.New("rollback")
		}

		err = tx.Exec("update payments set card_id = ? where order_id = ?", card.ID, payment.OrderID).Error
		if err != nil {
			return errors.New("rollback")
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

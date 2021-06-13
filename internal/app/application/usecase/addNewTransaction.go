package usecase

import (
	"github.com/cadasmeq/golang-ddd/internal/app/domain"
	"github.com/cadasmeq/golang-ddd/internal/app/domain/repository"
	"github.com/cadasmeq/golang-ddd/internal/app/domain/valueobject"
	"github.com/google/uuid"
)

func AddNewTransaction(o repository.OrderRepository) domain.Order {
	order := o.Get()
	newCardBrand := valueobject.VISA
	if order.Payment.Card.Brand == valueobject.VISA {
		newCardBrand = valueobject.MASTERCARD
	}
	newCard := valueobject.Card{
		ID:    uuid.New().String(),
		Brand: newCardBrand,
	}

	order.Payment.Card = newCard
	o.Update(order)
	return order
}

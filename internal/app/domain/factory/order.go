package factory

import (
	"github.com/cadasmeq/golang-ddd/internal/app/domain"
	"github.com/cadasmeq/golang-ddd/internal/app/domain/valueobject"
)

type Order struct{}

func (o Order) Generate(
	personID string,
	name string,
	cardID string,
	brand string,
	orderID string,
) domain.Order {
	person := domain.Person{
		ID:   personID,
		Name: name,
	}
	cardBrand := valueobject.ConvertToCardBrand(brand)
	card := valueobject.Card{
		ID:    cardID,
		Brand: cardBrand,
	}
	payment := valueobject.Payment{
		Card: card,
	}

	return domain.Order{
		ID:      orderID,
		Payment: payment,
		Person:  person,
	}
}

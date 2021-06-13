package domain

import "github.com/cadasmeq/golang-ddd/internal/app/domain/valueobject"

type Order struct {
	ID      string
	Payment valueobject.Payment
	Person  Person
}

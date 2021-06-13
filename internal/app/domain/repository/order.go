package repository

import "github.com/cadasmeq/golang-ddd/internal/app/domain"

type OrderRepository interface {
	Get() domain.Order
	Update(domain.Order)
}

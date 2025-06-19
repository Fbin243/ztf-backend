package business

import "ztf-backend/internal/entity"

type IOrderRepo interface {
	findAll() ([]entity.Order, error)
	findById(id string) (*entity.Order, error)
	insertOne(order *entity.Order) (string, error)
	updateOne(order *entity.Order) (string, error)
	deleteOne(id string) (string, error)
}

type OrderBusiness struct {
	orderRepo IOrderRepo
}

func NewOrderBusiness(orderRepo IOrderRepo) *OrderBusiness {
	return &OrderBusiness{
		orderRepo: orderRepo,
	}
}

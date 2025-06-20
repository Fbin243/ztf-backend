package biz

import "ztf-backend/internal/entity"

type IOrderRepo interface {
	FindAll() ([]entity.Order, error)
	FindById(id uint) (*entity.Order, error)
	InsertOne(order *entity.Order) (uint, error)
	UpdateOne(order *entity.Order) (uint, error)
	DeleteOne(id uint) (uint, error)
}

type OrderBusiness struct {
	orderRepo IOrderRepo
}

func NewOrderBusiness(orderRepo IOrderRepo) *OrderBusiness {
	return &OrderBusiness{
		orderRepo: orderRepo,
	}
}

package biz

import "ztf-backend/internal/entity"

type IOrderRepo interface {
	Exists(id uint) (bool, error)
	FindAll() ([]entity.Order, error)
	FindById(id uint) (*entity.Order, error)
	InsertOne(order *entity.Order) (uint, error)
	UpdateOne(order *entity.Order) (uint, error)
	DeleteOne(id uint) (uint, error)
}

type OrderBusiness struct {
	orderRepo  IOrderRepo
	couponRepo ICouponRepo
}

func NewOrderBusiness(orderRepo IOrderRepo, couponRepo ICouponRepo) *OrderBusiness {
	return &OrderBusiness{
		orderRepo:  orderRepo,
		couponRepo: couponRepo,
	}
}

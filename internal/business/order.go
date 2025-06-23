package biz

import (
	"ztf-backend/internal/entity"

	"github.com/google/uuid"
)

type IOrderRepo interface {
	Exists(id uuid.UUID) (bool, error)
	FindAll() ([]entity.Order, error)
	FindById(id uuid.UUID) (*entity.Order, error)
	InsertOne(order *entity.Order) (uuid.UUID, error)
	UpdateOne(order *entity.Order) (uuid.UUID, error)
	DeleteOne(id uuid.UUID) (uuid.UUID, error)
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

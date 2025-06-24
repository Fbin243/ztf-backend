package biz

import (
	entity2 "ztf-backend/order/internal/entity"
)

type IUserRepo interface {
	Exists(id string) (bool, error)
	FindByIds(ids []string) ([]entity2.User, error)
}

type IMerchantRepo interface {
	Exists(id string) (bool, error)
	FindByIds(ids []string) ([]entity2.Merchant, error)
}

type IOrderRepo interface {
	Exists(id string) (bool, error)
	FindAll() ([]entity2.Order, error)
	FindById(id string) (*entity2.Order, error)
	FindByIds(ids []string) ([]entity2.Order, error)
	InsertOne(order *entity2.Order) (string, error)
	UpdateOne(order *entity2.Order) (string, error)
	DeleteOne(id string) (string, error)
	FindByIdWithMerchantAndUser(id string) (*entity2.Order, error)
	UpdateUserId(id string, userId string) (string, error)
}

type OrderBusiness struct {
	orderRepo    IOrderRepo
	merchantRepo IMerchantRepo
	userRepo     IUserRepo
	couponRepo   ICouponRepo
}

func NewOrderBusiness(
	orderRepo IOrderRepo,
	couponRepo ICouponRepo,
	userRepo IUserRepo,
	merchantRepo IMerchantRepo,
) *OrderBusiness {
	return &OrderBusiness{
		orderRepo:    orderRepo,
		couponRepo:   couponRepo,
		userRepo:     userRepo,
		merchantRepo: merchantRepo,
	}
}

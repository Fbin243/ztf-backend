package biz

import (
	"ztf-backend/order/internal/entity"
)

func (b *OrderBusiness) FindAll() ([]entity.Order, error) {
	return b.orderRepo.FindAll()
}

func (b *OrderBusiness) FindById(id string) (*entity.Order, error) {
	return b.orderRepo.FindById(id)
}

func (b *OrderBusiness) FindByIdWithMerchantAndUser(id string) (*entity.Order, error) {
	order, err := b.orderRepo.FindByIdWithMerchantAndUser(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (b *OrderBusiness) FindByIds(ids []string) ([]entity.Order, error) {
	return b.orderRepo.FindByIds(ids)
}

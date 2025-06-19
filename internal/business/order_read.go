package biz

import "ztf-backend/internal/entity"

func (b *OrderBusiness) FindAll() ([]entity.Order, error) {
	return b.orderRepo.FindAll()
}

func (b *OrderBusiness) FindById(id string) (*entity.Order, error) {
	return b.orderRepo.FindById(id)
}

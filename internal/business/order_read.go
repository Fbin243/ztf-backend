package business

import "ztf-backend/internal/entity"

func (b *OrderBusiness) FindAll() ([]entity.Order, error) {
	return b.orderRepo.findAll()
}

func (b *OrderBusiness) FindById(id string) (*entity.Order, error) {
	return b.orderRepo.findById(id)
}

package biz

import (
	"context"
	"ztf-backend/services/order/internal/entity"
)

func (b *OrderBusiness) FindAll(ctx context.Context) ([]entity.Order, error) {
	return b.orderRepo.FindAll(ctx)
}

func (b *OrderBusiness) FindById(ctx context.Context, id string) (*entity.Order, error) {
	return b.orderRepo.FindById(ctx, id)
}

func (b *OrderBusiness) FindByIdWithMerchantAndUser(
	ctx context.Context,
	id string,
) (*entity.Order, error) {
	order, err := b.orderRepo.FindByIdWithMerchantAndUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (b *OrderBusiness) FindByIds(ctx context.Context, ids []string) ([]entity.Order, error) {
	return b.orderRepo.FindByIds(ctx, ids)
}

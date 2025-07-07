package biz

import (
	"context"
	"ztf-backend/services/order/internal/entity"
)

func (b *OrderBusiness) GetOrderList(ctx context.Context) ([]entity.Order, error) {
	return b.orderRepo.FindAll(ctx)
}

func (b *OrderBusiness) GetOrder(ctx context.Context, id int64) (*entity.Order, error) {
	return b.orderRepo.FindById(ctx, id)
}

func (b *OrderBusiness) GetOrderWithMerchantAndUser(
	ctx context.Context,
	id int64,
) (*entity.Order, error) {
	order, err := b.orderRepo.FindByIdWithMerchantAndUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (b *OrderBusiness) FindByIds(ctx context.Context, ids []int64) ([]entity.Order, error) {
	return b.orderRepo.FindByIds(ctx, ids)
}

package biz

import (
	"context"
	"errors"
	"log"
	"ztf-backend/services/order/internal/auth"
	"ztf-backend/services/order/internal/entity"

	"github.com/jinzhu/copier"

	errs "ztf-backend/services/order/internal/errors"
)

func (b *OrderBusiness) CreateOrder(
	ctx context.Context,
	input *entity.CreateOrderInput,
) (int64, error) {
	// Check if the merchant exists
	exists, err := b.merchantRepo.Exists(ctx, input.MerchantId)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, errs.ErrorNotFound
	}

	newOrder := &entity.Order{}
	err = copier.Copy(newOrder, input)
	if err != nil {
		return 0, err
	}

	return b.orderRepo.InsertOne(ctx, newOrder)
}

func (b *OrderBusiness) UpdateOrder(
	ctx context.Context,
	id int64,
	input *entity.UpdateOrderInput,
) (int64, error) {
	// Get the existing order
	existingOrder, err := b.orderRepo.FindById(ctx, id)
	if err != nil {
		return 0, err
	}

	err = copier.Copy(existingOrder, input)
	if err != nil {
		return 0, err
	}

	return b.orderRepo.UpdateOne(ctx, existingOrder)
}

func (b *OrderBusiness) DeleteOrder(ctx context.Context, id int64) (int64, error) {
	// Check if the order exists
	exists, err := b.orderRepo.Exists(ctx, id)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, errs.ErrorNotFound
	}

	return b.orderRepo.DeleteOne(ctx, id)
}

func (b *OrderBusiness) PayForOrder(
	ctx context.Context,
	id int64,
	input *entity.PayOrderInput,
) (int64, error) {
	userId, err := auth.GetAuthKey(ctx)
	if err != nil {
		return 0, err
	}

	// Check if the order exists
	order, err := b.orderRepo.FindById(ctx, id)
	if err != nil {
		return 0, err
	}

	if order.UserId != nil {
		return 0, errors.New("order is already paid")
	}

	// Check the amount of the order
	if order.Amount != input.Amount ||
		order.Amount != input.PromotionAmount+input.PayAmount {
		return 0, errors.New("invalid amount")
	}

	// Check if the user exists
	exists, err := b.userRepo.Exists(ctx, userId)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, errs.ErrorNotFound
	}

	// Apply the promotion
	if input.PromotionId != nil {
		success, err := b.promotionClient.ApplyPromotion(ctx, &entity.ApplyPromotionReq{
			PromotionId:     *input.PromotionId,
			UserId:          userId,
			OrderId:         id,
			Amount:          input.Amount,
			PromotionAmount: input.PromotionAmount,
		})
		if err != nil {
			return 0, err
		}
		if !success {
			return 0, errors.New("promotion is not applied")
		}
	} else if input.PromotionAmount != 0 {
		return 0, errors.New("missing promotion id")
	}

	log.Printf("Promotion is applied")

	// Pay for the order
	order.UserId = &userId
	order.PromotionId = input.PromotionId
	order.PromotionAmount = input.PromotionAmount
	order.PayAmount = input.PayAmount

	id, err = b.orderRepo.UpdatePaymentInfo(ctx, id, order)
	if errors.Is(err, errs.ErrorNoRowsAffected) {
		return 0, errors.New("order is already paid")
	}
	if err != nil {
		return 0, err
	}

	return id, nil
}

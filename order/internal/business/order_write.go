package biz

import (
	"context"
	"errors"

	"ztf-backend/order/internal/entity"
	"ztf-backend/pkg/db/base"
	errs "ztf-backend/pkg/errors"

	"github.com/jinzhu/copier"
)

func (b *OrderBusiness) InsertOne(input *entity.CreateOrderInput) (string, error) {
	// Check if the merchant exists
	exists, err := b.merchantRepo.Exists(input.MerchantId)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errs.ErrorNotFound
	}

	newOrder := &entity.Order{
		BaseEntity: &base.BaseEntity{},
	}
	err = copier.Copy(newOrder, input)
	if err != nil {
		return "", err
	}

	return b.orderRepo.InsertOne(newOrder)
}

func (b *OrderBusiness) UpdateOne(id string, input *entity.UpdateOrderInput) (string, error) {
	// Get the existing order
	existingOrder, err := b.orderRepo.FindById(id)
	if err != nil {
		return "", err
	}

	err = copier.Copy(existingOrder, input)
	if err != nil {
		return "", err
	}

	return b.orderRepo.UpdateOne(existingOrder)
}

func (b *OrderBusiness) DeleteOne(id string) (string, error) {
	// Check if the order exists
	exists, err := b.orderRepo.Exists(id)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errs.ErrorNotFound
	}

	return b.orderRepo.DeleteOne(id)
}

func (b *OrderBusiness) PayForOrder(id string, input *entity.PayOrderInput) (string, error) {
	// Check if the order exists
	order, err := b.orderRepo.FindById(id)
	if err != nil {
		return "", err
	}

	if order.UserId != nil {
		return "", errors.New("order is already paid")
	}

	// Check the amount of the order
	if order.Amount != input.Amount ||
		order.Amount != input.PromotionAmount+input.PayAmount {
		return "", errors.New("invalid amount")
	}

	// Check if the user exists
	exists, err := b.userRepo.Exists(input.UserId)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errs.ErrorNotFound
	}

	// Verify the promotion
	if input.PromotionId != nil {
		verified, err := b.promotionClient.VerifyPromotion(context.Background(), &entity.VerifyPromotionReq{
			PromotionId:     *input.PromotionId,
			UserId:          input.UserId,
			OrderId:         id,
			Amount:          input.Amount,
			PromotionAmount: input.PromotionAmount,
		})
		if err != nil {
			return "", err
		}
		if !verified {
			return "", errors.New("promotion is not valid")
		}
	} else if input.PromotionAmount != 0 {
		return "", errors.New("missing promotion id")
	}

	// Pay for the order
	order.UserId = &input.UserId
	order.PromotionId = input.PromotionId
	order.PromotionAmount = input.PromotionAmount
	order.PayAmount = input.PayAmount

	id, err = b.orderRepo.UpdatePaymentInfo(id, order)
	if errors.Is(err, errs.ErrorNoRowsAffected) {
		return "", errors.New("order is already paid")
	}
	if err != nil {
		return "", err
	}

	return id, nil
}

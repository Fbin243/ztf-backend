package biz

import (
	"context"
	"errors"
	"log"
	"time"

	"ztf-backend/pkg/db"
	"ztf-backend/pkg/db/base"
	errs "ztf-backend/pkg/errors"
	"ztf-backend/promotion/internal/entity"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func (b *PromotionBusiness) InsertOne(
	ctx context.Context,
	input *entity.CreatePromotionInput,
) (string, error) {
	newPromotion := &entity.Promotion{
		BaseEntity: &base.BaseEntity{},
	}

	err := copier.Copy(newPromotion, input)
	if err != nil {
		return "", err
	}

	return b.promotionRepo.InsertOne(ctx, newPromotion)
}

func (b *PromotionBusiness) UpdateOne(
	ctx context.Context,
	id string,
	input *entity.UpdatePromotionInput,
) (string, error) {
	// Check if the promotion exists
	existingPromotion, err := b.promotionRepo.FindById(ctx, id)
	if err != nil {
		return "", err
	}

	err = copier.Copy(existingPromotion, input)
	if err != nil {
		return "", err
	}

	return b.promotionRepo.UpdateOne(ctx, existingPromotion)
}

func (b *PromotionBusiness) DeleteOne(ctx context.Context, id string) (string, error) {
	// Check if the promotion exists
	exists, err := b.promotionRepo.Exists(ctx, id)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errs.ErrorNotFound
	}

	return b.promotionRepo.DeleteOne(ctx, id)
}

func (b *PromotionBusiness) VerifyPromotion(
	ctx context.Context,
	req *entity.VerifyPromotionReq,
) (bool, error) {
	log.Printf("Verifying promotion: %+v", req)

	// Check if the promotion exists
	promotion, err := b.promotionRepo.FindById(ctx, req.PromotionId)
	if err != nil {
		return false, err
	}

	// Check if the promotion is expired
	if promotion.ExpirationDate.Before(time.Now()) {
		return false, errors.New("promotion is expired")
	}

	// Check the promotion amount
	if promotion.CalculatePromotionAmount(req.Amount) != req.PromotionAmount {
		return false, errors.New("promotion amount is invalid")
	}

	// Check if the promotion belongs to the user
	if !promotion.IsForAll {
		exists, err := b.userPromotionRepo.Exists(ctx, req.UserId, req.PromotionId)
		if err != nil {
			return false, err
		}
		if !exists {
			return false, errors.New("promotion is not for the user")
		}
	}

	return true, nil
}

func (b *PromotionBusiness) ApplyPromotion(
	ctx context.Context,
	req *entity.VerifyPromotionReq,
) (bool, error) {
	// Validate the promotion
	valid, err := b.VerifyPromotion(ctx, req)
	if err != nil {
		return false, err
	}
	if !valid {
		return false, errors.New("promotion is invalid")
	}

	// Apply the promotion
	promotion, err := b.promotionRepo.FindById(ctx, req.PromotionId)
	if err != nil {
		return false, err
	}

	if promotion.IsForAll {
		// - Make a transaction
		db.GetDB().Transaction(func(tx *gorm.DB) error {
			// -- Check if the promotion is used by the user
			userPromotion, err := b.userPromotionRepo.WithTx(tx).FindByUserIdAndPromotionId(ctx, req.UserId, req.PromotionId)
			if err == errs.ErrorNotFound {
				userPromotion = &entity.UserPromotion{
					UserId:      req.UserId,
					PromotionId: req.PromotionId,
				}
				_, _, err = b.userPromotionRepo.WithTx(tx).UpsertOne(ctx, userPromotion)
				if err != nil {
					return err
				}
			}
			if err != nil {
				return err
			}

			// -- Reduce the remaining_count in promotion by 1
			err = b.promotionRepo.WithTx(tx).UpdateRemainingCount(ctx, req.PromotionId)
			if err != nil {
				return err
			}

			// -- Mark as used
			err = b.userPromotionRepo.WithTx(tx).MarkAsUsed(ctx, &entity.MarkAsUsedReq{
				UserId:      req.UserId,
				PromotionId: req.PromotionId,
			})
			if err != nil {
				return err
			}

			return nil
		})
	} else {
		// - Mark as used
		err = b.userPromotionRepo.MarkAsUsed(ctx, &entity.MarkAsUsedReq{
			UserId:      req.UserId,
			PromotionId: req.PromotionId,
		})
		if errors.Is(err, errs.ErrorNoRowsAffected) {
			return false, errors.New("promotion already used max times")
		}
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

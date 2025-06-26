package biz

import (
	"context"
	"errors"
	"log"
	"time"

	"ztf-backend/pkg/db"
	"ztf-backend/pkg/db/base"
	errs "ztf-backend/pkg/errors"
	"ztf-backend/services/promotion/internal/entity"

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

func (b *PromotionBusiness) CollectPromotion(ctx context.Context, userId string, promotionId string) (bool, error) {
	// Check if the promotion exists and its type is not for all
	promotion, err := b.promotionRepo.FindById(ctx, promotionId)
	if err != nil {
		return false, err
	}
	if promotion.IsForAll {
		return false, errors.New("promotion is for all")
	}

	// Check if the promotion is expired
	if promotion.ExpirationDate.Before(time.Now()) {
		return false, errors.New("promotion is expired")
	}

	// Check if user exists
	valid, err := b.orderClient.ValidateUser(ctx, &entity.ValidateUserReq{
		UserId: userId,
	})
	if err != nil {
		return false, err
	}
	if !valid {
		return false, errors.New("user is not valid")
	}

	// Check if the promotion is already collected by the user
	exists, err := b.userPromotionRepo.Exists(ctx, userId, promotionId)
	if err != nil {
		return false, err
	}
	if exists {
		return false, errors.New("promotion is already collected")
	}

	// Collect the promotion
	err = db.GetDB().Transaction(func(tx *gorm.DB) error {
		// - Reduce the remaining_count in promotion by 1
		err = b.promotionRepo.WithTx(tx).UpdateRemainingCount(ctx, promotionId)
		if err != nil {
			return err
		}

		// - Insert a new user_promotion
		_, _, err = b.userPromotionRepo.WithTx(tx).UpsertOne(ctx, &entity.UserPromotion{
			UserId:      userId,
			PromotionId: promotionId,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, errors.New("failed to collect promotion: " + err.Error())
	}

	return true, nil
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
	req *entity.ApplyPromotionReq,
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
	log.Printf("Promotion amount is valid")

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
	req *entity.ApplyPromotionReq,
) (bool, error) {
	// Validate the promotion
	valid, err := b.VerifyPromotion(ctx, req)
	if err != nil {
		return false, err
	}
	if !valid {
		return false, errors.New("promotion is invalid")
	}

	log.Printf("Promotion is valid")

	// Apply the promotion
	promotion, err := b.promotionRepo.FindById(ctx, req.PromotionId)
	if err != nil {
		return false, err
	}

	log.Printf("Promotion: %+v", promotion)
	if promotion.IsForAll {
		// - Make a transaction
		err := db.GetDB().Transaction(func(tx *gorm.DB) error {
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
			if errors.Is(err, errs.ErrorNoRowsAffected) {
				return errors.New("promotion already used max times")
			}
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return false, err
		}

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

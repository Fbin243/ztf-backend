package biz

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"ztf-backend/services/order/pkg/convert"
	"ztf-backend/services/order/pkg/locker"
	"ztf-backend/services/promotion/internal/auth"
	"ztf-backend/services/promotion/internal/entity"

	"github.com/jinzhu/copier"

	errs "ztf-backend/services/promotion/internal/errors"
)

func (b *PromotionBusiness) CreatePromotion(
	ctx context.Context,
	input *entity.CreatePromotionInput,
) (int64, error) {
	newPromotion := &entity.Promotion{}

	err := copier.Copy(newPromotion, input)
	if err != nil {
		return 0, err
	}

	return b.promotionRepo.InsertOne(ctx, newPromotion)
}

func (b *PromotionBusiness) UpdatePromotion(
	ctx context.Context,
	id int64,
	input *entity.UpdatePromotionInput,
) (int64, error) {
	// Check if the promotion exists
	existingPromotion, err := b.promotionRepo.FindById(ctx, id)
	if err != nil {
		return 0, err
	}

	err = copier.Copy(existingPromotion, input)
	if err != nil {
		return 0, err
	}

	return b.promotionRepo.UpdateOne(ctx, existingPromotion)
}

func (b *PromotionBusiness) CollectPromotion(
	ctx context.Context,
	promotionId int64,
) (bool, error) {
	userId, err := auth.GetAuthKey(ctx)
	if err != nil {
		return false, err
	}

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

	// Acquire a lock
	locker := locker.NewLocker(
		"promotion:collect:"+convert.ConvIntToStr(promotionId),
		10*time.Second,
		3,
		500*time.Millisecond,
	)
	err = locker.Acquire()
	if err != nil {
		return false, errors.New("failed to acquire lock: " + err.Error())
	}

	// log.Printf("Promotion %d is being collected by user %d", promotionId, userId)
	// time.Sleep(30000 * time.Millisecond) // Simulate some processing time

	// Collect the promotion
	err = b.txRunner.Transaction(ctx, func(tx Tx) error {
		// - Reduce the remaining_count in promotion by 1
		err = tx.PromotionRepo().UpdateRemainingCount(ctx, promotionId)
		if errors.Is(err, errs.ErrorNoRowsAffected) {
			return errors.New("promotion is collected out")
		}
		if err != nil {
			return err
		}

		// - Insert a new user_promotion
		_, _, err = tx.UserPromotionRepo().UpsertOne(ctx, &entity.UserPromotion{
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

	// Release the lock
	realeased, err := locker.Release()
	if err != nil || !realeased {
		return false, fmt.Errorf("released: %v, failed to release lock: %v", realeased, err)
	}

	return true, nil
}

func (b *PromotionBusiness) DeletePromotion(ctx context.Context, id int64) (int64, error) {
	// Check if the promotion exists
	exists, err := b.promotionRepo.Exists(ctx, id)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, errs.ErrorNotFound
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
		err := b.txRunner.Transaction(ctx, func(tx Tx) error {
			// -- Check if the promotion is used by the user
			var userPromotion *entity.UserPromotion
			_, err := tx.UserPromotionRepo().
				FindByUserIdAndPromotionId(ctx, req.UserId, req.PromotionId)
			if errors.Is(err, errs.ErrorNotFound) {
				userPromotion = &entity.UserPromotion{
					UserId:      req.UserId,
					PromotionId: req.PromotionId,
				}
				_, _, err = tx.UserPromotionRepo().UpsertOne(ctx, userPromotion)
				if err != nil {
					return err
				}
			}
			if err != nil {
				return err
			}

			// -- Reduce the remaining_count in promotion by 1
			err = tx.PromotionRepo().UpdateRemainingCount(ctx, req.PromotionId)
			if errors.Is(err, errs.ErrorNoRowsAffected) {
				return errors.New("promotion is used out")
			}
			if err != nil {
				return err
			}

			// -- Mark as used
			err = tx.UserPromotionRepo().MarkAsUsed(ctx, &entity.MarkAsUsedReq{
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

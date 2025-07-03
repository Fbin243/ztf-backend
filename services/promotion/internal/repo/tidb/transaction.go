package tidb

import (
	"context"
	biz "ztf-backend/services/promotion/internal/business"

	"gorm.io/gorm"
)

type gormTx struct {
	*gorm.DB
}

func (t *gormTx) PromotionRepo() biz.IPromotionRepo {
	return &PromotionRepo{DB: t.DB}
}

func (t *gormTx) UserPromotionRepo() biz.IUserPromotionRepo {
	return &UserPromotionRepo{DB: t.DB}
}

type GormTxRunner struct {
	*gorm.DB
}

func NewGormTxRunner(db *gorm.DB) *GormTxRunner {
	return &GormTxRunner{db}
}

func (r *GormTxRunner) Transaction(ctx context.Context, fn func(tx biz.Tx) error) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		return fn(&gormTx{tx})
	})
	if err != nil {
		return err
	}

	return nil
}

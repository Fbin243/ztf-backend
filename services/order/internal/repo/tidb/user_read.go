package tidb

import (
	"context"
	"errors"
	"ztf-backend/services/order/internal/entity"

	"gorm.io/gorm"

	errs "ztf-backend/services/order/internal/errors"
)

func (r *UserRepo) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	if err := r.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) FindById(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	err := r.DB.WithContext(ctx).First(&user, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) FindByIds(ctx context.Context, ids []string) ([]entity.User, error) {
	var users []entity.User
	if err := r.DB.WithContext(ctx).Where("id IN (?)", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) Exists(ctx context.Context, id string) (bool, error) {
	var count int64
	var user entity.User
	err := r.DB.WithContext(ctx).Model(&user).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepo) FindUsersByIds(ctx context.Context, ids []string) ([]entity.User, error) {
	var users []entity.User
	if err := r.DB.WithContext(ctx).Where("id IN (?)", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

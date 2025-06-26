package base

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
	errs "ztf-backend/pkg/errors"
)

type IBaseRepo[E IBaseEntity] interface {
	FindAll(ctx context.Context) ([]E, error)
	FindById(ctx context.Context, id string) (*E, error)
	FindByIds(ctx context.Context, ids []string) ([]E, error)
	InsertOne(ctx context.Context, entity *E) (string, error)
	UpdateOne(ctx context.Context, entity *E) (string, error)
	DeleteOne(ctx context.Context, id string) (string, error)
	Exists(ctx context.Context, id string) (bool, error)
}

type BaseRepo[E IBaseEntity] struct {
	*gorm.DB
}

func NewBaseRepo[E IBaseEntity](db *gorm.DB) *BaseRepo[E] {
	return &BaseRepo[E]{db}
}

func (r *BaseRepo[E]) FindAll(ctx context.Context) ([]E, error) {
	var entities []E
	if err := r.DB.WithContext(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseRepo[E]) FindById(ctx context.Context, id string) (*E, error) {
	var e E
	err := r.DB.WithContext(ctx).First(&e, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *BaseRepo[E]) FindByIds(ctx context.Context, ids []string) ([]E, error) {
	var entities []E
	if err := r.DB.WithContext(ctx).Where("id IN (?)", ids).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseRepo[E]) InsertOne(ctx context.Context, entity *E) (string, error) {
	lo.FromPtr(entity).SetId(uuid.New().String())
	if err := r.DB.WithContext(ctx).Create(entity).Error; err != nil {
		return "", err
	}
	return lo.FromPtr(entity).GetId(), nil
}

func (r *BaseRepo[E]) UpdateOne(ctx context.Context, entity *E) (string, error) {
	if err := r.DB.WithContext(ctx).Save(entity).Error; err != nil {
		return "", err
	}
	return lo.FromPtr(entity).GetId(), nil
}

func (r *BaseRepo[E]) DeleteOne(ctx context.Context, id string) (string, error) {
	var e E
	err := r.DB.WithContext(ctx).Delete(&e, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errs.ErrorNotFound
	}
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *BaseRepo[E]) Exists(ctx context.Context, id string) (bool, error) {
	var count int64
	var e E
	err := r.DB.WithContext(ctx).Model(&e).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

package base

import (
	"errors"
	"github.com/google/uuid"
	errs "ztf-backend/shared/errors"

	"github.com/samber/lo"
	"gorm.io/gorm"
	"ztf-backend/shared/pkg/db"
)

type BaseRepo[E IBaseEntity] struct {
	*gorm.DB
}

func NewBaseRepo[E IBaseEntity]() *BaseRepo[E] {
	return &BaseRepo[E]{db.GetDB()}
}

func (r *BaseRepo[E]) FindAll() ([]E, error) {
	var entities []E
	if err := r.DB.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseRepo[E]) FindById(id string) (*E, error) {
	var e E
	err := r.DB.First(&e, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *BaseRepo[E]) FindByIds(ids []string) ([]E, error) {
	var entities []E
	if err := r.DB.Where("id IN (?)", ids).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseRepo[E]) InsertOne(entity *E) (string, error) {
	lo.FromPtr(entity).SetID(uuid.New().String())
	if err := r.DB.Create(entity).Error; err != nil {
		return "", err
	}
	return lo.FromPtr(entity).GetID(), nil
}

func (r *BaseRepo[E]) UpdateOne(entity *E) (string, error) {
	if err := r.DB.Save(entity).Error; err != nil {
		return "", err
	}
	return lo.FromPtr(entity).GetID(), nil
}

func (r *BaseRepo[E]) DeleteOne(id string) (string, error) {
	var e E
	err := r.DB.Delete(&e, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errs.ErrorNotFound
	}
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *BaseRepo[E]) Exists(id string) (bool, error) {
	var count int64
	var e E
	err := r.DB.Model(&e).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

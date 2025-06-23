package db

import (
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type BaseRepo[E entity.IBaseEntity] struct {
	*gorm.DB
}

func NewBaseRepo[E entity.IBaseEntity]() *BaseRepo[E] {
	return &BaseRepo[E]{GetDB()}
}

func (r *BaseRepo[E]) FindAll() ([]E, error) {
	var entities []E
	if err := r.DB.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseRepo[E]) FindById(id uuid.UUID) (*E, error) {
	var entity E
	err := r.DB.First(&entity, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, utils.ErrorNotFound
	}
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepo[E]) InsertOne(entity *E) (uuid.UUID, error) {
	if err := r.DB.Create(entity).Error; err != nil {
		return uuid.Nil, err
	}
	return lo.FromPtr(entity).GetID(), nil
}

func (r *BaseRepo[E]) UpdateOne(entity *E) (uuid.UUID, error) {
	if err := r.DB.Save(entity).Error; err != nil {
		return uuid.Nil, err
	}
	return lo.FromPtr(entity).GetID(), nil
}

func (r *BaseRepo[E]) DeleteOne(id uuid.UUID) (uuid.UUID, error) {
	var entity E
	err := r.DB.Delete(&entity, id).Error
	if err == gorm.ErrRecordNotFound {
		return uuid.Nil, utils.ErrorNotFound
	}
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *BaseRepo[E]) Exists(id uuid.UUID) (bool, error) {
	var count int64
	var entity E
	err := r.DB.Model(&entity).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

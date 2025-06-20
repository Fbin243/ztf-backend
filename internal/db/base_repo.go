package db

import (
	"ztf-backend/internal/entity"

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

func (r *BaseRepo[E]) FindById(id uint) (*E, error) {
	var entity E
	if err := r.DB.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepo[E]) InsertOne(entity *E) (uint, error) {
	if err := r.DB.Create(entity).Error; err != nil {
		return 0, err
	}

	return lo.FromPtr(entity).GetID(), nil
}

func (r *BaseRepo[E]) UpdateOne(entity *E) (uint, error) {
	if err := r.DB.Save(entity).Error; err != nil {
		return 0, err
	}

	return lo.FromPtr(entity).GetID(), nil
}

func (r *BaseRepo[E]) DeleteOne(id uint) (uint, error) {
	var entity E
	if err := r.DB.First(&entity, "id = ?", id).Error; err != nil {
		return 0, err
	}

	if err := r.DB.Delete(&entity).Error; err != nil {
		return 0, err
	}

	return id, nil
}

package db

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"
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

func (r *BaseRepo[E]) FindById(id string) (*E, error) {
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
	var entity E
	err := r.DB.Delete(&entity, id).Error
	if err == gorm.ErrRecordNotFound {
		return "", utils.ErrorNotFound
	}
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *BaseRepo[E]) Exists(id string) (bool, error) {
	var count int64
	var entity E
	err := r.DB.Model(&entity).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

package repository

import (
	"gorm.io/gorm"
)

type GenericRepository[T any] struct {
	db *gorm.DB
}

func NewGenericRepository[T any](db *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{db: db}
}

func (r *GenericRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *GenericRepository[T]) FindAll(entities *[]T) error {
	return r.db.Find(entities).Error
}

func (r *GenericRepository[T]) FindByID(id any, entity *T) error {
	return r.db.First(entity, id).Error
}

func (r *GenericRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *GenericRepository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}

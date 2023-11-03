package services

import (
	"gorm.io/gorm"
)

type CrudServiceInterface[T any] interface {
	GetByID(id int) (*T, error)
	CreateObject(obj *T) (*T, error)
}

type BaseCrudService[T any] struct {
	db *gorm.DB
}

func NewBaseDBService[T any](db *gorm.DB) *BaseCrudService[T] {
	return &BaseCrudService[T]{db: db}
}

func (service *BaseCrudService[T]) GetByID(id int) (*T, error) {
	var obj T
	result := service.db.First(&obj, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &obj, nil
}

func (service *BaseCrudService[T]) CreateObject(obj *T) (*T, error) {
	result := service.db.Create(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

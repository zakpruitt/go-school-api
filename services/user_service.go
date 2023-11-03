package services

import (
	"gorm.io/gorm"
	"school-api/models"
)

type UserService struct {
	*BaseCrudService[models.User]
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		BaseCrudService: NewBaseDBService[models.User](db),
	}
}

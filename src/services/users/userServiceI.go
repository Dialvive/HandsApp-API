package services

import (
	"API/models"
	"gorm.io/gorm"
)

type UserServiceI interface {
	Save(models.User, ...string) (string, error)
	Login(models.LoginForm, func(db *gorm.DB) *gorm.DB) (string, error)
}

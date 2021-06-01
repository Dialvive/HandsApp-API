package services

import "API/models"

type UserServiceI interface {
	Save(receiver models.User, omitColumns ...string) (models.User, error)
}

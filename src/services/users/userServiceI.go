package users

import "API/models"

type userServiceI interface {
	Save() (models.User, error)
}

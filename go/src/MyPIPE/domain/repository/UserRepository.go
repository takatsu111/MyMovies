package repository

import (
	"MyPIPE/domain/model"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	FindById(id model.UserID) (*model.User, error)
	FindByToken(token model.UserToken) (*model.User, error)
	FindByEmail(email model.UserEmail) (*model.User, error)
	SetUser(*model.User) error
}
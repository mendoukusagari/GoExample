package repository

import (
	"github.com/seishino/go-example/core/domain"
)

type UserRepository interface {
	Add(user domain.User) (createdUser domain.User, err error)
	FindAll() (users []domain.User, err error)
	Delete(id int) error
	Update(user domain.User) (updatedUser domain.User, err error)
	FindById(id int) (user domain.User, err error)
	FindByUsername(username string) (domain.User, error)
}

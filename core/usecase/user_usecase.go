package usecase

import "github.com/seishino/go-example/core/domain"

type UserUsecase interface {
	Add(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	GetAll() ([]domain.User, error)
	GetById(id int) (domain.User, error)
	Login(user domain.User) (domain.User, error)
	DeleteById(id int) error
}

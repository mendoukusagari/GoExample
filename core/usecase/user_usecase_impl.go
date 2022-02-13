package usecase

import (
	"github.com/seishino/go-example/core/domain"
	"github.com/seishino/go-example/data/mysql/repository"
	security "github.com/seishino/go-example/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	repository repository.UserRepository
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return UserUsecaseImpl{
		repository: repository,
	}
}

func (usecase UserUsecaseImpl) Add(user domain.User) (domain.User, error) {
	user.Password, _ = HashPassword(user.Password)
	createdUser, err := usecase.repository.Add(user)
	return createdUser, err
}
func (usecase UserUsecaseImpl) Update(user domain.User) (domain.User, error) {
	user.Password, _ = HashPassword(user.Password)
	updatedUser, err := usecase.repository.Update(user)
	return updatedUser, err
}
func (usecase UserUsecaseImpl) GetAll() ([]domain.User, error) {
	users, err := usecase.repository.FindAll()
	return users, err
}
func (usecase UserUsecaseImpl) GetById(id int) (domain.User, error) {
	user, err := usecase.repository.FindById(id)
	return user, err
}
func (usecase UserUsecaseImpl) DeleteById(id int) error {
	err := usecase.repository.Delete(id)
	return err
}
func (usecase UserUsecaseImpl) Login(user domain.User) (domain.User, error) {
	searchedUser, err := usecase.repository.FindByUsername(user.Username)
	CheckPasswordHash(user.Password, searchedUser.Password)
	res := CheckPasswordHash(user.Password, searchedUser.Password)
	if err == nil && res {
		jwtUtil := security.NewJWTAuthService()
		searchedUser.Token = jwtUtil.CreateToken(searchedUser)
		return searchedUser, err
	}
	return domain.User{}, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

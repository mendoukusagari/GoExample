package repository

import (
	"github.com/seishino/go-example/core/domain"
	"github.com/seishino/go-example/data/mysql/entity"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepositoryImpl{
		db: db,
	}
}

func (repository userRepositoryImpl) Add(user domain.User) (domain.User, error) {
	userModel := entity.From(user)
	err := repository.db.Create(&userModel).Error
	return userModel.FromThis(), err
}
func (repository userRepositoryImpl) Delete(id int) error {
	err := repository.db.Delete(&entity.User{}, id).Error
	return err
}
func (repository userRepositoryImpl) FindById(id int) (user domain.User, err error) {
	err = repository.db.First(&user, id).Error
	return user, err
}
func (repository userRepositoryImpl) FindAll() ([]domain.User, error) {
	var userModels []entity.User
	err := repository.db.Find(&userModels).Error
	return entity.FromListThis(userModels), err
}
func (repository userRepositoryImpl) Update(user domain.User) (domain.User, error) {
	err := repository.db.Updates(&user).Error
	return user, err
}
func (repository userRepositoryImpl) FindByUsername(username string) (user domain.User, err error) {
	err = repository.db.Where("username = ?", username).First(&user).Error
	return user, err
}

package usecase

import (
	"testing"

	"github.com/seishino/go-example/config"
	"github.com/seishino/go-example/core/domain"
	"github.com/seishino/go-example/data/mysql/repository"
	"github.com/stretchr/testify/assert"
)

func CreateUsecase() UserUsecase {
	repo := repository.NewUserRepository(config.GetDB())
	return NewUserUsecase(repo)
}

func Test_Add(t *testing.T) {
	usecase := CreateUsecase()
	user := domain.User{
		Username: "test5",
		Password: "test",
	}
	_, err := usecase.Add(user)
	assert.Nil(t, err)
}
func Test_FindAll(t *testing.T) {
	usecase := CreateUsecase()
	_, err := usecase.GetAll()
	assert.Nil(t, err)

}
func Test_FindById(t *testing.T) {
	usecase := CreateUsecase()
	user, err := usecase.GetById(14)
	assert.Nil(t, err)
	assert.NotNil(t, user)

}
func Test_Delete(t *testing.T) {
	usecase := CreateUsecase()
	user := domain.User{
		Username: "test",
		Password: "test",
	}
	user, err := usecase.Add(user)

	err = usecase.DeleteById(int(user.ID))
	assert.Nil(t, err)

}
func Test_Update(t *testing.T) {
	usecase := CreateUsecase()
	user := domain.User{
		Username: "test",
		Password: "test",
	}
	user, err := usecase.Add(user)

	user.Password = "Password"

	updatedUser, err := usecase.Update(user)
	assert.Nil(t, err)
	user, err = usecase.GetById(int(updatedUser.ID))
	assert.Equal(t, user.ID, updatedUser.ID)
	assert.Equal(t, user.Username, updatedUser.Username)
	assert.Equal(t, user.Password, updatedUser.Password)
	assert.Equal(t, user.Token, updatedUser.Token)
}
func Test_Login(t *testing.T) {
	usecase := CreateUsecase()
	user := domain.User{
		Username: "User",
		Password: "Pass",
	}
	searchedUser, _ := usecase.Login(user)
	assert.NotEmpty(t, searchedUser.Username)
}

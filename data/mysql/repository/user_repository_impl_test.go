package repository

import (
	"testing"

	"github.com/seishino/go-example/config"
	"github.com/seishino/go-example/core/domain"
	"github.com/stretchr/testify/assert"
)

func CreateRepository() UserRepository {
	return NewUserRepository(config.GetDB())
}

func Test_Add(t *testing.T) {
	repo := CreateRepository()
	user := domain.User{
		Username: "test",
		Password: "test",
	}
	_, err := repo.Add(user)
	assert.Nil(t, err)
}
func Test_FindAll(t *testing.T) {
	repo := CreateRepository()
	_, err := repo.FindAll()
	assert.Nil(t, err)

}
func Test_FindById(t *testing.T) {
	repo := CreateRepository()
	user, err := repo.FindById(14)
	assert.Nil(t, err)
	assert.NotNil(t, user)

}
func Test_Delete(t *testing.T) {
	repo := CreateRepository()

	user := domain.User{
		Username: "test",
		Password: "test",
	}
	user, err := repo.Add(user)

	err = repo.Delete(int(user.ID))
	assert.Nil(t, err)

}
func Test_Update(t *testing.T) {
	repo := CreateRepository()
	user := domain.User{
		Username: "test",
		Password: "test",
	}
	user, err := repo.Add(user)

	user.Password = "Password"

	updatedUser, err := repo.Update(user)
	assert.Nil(t, err)
	user, err = repo.FindById(int(updatedUser.ID))
	assert.Equal(t, user.ID, updatedUser.ID)
	assert.Equal(t, user.Username, updatedUser.Username)
	assert.Equal(t, user.Password, updatedUser.Password)
	assert.Equal(t, user.Token, updatedUser.Token)
}
func Test_FindByUsername(t *testing.T) {
	repository := CreateRepository()
	searchedUser, _ := repository.FindByUsername("test")
	assert.NotEmpty(t, searchedUser.Username)
}

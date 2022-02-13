package entity

import "github.com/seishino/go-example/core/domain"

type User struct {
	ID       uint
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Token    string `gorm:"column:token"`
}

func From(user domain.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Token:    user.Token,
	}
}
func FromList(users []domain.User) []User {
	var userModels []User
	for _, user := range users {
		userModels = append(userModels, From(user))
	}
	return userModels
}
func (user *User) FromThis() domain.User {
	return domain.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Token:    user.Token,
	}
}
func FromListThis(userModels []User) []domain.User {
	var users []domain.User
	for _, user := range userModels {
		users = append(users, user.FromThis())
	}
	return users
}

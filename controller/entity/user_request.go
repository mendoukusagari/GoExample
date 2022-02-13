package entity

import "github.com/seishino/go-example/core/domain"

type UserRequest struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ToUserRequest(user domain.User) UserRequest {
	return UserRequest{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}
}
func ToUserRequests(users []domain.User) []UserRequest {
	var userRequests []UserRequest
	for _, user := range users {
		userRequests = append(userRequests, ToUserRequest(user))
	}
	return userRequests
}
func (user *UserRequest) FromUserRequest() domain.User {
	return domain.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}
}
func FromUserRequests(userRequests []UserRequest) []domain.User {
	var users []domain.User
	for _, user := range userRequests {
		users = append(users, user.FromUserRequest())
	}
	return users
}

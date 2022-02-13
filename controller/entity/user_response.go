package entity

import "github.com/seishino/go-example/core/domain"

type UserResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		Username: user.Username,
		Token:    user.Token,
	}
}
func ToUserResponses(users []domain.User) []UserResponse {
	var userResponse []UserResponse
	for _, user := range users {
		userResponse = append(userResponse, ToUserResponse(user))
	}
	return userResponse
}
func (user *UserResponse) FromUserResponse() domain.User {
	return domain.User{
		Username: user.Username,
		Token:    user.Token,
	}
}
func FromUserResponses(userResponses []UserResponse) []domain.User {
	var users []domain.User
	for _, user := range userResponses {
		users = append(users, user.FromUserResponse())
	}
	return users
}

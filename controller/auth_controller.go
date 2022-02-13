package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seishino/go-example/controller/entity"
	"github.com/seishino/go-example/core/usecase"
)

type AuthController struct {
	userUsecase usecase.UserUsecase
}

func NewAuthController(usecase usecase.UserUsecase) AuthController {
	return AuthController{
		userUsecase: usecase,
	}
}

func (userController *AuthController) CreateUser(c *gin.Context) {
	var userRequest entity.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Value should be JSON"})
		return
	}
	user, err := userController.userUsecase.Add(userRequest.FromUserRequest())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payload": entity.ToUserRequest(user)})
}

func (userController *AuthController) LoginUser(c *gin.Context) {
	var userRequest entity.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Value should be JSON"})
		return
	}
	user, err := userController.userUsecase.Login(userRequest.FromUserRequest())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Login is failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payload": entity.ToUserResponse(user)})
}

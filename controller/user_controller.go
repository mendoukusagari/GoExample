package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/seishino/go-example/controller/entity"
	"github.com/seishino/go-example/core/usecase"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) UserController {
	return UserController{
		userUsecase: usecase,
	}
}

func (userController *UserController) GetById(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	user, err := userController.userUsecase.GetById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payload": entity.ToUserRequest(user)})
}

func (userController *UserController) GetAllUser(c *gin.Context) {
	users, err := userController.userUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payload": entity.ToUserRequests(users)})
}

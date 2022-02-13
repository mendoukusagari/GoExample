package router

import (
	"github.com/seishino/go-example/controller"
	"github.com/seishino/go-example/lib"
)

type AuthRouter struct {
	handler        lib.RequestHandler
	authController controller.AuthController
}

func (authRouter AuthRouter) Setup() {
	api := authRouter.handler.Gin.Group("/api/auth")
	api.POST("/register", authRouter.authController.CreateUser)
	api.POST("/login", authRouter.authController.LoginUser)

}

func NewAuthRouter(
	handler lib.RequestHandler,
	authController controller.AuthController,
) AuthRouter {
	return AuthRouter{
		handler:        handler,
		authController: authController,
	}
}

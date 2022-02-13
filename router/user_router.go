package router

import (
	"github.com/seishino/go-example/controller"
	"github.com/seishino/go-example/lib"
	"github.com/seishino/go-example/middleware"
)

type UserRouter struct {
	handler        lib.RequestHandler
	userController controller.UserController
	jwtMiddleware  middleware.JwtMiddleware
}

func (userRouter UserRouter) Setup() {
	api := userRouter.handler.Gin.Group("/api/user").Use(userRouter.jwtMiddleware.Handler())
	api.GET("/:id", userRouter.userController.GetById)
	api.GET("/", userRouter.userController.GetAllUser)

}

func NewUserRouter(
	handler lib.RequestHandler,
	userController controller.UserController,
	jwtMiddleware middleware.JwtMiddleware,
) UserRouter {
	return UserRouter{
		handler:        handler,
		userController: userController,
		jwtMiddleware:  jwtMiddleware,
	}
}

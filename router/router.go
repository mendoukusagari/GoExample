package router

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRouter),
	fx.Provide(NewRoutes),
	fx.Provide(NewAuthRouter),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	userRouter UserRouter,
	authRouter AuthRouter,
) Routes {
	return Routes{
		userRouter,
		authRouter,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}

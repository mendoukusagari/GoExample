package bootstrap

import (
	"github.com/seishino/go-example/config"
	"github.com/seishino/go-example/controller"
	"github.com/seishino/go-example/core/usecase"
	"github.com/seishino/go-example/data/mysql/repository"
	"github.com/seishino/go-example/lib"
	"github.com/seishino/go-example/middleware"
	"github.com/seishino/go-example/router"
	"github.com/seishino/go-example/utils"
	"go.uber.org/fx"
)

var Module = fx.Options(
	router.Module,
	repository.Module,
	usecase.Module,
	controller.Module,
	lib.Module,
	config.Module,
	middleware.Module,
	utils.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	handler lib.RequestHandler,
	routes router.Routes,
) {
	routes.Setup()
	handler.Gin.Run(":8002")
}

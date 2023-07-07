//go:build wireinject
// +build wireinject

package di

import (
	"example.com/gin_wire/app/controllers"
	"example.com/gin_wire/app/repositories"
	"example.com/gin_wire/app/usecases"
	"github.com/google/wire"
)

func InitializeServer() *Router {
	wire.Build(
		NewRouter,
		controllers.NewSampleController,
		repositories.NewSampleRepository,
		usecases.NewSampleUsecase,
	)
	return nil
}

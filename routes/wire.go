//go:build wireinject
// +build wireinject

package routes

import (
	"github.com/google/wire"
	"golang_sample/controllers"
	"golang_sample/repository"
	"golang_sample/services"
)

func InitializeUserController() *controllers.UserController {
	wire.Build(controllers.NewUserController, services.NewUserService, repository.NewUserRepo)
	return &controllers.UserController{}
}

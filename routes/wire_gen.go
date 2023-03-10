// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package routes

import (
	"golang_sample/controllers"
	"golang_sample/repository"
	"golang_sample/services"
)

// Injectors from wire.go:

func InitializeUserController() *controllers.UserController {
	userRepo := repository.NewUserRepo()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	return userController
}

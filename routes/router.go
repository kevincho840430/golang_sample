package routes

import (
	// "golang_sample/controllers"
	// "golang_sample/services"
	// "golang_sample/repository"
	"github.com/gin-gonic/gin"
)

func SetRouter(g *gin.Engine) *gin.Engine {
	// userRepo := repository.NewUserRepo()
	// userService := services.NewUserService(userRepo)
	// userCtrl := controllers.NewUserController(userService)

	userGroup := g.Group("/user")
	{
		//增加用户User
		userGroup.GET("/getAll", InitializeUserController().GetUser)
		userGroup.POST("/add", InitializeUserController().Create)
		userGroup.POST("/edit", InitializeUserController().Update)
		userGroup.POST("/delete", InitializeUserController().Delete)
	}

	return g
}

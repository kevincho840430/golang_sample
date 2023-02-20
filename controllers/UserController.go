package controllers

import (
	"github.com/gin-gonic/gin"
	"golang_sample/models"
	"golang_sample/services"
	"net/http"
	//    "fmt"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: *userService}
}

func (u *UserController) GetUser(c *gin.Context) {
	var userList []*models.User

	// c.BindJSON(&user)

	userList, err := u.UserService.GetUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": userList,
		})
	}
}

func (u *UserController) Create(c *gin.Context) {
	err := u.UserService.Create(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func (u *UserController) Update(c *gin.Context) {
	err := u.UserService.Update(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func (u *UserController) Delete(c *gin.Context) {
	err := u.UserService.Delete(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

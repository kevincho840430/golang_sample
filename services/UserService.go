package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/liudng/godump"
	"golang_sample/models"
	"golang_sample/repository"
	"net/http"
	// "errors"
)

type UserService struct {
	UserRepo repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{UserRepo: *userRepo}
}

// get user
func (u *UserService) GetUser(c *gin.Context) (userList []*models.User, err error) {
	// error := errors.New("建立錯誤訊息")
	result, error := u.UserRepo.GetAllUser()
	return result, error
}

//	func RegisterOneUser(account string, password string, email string) error {
//		if !CheckOneUser(account) {
//			return fmt.Errorf("User exists.")
//		}
//		user := model.User{
//			Account: account,
//			Password: password,
//			Email: email,
//		}
//		insertErr := dao.SqlSession.Model(&model.User{}).Create(&user).Error
//		return insertErr
//	}
func (u *UserService) Create(c *gin.Context) (err error) {
	// error := errors.New("建立錯誤訊息")
	user := models.User{}
	godump.Dump(c.PostForm("Name"))
	c.JSON(http.StatusOK, c.Request.PostForm)

	c.BindJSON(&user)
	user.UserId = uuid.New().String()
	user.Name = c.Request.FormValue("Name")
	user.Password = c.Request.FormValue("Password")
	user.Email = c.Request.FormValue("Email")

	godump.Dump(user)
	error := u.UserRepo.Create(&user)

	return error
}

func (u *UserService) Update(c *gin.Context) (err error) {
	// error := errors.New("建立錯誤訊息")
	user := models.User{}
	godump.Dump(c.PostForm("Name"))
	c.BindJSON(&user)

	user.UserId = c.Request.FormValue("UserId")
	user.Name = c.Request.FormValue("Name")
	user.Password = c.Request.FormValue("Password")
	user.Email = c.Request.FormValue("Email")

	error := u.UserRepo.Update(&user)

	return error
}

func (u *UserService) Delete(c *gin.Context) (err error) {
	// error := errors.New("建立錯誤訊息")
	user := models.User{}
	godump.Dump(c.PostForm("UserId"))
	c.BindJSON(&user)
	user.UserId = c.Request.FormValue("UserId")
	error := u.UserRepo.Delete(&user)
	return error
}

package repository

import (
	"golang_sample/models"
	"golang_sample/pkg"
	// "fmt"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (u *UserRepo) GetAllUser() (userModel []*models.User, err error) {
	user := []*models.User{}
	error := pkg.SqlSession.Find(&user).Error

	return user, error
}

func (u *UserRepo) Create(user *models.User) (err error) {
	error := pkg.SqlSession.Create(&user).Error
	return error
}

func (u *UserRepo) Update(user *models.User) (err error) {
	error := pkg.SqlSession.Save(&user).Error
	return error
}

func (u *UserRepo) Delete(user *models.User) (err error) {
	error := pkg.SqlSession.Where("user_id=?", user.UserId).Delete(&models.User{}).Error
	return error
}

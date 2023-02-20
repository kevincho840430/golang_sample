package main

import (
	"golang_sample/models"
	"golang_sample/pkg"
	"gorm.io/gorm"
)

func main() {
	pkg.InitMySql()
	user(pkg.SqlSession)
}

func user(db *gorm.DB) {
	db.AutoMigrate(&models.User{})

	//刪除的時候再用
	//genderExisted := db.Migrator().HasColumn(&models.User{}, "Gender")
	//if genderExisted {
	//	db.Migrator().DropColumn(&models.User{}, "Gender")
	//}
}

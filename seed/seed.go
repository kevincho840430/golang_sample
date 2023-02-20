package main

import (
	"golang_sample/models"
	"golang_sample/pkg"

	"gorm.io/gorm"
)

func main() {
	pkg.InitMySql()
	seedUser(pkg.SqlSession)
}

func seedUser(db *gorm.DB) {
	users := []models.User{
		{
			UserId:   "4c494980-1b8f-489f-91c9-63d268afa5a2",
			Name:     "Kevin",
			Password: "kevinSoGood",
			Email:    "8286kevin@gmail.com",
		},
		{
			UserId:   "5a494980-1b8f-489f-91c9-63d268afa5a2",
			Name:     "Frank",
			Password: "FrankSoGood",
			Email:    "8833frank@gmail.com",
		},
		{
			UserId:   "6b494980-1b8f-489f-91c9-63d268afa5a2",
			Name:     "Nick",
			Password: "NickSoGood",
			Email:    "8888nick@gmail.com",
		},
	}
	for _, u := range users {
		db.Create(&u)
	}
}

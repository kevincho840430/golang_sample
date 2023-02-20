package main

import (
	"github.com/gin-gonic/gin"
	"golang_sample/routes"
	// "fmt"
	"golang_sample/pkg"
)

func main() {
	err, _ := pkg.InitMySql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	// defer pkg.Close()
	router := gin.Default()
	routes.SetRouter(router)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "service is running",
		})
	})

	router.Run(":8000")
}

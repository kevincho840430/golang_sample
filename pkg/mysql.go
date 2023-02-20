package pkg

import (
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

var SqlSession *gorm.DB

type conf struct {
	Url string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName string `yaml:"dbname"`
	Port string `yaml:"post"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("./configs/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func InitMySql()(err error , db *gorm.DB)  {
	var c conf
	conf:=c.getConf()
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)
	SqlSession,err =gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		panic(err)
	}
	return err,SqlSession
}

// func Close()  {
// 	SqlSession.Close()
// }

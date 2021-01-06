package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
)

type conf struct {
	DbConnect string   `yaml:"db_connect"`
	Port      int      `yaml:"port"`
	SecretKey string   `yaml:"secret_key"`
	WriteDb   string   `yaml:"write_db"`
	ReadDb    []string `yaml:"read_db"`
	TokenTime int64    `yaml:"token_time"`
	OrderTime int      `yaml:"order_time"`
	ImgServer string   `yaml:"img_server"`
}

var c conf

func GetConf() conf {
	yamlFile, err := ioutil.ReadFile("Config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func GetBindPort() int {
	return c.Port
}

func GetJWTSecret() string {
	return c.SecretKey
}
func GetOneDB() *gorm.DB {
	db, err := gorm.Open("mysql", c.WriteDb)
	if err != nil {
		panic(err)
	}
	db.DB().SetConnMaxIdleTime(5)
	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)
	return db
}

func GetRandomSlaveDB() *gorm.DB {
	choiceDbIdx := rand.Intn(len(c.ReadDb))
	db, err := gorm.Open("mysql", c.ReadDb[choiceDbIdx])
	if err != nil {
		panic(err)
	}
	db.DB().SetConnMaxIdleTime(5)
	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)
	return db
}

func GetTokenValidTime() int64 {
	return c.TokenTime
}

func GetOrderTime() int {
	return c.OrderTime
}

func GetImgServer() string {
	return c.ImgServer
}

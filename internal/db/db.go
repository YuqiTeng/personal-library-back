package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	userName string
	password string
	host     string
	dbName   string
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("internal/")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	userName = viper.GetString("DB_USERNAME")
	password = viper.GetString("DB_PASSWORD")
	host 	 = viper.GetString("DB_HOST")
	dbName   = viper.GetString("DB_NAME")

	dbHost := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, host, dbName)
	db, err = gorm.Open("mysql", dbHost)
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

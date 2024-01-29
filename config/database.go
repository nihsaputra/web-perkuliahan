package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectionDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/perkuliahan_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	log.Println("connected to database success")
}

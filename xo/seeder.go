package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"./models"
)

func dbConnect() *gorm.DB {
	dbUser := "root"
	dbPass := "test"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "orange"

	dbPath := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbPath)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	db := dbConnect()
	defer db.Close()

	user := models.User{
		Email: "user@example.com",
	}

	db.Create(&user)
}

package main

import (
	"../models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	db, err := gorm.Open("mysql", "root:@/gorm?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	db.CreateTable(&models.User{})
}

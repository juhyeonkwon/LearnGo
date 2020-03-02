package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
type UserModel struct {
	Id      int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
}

func gorm_SQLparameters() {

	db, err := gorm.Open("mysql", "golang:wngus123@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()

	if err != nil {
		log.Println("Error!")
	} else {
		log.Println("Connected!")
	}

	db.CreateTable(&UserModel{})

}
*/

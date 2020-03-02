package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //You could import dialect
)

//User schema
type User struct {
	ID       int
	Username string
}

func main() {

	//connection
	db, err := gorm.Open("mysql", "golang:wngus123@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()

	if err != nil {
		log.Println("Connection Failed")
	} else {
		log.Println("Connected!")
	}
	//

	db.CreateTable(&User{})

}

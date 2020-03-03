package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserModel struct {
	Id      int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
}

func main() {

	db, err := gorm.Open("mysql", "golang:wngus123@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()

	if err != nil {
		log.Println("ERROR")
	} else {
		log.Println("CONNECTED!")
	}

	//간단한 insert
	//user := &UserModel{Name: "hong", Address: "JinJu"}
	//db.Create(user)

	//여러개 insert
	/*var users []UserModel = []UserModel{
		UserModel{Name: "kim", Address: "Seoul"},
		UserModel{Name: "kwon", Address: "jinju"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	*/

	//Update
	user := &UserModel{Name: "kim", Address: "Seoul"}

	//1
	db.Find(&user)          //select
	user.Address = "munsan" //edit
	db.Save(&user)          //save

	//2
	/*db.Model(&user).Update("name", "kwon")

	db.Model(&user).Updates(
		map[string]interface{}{
			"name":    "aa",
			"Address": "moonsan",
		})*/

}

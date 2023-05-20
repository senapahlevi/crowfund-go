package main

import (
	"crowfund/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowfund?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())

	}
	fmt.Println("good connect ya")

	// var users []user.User

	// db.Find(&users)
	// for _, user := range users {

	// 	fmt.Println("hello", user.Name)
	// 	fmt.Println("hello email", user.Email)

	// }
	userRepository := user.NewRepository(db)
	user := user.User{
		Name: "coba dari golang",
	}
	userRepository.Save(user)
}

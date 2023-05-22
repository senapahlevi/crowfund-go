package main

import (
	"crowfund/handler"
	"crowfund/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	userByEmail, err := userRepository.FindByEmail("sena3il@gmaisl.com")
	if err != nil {
		fmt.Println(err.Error())

	}
	if userByEmail.ID == 0 {
		fmt.Println("maaf ga ketemu")
	}

	fmt.Println(userByEmail.Name, "hello")

	router := gin.Default()
	api := router.Group("/api/v1/")
	api.POST("/users", userHandler.RegisterUser)
	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "sena service"
	// userInput.Occupation = "backend engineer"
	// userInput.Password = "123456789"
	// userService.RegisterUser(userInput)
	// user := user.User{ udah ga make ini yang gw komen
	// 	Name: "coba dari golang",
	// }
	// userRepository.Save(user)
}

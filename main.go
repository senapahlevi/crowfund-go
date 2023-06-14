package main

import (
	"crowfund/auth"
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
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)
	// fmt.Println(authService.GenerateToken(1000)) // print token
	// userByEmail, err := userRepository.FindByEmail("sena3il@gmaisl.com")
	// if err != nil {
	// 	fmt.Println(err.Error())

	// }
	//test token

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyNn0.1r0MdmrpNNBygCcchvfjGKyvqX7-2cAH9D8xyIHXH_8")
	if err != nil {
		fmt.Println("error tokenee hello ")

	}
	if token.Valid {
		fmt.Println("hello neee token valid")
	} else {
		fmt.Println("hello neee token ga validdd!")

	}
	// if userByEmail.ID == 0 {
	// 	fmt.Println("maaf ga ketemu")
	// }

	// input := user.LoginInput{
	// 	Email:    "sena1@gmail.com",
	// 	Password: "123456789",
	// }
	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println(err.Error(), "login terjadi kesalahan ")
	// }
	// fmt.Println(user.Name, "login berhasil hello nama")

	// fmt.Println(userByEmail.Name, "hello")

	router := gin.Default()
	api := router.Group("/api/v1/")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckAvailabilityEmail)
	api.POST("/avatars", userHandler.UploadAvatar)
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

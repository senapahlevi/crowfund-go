package handler

import (
	"crowfund/helper"
	"crowfund/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct diatas kita passing sebagai parameter service
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input) //ini supaya kebaca dari format .json ke struct

	if err != nil {
		errors := user.FormatError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("data Register failed ", http.StatusUnprocessableEntity, "error", errorMessage) // nil karena ga ada token

		c.JSON(http.StatusUnprocessableEntity, response) //jika nil maka response json postman null
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("user Registered failed ", http.StatusBadRequest, "error", nil) // nil karena ga ada token
		c.JSON(http.StatusBadRequest, response)
	}

	fmt.Println("hello woy")
	formatter := user.FormatUser(newUser, "initoken")
	response := helper.APIResponse("Account success register", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {

	//user melakukan input(email & password)
	//input ditangkap handler
	//mapping dari input user ke input struct
	//input struct passing service
	//di service mencari bantuan dgn bantuan repository user dengan email x
	//mencocokan passowrd
}

package handler

import (
	"crowfund/helper"
	"crowfund/user"
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
		errors := user.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

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
	var input user.LoginInput

	err := c.ShouldBindJSON(&input) //ini supaya kebaca dari format .json ke struct

	if err != nil {
		errors := user.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "failed", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, "initoken")
	response := helper.APIResponse("Login success yeay", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {

}

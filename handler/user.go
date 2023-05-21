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
		response := helper.APIResponse("Registered failed ", http.StatusBadRequest, "error", nil) // nil karena ga ada token

		c.JSON(http.StatusBadGateway, response) //jika nil maka response json postman null
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Registered failed ", http.StatusBadRequest, "error", nil) // nil karena ga ada token

		c.JSON(http.StatusBadRequest, response)
	}
	formatter := user.APIFormatter(newUser, "initoken")
	response := helper.APIResponse("Account success registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

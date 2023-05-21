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
		c.JSON(http.StatusBadGateway, nil)
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	response := helper.APIResponse("Account success registered", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)

}

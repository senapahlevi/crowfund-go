package handler

import (
	"crowfund/auth"
	"crowfund/helper"
	"crowfund/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, token)
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

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Login success yeay", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckAvailabilityEmail(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := user.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Email check Failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	isEmailAvailable, err := h.userService.IsEmailAvailableInput(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Email checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{"is_available": isEmailAvailable}
	metaMessage := "email already register"
	if isEmailAvailable {
		metaMessage = "email is available"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// input dari user
	// simpan dari gambar nya folder "images/"
	// di service kita panggil repo
	// JWT (sementara hardcode jadi user yang login ID =1 )
	// repo diambil dari user yang login ID =1
	// repo update data user simpan lokasi file
	file, err := c.FormFile("avatar")
	if err != nil {

		data := gin.H{"is_upload": false}

		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//harusnya dapat dari jwt tapi sementara 1 dulu
	// userID := 1

	currentUser := c.MustGet("currentUser").(user.User) //balikan default interface karena c.Set("currentUser", user) //typenya user User //jadi user saat ini login dan dapet informasi payloadnya kayak ID user, name ,dll

	userID := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename) //bikin folder local di golang

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveUserAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_upload": true}
	response := helper.APIResponse("Avatar successfully upload image", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
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
	//map input dari user ke struct RegisterUserinput
	//struct diatas kita passing sbg parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Register Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {

		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokenexample")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	//user memasukkan input (email & password)
	//input ditangkap user
	//mapping dari input user ke input struct
	//input struct passing service
	//di service , mencari dgn bantuan repository user dgn email X
	//validasi password
	var input user.LoginInput
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Login Failed", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	userLogged, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	formatter := user.FormatUser(userLogged, "tokenexample")
	response := helper.APIResponse("Login Success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

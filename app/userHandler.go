package app

import (
	"fmt"
	"miniproject_products/domain"
	"miniproject_products/dto"
	"miniproject_products/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userDTO struct {
	ID    int    `json:"id"`
	Username string `json:"username"`
	Token string `json:"token"`
}

func FormatuserDTO(user domain.Users, token string) userDTO {
	userDTO := userDTO{}
	userDTO.ID = user.ID
	userDTO.Username = user.Username
	userDTO.Token = token
	return userDTO
}

type UserHandler struct {
	service service.UserService
	authService service.AuthService

}

func (ch *UserHandler)registerUser (c *gin.Context){
	var request dto.UserRequest
	err := c.ShouldBindJSON(&request)

	users, _ := ch.service.RegisterUser(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ch *UserHandler) loginUser (c *gin.Context){
	var request dto.Login
	err := c.ShouldBindJSON(&request)

	users, _ := ch.service.LoginUser(request)
	if err != nil{
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	

	token, err := ch.authService.GenerateToken(users.ID)
	if err != nil{
		fmt.Println("Test", token)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	userDTO := FormatuserDTO(users, token)
	c.JSON(http.StatusOK, userDTO)
}
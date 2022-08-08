package app

import (
	"miniproject_products/dto"
	"miniproject_products/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
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
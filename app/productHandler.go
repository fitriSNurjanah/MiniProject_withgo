package app

import (
	"miniproject_products/domain"
	"miniproject_products/dto"
	"miniproject_products/helper"
	"miniproject_products/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductService
}

func getCurrentUserJWT(c *gin.Context) int {
	currentUser := c.MustGet("currentUser").(domain.Users)
	return currentUser.ID
}


func (ch *ProductHandler) getAllProduct(c*gin.Context){
	
	userID := getCurrentUserJWT(c)
	pagination := helper.GeneratePaginationRequest(c)
	products, err := ch.service.GetAllProduct(*pagination, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, products)
}


func (ch *ProductHandler) getProductID(c *gin.Context) {

	id := c.Param("id")
	newId, _ := strconv.Atoi(id)
	products, err := ch.service.GetProductID(newId)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, products)
}

func(ch * ProductHandler) createProduct(c *gin.Context){
	var request dto.ProductRequest
	err := c.ShouldBindJSON(&request)

	products, _:= ch.service.CreateProduct(request)
	if err!= nil{
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, products)
}


func (ch *ProductHandler)UpdateProduct(c *gin.Context){
	id := c.Param("id")
	newID, _ :=strconv.Atoi(id)
	var request dto.ProductRequest
	err := c.ShouldBindJSON(&request)
	products,_ := ch.service.UpdateProduct(newID, request)
	if err != nil{
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, products)
}


func (ch *ProductHandler) DeleteProduct(c *gin.Context) {

	id := c.Param("id")
	newId, _ := strconv.Atoi(id)
	Products, err := ch.service.DeleteProduct(newId)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, Products)
}
package app

import (
	"miniproject_products/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductService
}

func (ch *ProductHandler) getAllProduct(c*gin.Context){
	
	products, err :=ch.service.GetAllProduct()
	if err != nil{
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
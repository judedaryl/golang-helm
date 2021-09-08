package product

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type resource struct {
	service Service
}

func RegisterHandler(router *gin.RouterGroup, service Service) {
	res := resource{service}
	router.GET("/:productid", res.get)
	router.GET("/", res.getAll)
	router.POST("/", res.create)
}

func (res resource) get(c *gin.Context) {
	productid := c.Param("productid")
	product, err := res.service.Get(productid)
	if err == nil {
		c.JSON(200, product)
	}
	if err != nil {
		fmt.Print("error found")
		c.Error(err)
	}
}

func (res resource) getAll(c *gin.Context) {
	product, err := res.service.GetAll()
	if err == nil {
		c.JSON(200, product)
	}
	if err != nil {
		fmt.Print("error found")
		c.Error(err)
	}
}

func (res resource) create(c *gin.Context) {
	var req CreateProductRequest
	c.BindJSON(&req)
	product, err := res.service.Create(req)
	if err == nil {
		fmt.Print("no error found")
		c.JSON(200, product)
	}
	if err != nil {
		fmt.Print("error found")
		c.Error(err)
	}

}

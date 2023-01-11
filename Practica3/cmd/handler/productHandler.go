package handler

import (
	"practica3/internal/product"

	"github.com/gin-gonic/gin"
)

type HandlerStruct struct {
	sv product.ServiceInterface
}

func NewHandler(sv product.ServiceInterface) *HandlerStruct {
	return &HandlerStruct{sv: sv}
}

func (p *HandlerStruct) GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//request

		//process
		pr, err := p.sv.GetAllProducts()
		if err != nil {
			ctx.JSON(500, nil)
			return
		}

		// response
		ctx.JSON(200, pr)
	}
}
func (p *HandlerStruct) GetProduct() gin.HandlerFunc { //id int
	return func(ctx *gin.Context) {
		//request

		//process

		//response
	}
}
func (p *HandlerStruct) SearchProduct() gin.HandlerFunc { //priceGt float64
	return func(ctx *gin.Context) {
		//request

		//process

		//response
	}
}
func (p *HandlerStruct) PostProduct() gin.HandlerFunc { //product domain.Product
	return func(ctx *gin.Context) {
		//request

		//process

		//response
	}
}

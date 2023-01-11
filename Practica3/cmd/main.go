package main

import (
	"practica3/internal/product"

	"github.com/gin-gonic/gin"
)

func main() {
	//repo := products.NewRepository()
	//service := products.NewService(repo)
	//p := handler.NewProduct (service)
	//p.update()
	product.LoadProducts("products.json")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", product.GetAllProducts())
		products.GET("/:id", product.GetProduct())       //localhost:8080/products/100
		products.GET("/search", product.SearchProduct()) //localhost:8080/products/search?priceGt=200
		products.POST("", product.PostProduct())
	}
	r.Run(":8080")
}

// loadProducts carga los productos desde un archivo json

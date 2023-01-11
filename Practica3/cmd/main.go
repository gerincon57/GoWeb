package main

import (
	"practica3/cmd/handler"
	"practica3/internal/product"

	"github.com/gin-gonic/gin"
)

func main() {
	db := product.LoadProducts("./products.json")
	id := 4 //mirar como calculaR ID

	repo := product.NewRepository(db, id)
	service := product.NewService(repo)
	p := handler.NewHandler(service)
	//p.update()
	//product.LoadProducts("products.json")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", p.GetAllProducts())
		products.GET("/:id", p.GetProduct())       //localhost:8080/products/100
		products.GET("/search", p.SearchProduct()) //localhost:8080/products/search?priceGt=200
		products.POST("", p.PostProduct())
	}
	r.Run(":8080")
}

// loadProducts carga los productos desde un archivo json

package main

import (
	"encoding/json"
	"os"

	"github.com/bootcamp-go/Consignas-Go-Web.git/cmd/server/handler"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/product"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//importacion variable global:

	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	var productsList = []domain.Product{}
	loadProducts("../../products.json", &productsList)

	repo := product.NewRepository(productsList)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())        //localhost:8080/products
		products.GET(":id", productHandler.GetByID())    //localhost:8080/products/232
		products.GET("/search", productHandler.Search()) //localhost:8080/products/search?priceGt=100
		products.POST("", productHandler.Post())         //localhost:8080/products
		products.DELETE(":id", productHandler.Delete())  //localhost:8080/products/100
		products.PATCH(":id", productHandler.Patch())    //localhost:8080/products/200
		products.PUT(":id", productHandler.Put())        //localhost:8080/products/300
	}
	r.Run(":8080")
}

// loadProducts carga los productos desde un archivo json
func loadProducts(path string, list *[]domain.Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}

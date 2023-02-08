package main

import (
	"github.com/bootcamp-go/Consignas-Go-Web.git/cmd/server/handler"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/product"
	"github.com/bootcamp-go/Consignas-Go-Web.git/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	//"github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
	docs "github.com/practica5/cmd/server/docs"
)

// @BasePath /poducts

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /products [get]
func main() {

	//importacion variable global:

	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	//var productsList = []domain.Product{}
	//loadProducts("../../products.json", &productsList)
	productsList := storage.NewStore("../../products.json")

	repo := product.NewRepository(productsList)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/products"

	//r := gin.New()
	//r.Use(gin.Recovery())

	r.Use(handler.MiddlewareToken())

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
/*func loadProducts(path string, list *[]domain.Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}*/

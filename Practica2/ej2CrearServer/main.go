package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	CodeValue   string  `json:"code_value" binding:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

var sliceProducts = []Product{}

func main() {

	LeerJson("../products.json", &sliceProducts)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := router.Group("/products")
	{
		products.GET("", GetAllProducts())
		products.GET(":id", GetProductID()) //localhost:8080/products/5
		products.GET("/search", SearchProduct())
		products.POST("", CreateProduct())
	}
	router.Run(":8080")

}

func LeerJson(path string, slices *[]Product) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//fmt.Print("data:  ", string(data))

	err = json.Unmarshal([]byte(data), &slices)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}

func GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, sliceProducts)
	}

}
func GetProductID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		for _, product := range sliceProducts {
			if product.Id == id {
				c.JSON(http.StatusOK, product)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
	}
}
func SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("priceGt")
		priceGt, err := strconv.ParseFloat(query, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid price"})
			return
		}
		list := []Product{}
		for _, product := range sliceProducts {
			if product.Price > priceGt {
				list = append(list, product)
			}
		}
		c.JSON(200, list)
	}
}

func validarVacios(product *Product) (bool, error) {
	switch {
	case product.Name == "" || product.CodeValue == "" || product.Expiration == "":
		return false, errors.New("fields can't be empty")
	case product.Quantity <= 0 || product.Price <= 0:
		if product.Quantity <= 0 {
			return false, errors.New("quantity must be greater than 0")
		}
		if product.Price <= 0 {
			return false, errors.New("price must be greater than 0")
		}
	}
	return true, nil
}

func validarCodigo(codeValue string) bool {
	for _, product := range sliceProducts {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}

func validarExpiracion(product *Product) (bool, error) {
	dates := strings.Split(product.Expiration, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid expiration date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid expiration date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}

func CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid product"})
			return
		}
		valid, err := validarVacios(&product)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		valid, err = validarExpiracion(&product)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		valid = validarCodigo(product.CodeValue)
		if !valid {
			ctx.JSON(400, gin.H{"error": "code value already exists"})
			return
		}
		product.Id = len(sliceProducts) + 1
		sliceProducts = append(sliceProducts, product)
		ctx.JSON(201, product)
	}
}

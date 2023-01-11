package product

import (
	"encoding/json"
	"errors"
	"os"
	"practica3/internal/domain"
	"strconv"
	"strings"
)

type RepoInterface interface {

	//no se cual si va y cual no
	LoadProducts(path string)
	//ValidateEmptys(product *domain.Product) (bool, error)
	//ValidateExpiration(product *domain.Product) (bool, error)
	ValidateCodeValue(codeValue string) bool
	GetAllProducts() []domain.Product
	GetProduct(id int) (pr domain.Product, err error)
	SearchProduct(priceGt float64) []domain.Product
	PostProduct(product domain.Product) (err error)
}

type repoStruct struct {
	db     *[]domain.Product
	lastID int
}

//var productsList = []domain.Product{}

func NewRepository(db *[]domain.Product, lastID int) RepoInterface {
	return &repoStruct{db: db, lastID: lastID}
}

func (r *repoStruct) LoadProducts(path string) { //(, list *[]domain.Product)
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &r.db) //&list
	if err != nil {
		panic(err)
	}
}

// validateEmptys valida que los campos no esten vacios
// era para uno solo?
func ValidateEmptys(product *domain.Product) (bool, error) { //e(product *domain.Product)

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

// validateExpiration valida que la fecha de expiracion sea valida
// era para uno solo?
func ValidateExpiration(product *domain.Product) (bool, error) { //e(product *domain.Product)

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

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *repoStruct) ValidateCodeValue(codeValue string) bool {
	for _, product := range *r.db {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}

// GetAllProducts traer todos los productos almacenados
func (r *repoStruct) GetAllProducts() []domain.Product {
	return *r.db
	/* func(ctx *gin.Context) {
		ctx.JSON(200, productsList)
	}*/
}

// GetProduct traer un producto por id
func (r *repoStruct) GetProduct(id int) (pr domain.Product, err error) {

	/*gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}*/
	for _, product := range *r.db {
		if product.Id == id {
			//ctx.JSON(200, product)
			return product, nil
		}
	}
	//ctx.JSON(404, gin.H{"error": "product not found"})
	return domain.Product{}, errors.New("No se encontró id del product")

}

// SearchProduct traer un producto por nombre o categoria
func (r *repoStruct) SearchProduct(priceGt float64) []domain.Product {

	/*gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Query("priceGt")
		priceGt, err := strconv.ParseFloat(query, 32)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid price"})
			return
		}*/
	list := []domain.Product{}
	for _, product := range *r.db {
		if product.Price > priceGt {
			list = append(list, product)
		}
	}
	return list
}

// PostProduct crear un producto
func (r *repoStruct) PostProduct(product domain.Product) (err error) {

	/*gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product domain.Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid product"})
			return
		}*/

	valid, err := ValidateEmptys(&product)
	if !valid {
		// ctx.JSON(400, gin.H{"error": err.Error()})
		return errors.New("error desconocido")
	}
	valid, err = ValidateExpiration(&product)
	if !valid {
		//ctx.JSON(400, gin.H{"error": err.Error()})
		return errors.New("error desconocido")
	}
	valid = r.ValidateCodeValue(product.CodeValue)
	if !valid {
		//ctx.JSON(400, gin.H{"error": "code value already exists"})
		return errors.New("error desconocido")
	}
	product.Id = len(*r.db) + 1
	*r.db = append(*r.db, product)
	//ctx.JSON(201, product)
	return nil
}

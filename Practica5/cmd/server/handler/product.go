package handler

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/product"
	"github.com/bootcamp-go/Consignas-Go-Web.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	s product.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewProductHandler(s product.Service) *productHandler {
	return &productHandler{
		s: s,
	}
}

func MiddlewareToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			web.ResponseFail(c, 401, errors.New("invalid token"))
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
			//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		}
		c.Next()

	}
}

// GetAll obtiene todos los productos
func (h *productHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		products, _ := h.s.GetAll()
		web.ResponseOk(c, products, 200)
		//c.JSON(200, products)
	}
}

// GetByID obtiene un producto por su id
func (h *productHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {

		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.ResponseFail(c, 400, errors.New("invalid id"))
			//c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		product, err := h.s.GetByID(id)
		if err != nil {
			web.ResponseFail(c, 404, errors.New("product not found"))
			//c.JSON(404, gin.H{"error": "product not found"})
			return
		}
		web.ResponseOk(c, product, 200)
		//c.JSON(200, product)
	}
}

// Search busca un producto por precio mayor a un valor
func (h *productHandler) Search() gin.HandlerFunc {
	return func(c *gin.Context) {

		priceParam := c.Query("priceGt")
		price, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			web.ResponseFail(c, 400, errors.New("invalid price"))
			//c.JSON(400, gin.H{"error": "invalid price"})
			return
		}
		products, err := h.s.SearchPriceGt(price)
		if err != nil {
			web.ResponseFail(c, 404, errors.New("no products found"))
			//c.JSON(404, gin.H{"error": "no products found"})
			return
		}
		web.ResponseOk(c, products, 200)
		//c.JSON(200, products)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(product *domain.Product) (bool, error) {
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
func validateExpiration(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
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

// Post godoc
// @Summary      Create a new product
// @Description  Create a new product in repository
// @Tags         products
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Product true "Product"
// @Success      201 {object}  web.ResponseOk
// @Failure      400 {object}  web.ResponseFail
// @Failure      404 {object}  web.ResponseFail
// @Router       /products [post]
// @host localhost:8080

// Post crear un producto nuevo
func (h *productHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var product domain.Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			web.ResponseFail(ctx, 400, errors.New("invalid token"))
			//ctx.JSON(400, gin.H{"error": "invalid product"})
			return
		}
		valid, err := validateEmptys(&product)
		if !valid {
			web.ResponseFail(ctx, 400, err)
			//ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		valid, err = validateExpiration(product.Expiration)
		if !valid {
			web.ResponseFail(ctx, 400, err)
			//ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		p, err := h.s.Create(product)
		if err != nil {
			web.ResponseFail(ctx, 400, err)
			//ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		web.ResponseOk(ctx, p, 201)
	}
}

// Delete elimina un producto
func (h *productHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.ResponseFail(ctx, 400, errors.New("invalid id"))
			//ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.ResponseFail(ctx, 400, err)
			//ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		web.ResponseOk(ctx, nil, 200)
		//ctx.JSON(200, gin.H{"message": "product deleted"})
	}
}

// Put actualiza un producto
func (h *productHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.ResponseFail(ctx, 400, errors.New("invalid id"))
			//ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		var product domain.Product
		err = ctx.ShouldBindJSON(&product)
		if err != nil {
			web.ResponseFail(ctx, 400, errors.New("invalid product"))
			//ctx.JSON(400, gin.H{"error": "invalid product"})
			return
		}
		valid, err := validateEmptys(&product)
		if !valid {
			web.ResponseFail(ctx, 400, err)
			//ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		valid, err = validateExpiration(product.Expiration)
		if !valid {
			web.ResponseFail(ctx, 400, err)
			//ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		p, err := h.s.Update(id, product)
		if err != nil {
			web.ResponseFail(ctx, 409, err)
			//ctx.JSON(409, gin.H{"error": err.Error()})
			return
		}
		web.ResponseOk(ctx, p, 200)
		//ctx.JSON(200, p)
	}
}

// Patch update selected fields of a product WIP
func (h *productHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name        string  `json:"name,omitempty"`
		Quantity    int     `json:"quantity,omitempty"`
		CodeValue   string  `json:"code_value,omitempty"`
		IsPublished bool    `json:"is_published,omitempty"`
		Expiration  string  `json:"expiration,omitempty"`
		Price       float64 `json:"price,omitempty"`
	}
	return func(ctx *gin.Context) {

		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.ResponseFail(ctx, 400, errors.New("invalid id"))
			//ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			web.ResponseFail(ctx, 400, errors.New("invalid request"))
			//ctx.JSON(400, gin.H{"error": "invalid request"})
			return
		}
		update := domain.Product{
			Name:        r.Name,
			Quantity:    r.Quantity,
			CodeValue:   r.CodeValue,
			IsPublished: r.IsPublished,
			Expiration:  r.Expiration,
			Price:       r.Price,
		}
		if update.Expiration != "" {
			valid, err := validateExpiration(update.Expiration)
			if !valid {
				web.ResponseFail(ctx, 400, err)
				//ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.ResponseFail(ctx, 409, err)
			//ctx.JSON(409, gin.H{"error": err.Error()})
			return
		}
		web.ResponseOk(ctx, p, 200)
		//ctx.JSON(200, p)
	}
}

package product

import "practica3/internal/domain"

//controller

type ServiceInterface interface {

	//no se cual si va y cual no
	//LoadProducts(path string) //creo que no va
	//ValidateEmptys(product *domain.Product) (bool, error)
	//ValidateExpiration(product *domain.Product) (bool, error)
	ValidateCodeValue(codeValue string) bool
	GetAllProducts() ([]domain.Product, error)
	GetProduct(id int) (pr domain.Product, err error)
	SearchProduct(priceGt float64) ([]domain.Product, error)
	PostProduct(product domain.Product) (err error)
}

type service struct {
	// repo
	rp RepoInterface
}

// constructor
func NewService(rp RepoInterface) ServiceInterface {
	return &service{rp: rp}
}

//func (sv *service) LoadProducts(path string) {}

func (sv *service) ValidateCodeValue(codeValue string) bool {
	return sv.rp.ValidateCodeValue(codeValue)
}
func (sv *service) GetAllProducts() ([]domain.Product, error) {
	return sv.rp.GetAllProducts()
}
func (sv *service) GetProduct(id int) (pr domain.Product, err error) {
	return sv.rp.GetProduct(id)
}
func (sv *service) SearchProduct(priceGt float64) ([]domain.Product, error) {
	return sv.rp.SearchProduct(priceGt)
}
func (sv *service) PostProduct(product domain.Product) (err error) {
	return sv.rp.PostProduct(product)
}

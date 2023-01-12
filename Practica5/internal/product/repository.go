package product

import (
	"errors"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
	"github.com/bootcamp-go/Consignas-Go-Web.git/pkg/storage"
)

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Update(id int, p domain.Product) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	storage2 storage.Store
	//list []domain.Product
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage2 storage.Store) Repository {
	return &repository{storage2}
}

// GetAll devuelve todos los productos
func (r *repository) GetAll() []domain.Product {
	products, err := r.storage2.GetAll()
	if err != nil {
		return []domain.Product{}
	}
	return products
}

// GetByID busca un producto por su id
func (r *repository) GetByID(id int) (domain.Product, error) {
	product, err := r.storage2.GetOne(id)
	if err != nil {
		return domain.Product{}, errors.New("product not found")
	}
	return product, nil

}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (r *repository) SearchPriceGt(price float64) []domain.Product {
	var products []domain.Product
	list, err := r.storage2.GetAll()
	if err != nil {
		return products
	}
	for _, product := range list {
		if product.Price > price {
			products = append(products, product)
		}
	}
	return products
}

// Create agrega un nuevo producto
func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	err := r.storage2.AddOne(p)
	if err != nil {
		return domain.Product{}, errors.New("error creating product")
	}
	return p, nil
}

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *repository) validateCodeValue(codeValue string) bool {
	list, err := r.storage2.GetAll()
	if err != nil {
		return false
	}
	for _, product := range list {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}

// Delete elimina un producto
func (r *repository) Delete(id int) error {
	err := r.storage2.DeleteOne(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un producto
func (r *repository) Update(id int, p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	err := r.storage2.UpdateOne(p)
	if err != nil {
		return domain.Product{}, errors.New("error updating product")
	}
	return p, nil
}

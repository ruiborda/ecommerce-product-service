// filepath: /home/rui/ecommerce/ecommerce-product-service/src/repository/ProductRepository.go
package repository

import (
	"github.com/ruiborda/ecommerce-product-service/src/model"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) (*model.Product, error)
	GetProductById(id string) (*model.Product, error)
	UpdateProduct(product *model.Product) (*model.Product, error)
	DeleteProductById(id string) error
	GetProducts() ([]*model.Product, error)
}

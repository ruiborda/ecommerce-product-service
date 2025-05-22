package service

import (
	"github.com/ruiborda/ecommerce-product-service/src/dto/product"
)

// CategoryService define las operaciones de negocio para las categorías
type CategoryService interface {
	// CreateCategory crea una nueva categoría
	CreateCategory(createRequest *product.CreateCategoryRequest) (*product.CreateCategoryResponse, error)
	
	// UpdateCategory actualiza una categoría existente
	UpdateCategory(updateRequest *product.UpdateCategoryRequest) (*product.UpdateCategoryResponse, error)
	
	// GetCategories obtiene todas las categorías
	GetCategories() (*product.GetCategoriesResponse, error)
}
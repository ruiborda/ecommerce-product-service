package service

import (
	"github.com/ruiborda/ecommerce-product-service/src/dto/category"
	"github.com/ruiborda/ecommerce-product-service/src/dto/product"
)

// CategoryService define las operaciones de negocio para las categorías
type CategoryService interface {
	// CreateCategory crea una nueva categoría
	CreateCategory(createRequest *category.CreateCategoryRequest) (*product.CreateCategoryResponse, error)
	
	// UpdateCategory actualiza una categoría existente
	UpdateCategory(updateRequest *category.UpdateCategoryRequest) (*product.UpdateCategoryResponse, error)
	
	// GetAllCategoriesAsArray obtiene todas las categorías como un array
	GetAllCategoriesAsArray() *[]*category.GetCategoriesResponse
}

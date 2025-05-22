package service

import (
	"github.com/ruiborda/ecommerce-product-service/src/dto/category"
)

// CategoryService define las operaciones de negocio para las categorías
type CategoryService interface {
	// CreateCategory crea una nueva categoría
	CreateCategory(createRequest *category.CreateCategoryRequest) (*category.CreateCategoryResponse, error)

	// UpdateCategory actualiza una categoría existente
	UpdateCategory(updateRequest *category.UpdateCategoryRequest) (*category.UpdateCategoryResponse, error)

	// GetAllCategoriesAsArray obtiene todas las categorías como un array
	GetAllCategoriesAsArray() *[]*category.GetCategoriesResponse
}

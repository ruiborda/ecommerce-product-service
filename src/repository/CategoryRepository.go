package repository

import "github.com/ruiborda/ecommerce-product-service/src/model"

// CategoryRepository define las operaciones de acceso a datos para el modelo Category
type CategoryRepository interface {
	// CreateCategory crea una nueva categoría en la base de datos
	CreateCategory(category *model.Category) (*model.Category, error)
	
	// GetCategoryById obtiene una categoría por su ID
	GetCategoryById(id string) (*model.Category, error)
	
	// UpdateCategory actualiza una categoría existente
	UpdateCategory(category *model.Category) (*model.Category, error)
	
	// DeleteCategoryById elimina una categoría por su ID
	DeleteCategoryById(id string) error
	
	// GetCategories obtiene todas las categorías
	GetCategories() ([]*model.Category, error)
}
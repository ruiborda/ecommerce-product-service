package mapper

import (
	"github.com/ruiborda/ecommerce-product-service/src/dto/product"
	"github.com/ruiborda/ecommerce-product-service/src/model"
)

// CategoryMapper proporciona métodos para convertir entre DTOs y modelos de categoría
type CategoryMapper struct{}

// CreateRequestToCategory convierte un DTO de solicitud de creación a un modelo de categoría
func (cm *CategoryMapper) CreateRequestToCategory(request *product.CreateCategoryRequest) *model.Category {
	return &model.Category{
		Name: request.Name,
	}
}

// CategoryToCreateResponse convierte un modelo de categoría a un DTO de respuesta de creación
func (cm *CategoryMapper) CategoryToCreateResponse(category *model.Category) *product.CreateCategoryResponse {
	return &product.CreateCategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

// UpdateRequestToCategory convierte un DTO de solicitud de actualización a un modelo de categoría
func (cm *CategoryMapper) UpdateRequestToCategory(id string, request *product.UpdateCategoryRequest) *model.Category {
	return &model.Category{
		Id:   id,
		Name: request.Name,
	}
}

// CategoryToUpdateResponse convierte un modelo de categoría a un DTO de respuesta de actualización
func (cm *CategoryMapper) CategoryToUpdateResponse(category *model.Category) *product.UpdateCategoryResponse {
	return &product.UpdateCategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

// CategoriesToGetCategoriesResponse convierte una lista de modelos de categoría a un DTO de respuesta
func (cm *CategoryMapper) CategoriesToGetCategoriesResponse(categories []*model.Category) *product.GetCategoriesResponse {
	response := &product.GetCategoriesResponse{
		Categories: make([]product.CategoryDTO, 0, len(categories)),
	}

	for _, category := range categories {
		response.Categories = append(response.Categories, product.CategoryDTO{
			Id:   category.Id,
			Name: category.Name,
		})
	}

	return response
}
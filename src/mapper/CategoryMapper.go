package mapper

import (
	"github.com/ruiborda/ecommerce-product-service/src/dto/category"
	"github.com/ruiborda/ecommerce-product-service/src/dto/product"
	"github.com/ruiborda/ecommerce-product-service/src/model"
)

// CategoryMapper proporciona métodos para convertir entre DTOs y modelos de categoría
type CategoryMapper struct{}

// CreateRequestToCategory convierte un DTO de solicitud de creación a un modelo de categoría
func (cm *CategoryMapper) CreateRequestToCategory(request *category.CreateCategoryRequest) *model.Category {
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
func (cm *CategoryMapper) UpdateRequestToCategory(request *category.UpdateCategoryRequest) *model.Category {
	return &model.Category{
		Id:   request.Id,
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

// CategoriesToDTOArray convierte una lista de modelos de categoría a un array de DTOs
func (cm *CategoryMapper) CategoriesToDTOArray(categories []*model.Category) *[]*category.GetCategoriesResponse {
	responses := make([]*category.GetCategoriesResponse, 0, len(categories))
	
	for _, categoryModel := range categories {
		response := &category.GetCategoriesResponse{
			Id:   categoryModel.Id,
			Name: categoryModel.Name,
		}
		responses = append(responses, response)
	}
	
	return &responses
}
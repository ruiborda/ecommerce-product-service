package impl

import (
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/ruiborda/ecommerce-product-service/src/dto/product"
	"github.com/ruiborda/ecommerce-product-service/src/mapper"
	"github.com/ruiborda/ecommerce-product-service/src/repository"
	repoImpl "github.com/ruiborda/ecommerce-product-service/src/repository/impl"
	"github.com/ruiborda/ecommerce-product-service/src/service"
)

// CategoryServiceImpl implementa la interfaz CategoryService
type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
	categoryMapper     *mapper.CategoryMapper
}

// NewCategoryServiceImpl crea una nueva instancia de CategoryServiceImpl
func NewCategoryServiceImpl() service.CategoryService {
	return &CategoryServiceImpl{
		categoryRepository: repoImpl.NewCategoryRepositoryImpl(),
		categoryMapper:     &mapper.CategoryMapper{},
	}
}

// CreateCategory implementa la creación de una nueva categoría
func (s *CategoryServiceImpl) CreateCategory(createRequest *product.CreateCategoryRequest, authorId string) (*product.CreateCategoryResponse, error) {
	// Validar datos de entrada
	if createRequest == nil {
		return nil, errors.New("request cannot be nil")
	}
	
	if createRequest.Name == "" {
		return nil, errors.New("category name is required")
	}
	
	// Generar un ID único para la categoría
	categoryId := uuid.New().String()
	
	// Crear el modelo de categoría usando el mapper
	categoryModel := s.categoryMapper.CreateRequestToCategory(createRequest)
	categoryModel.Id = categoryId
	
	// Guardar la categoría en la base de datos
	createdCategory, err := s.categoryRepository.CreateCategory(categoryModel)
	if err != nil {
		slog.Error("Error creating category", "error", err)
		return nil, err
	}
	
	// Crear la respuesta usando el mapper
	return s.categoryMapper.CategoryToCreateResponse(createdCategory), nil
}

// UpdateCategory implementa la actualización de una categoría existente
func (s *CategoryServiceImpl) UpdateCategory(id string, updateRequest *product.UpdateCategoryRequest) (*product.UpdateCategoryResponse, error) {
	// Validar datos de entrada
	if updateRequest == nil {
		return nil, errors.New("request cannot be nil")
	}
	
	if updateRequest.Name == "" {
		return nil, errors.New("category name is required")
	}
	
	// Verificar si la categoría existe
	existingCategory, err := s.categoryRepository.GetCategoryById(id)
	if err != nil {
		slog.Error("Error fetching category for update", "id", id, "error", err)
		return nil, err
	}
	
	if existingCategory == nil {
		return nil, errors.New("category not found")
	}
	
	// Actualizar el modelo de categoría usando el mapper
	categoryModel := s.categoryMapper.UpdateRequestToCategory(id, updateRequest)
	
	// Guardar la categoría actualizada en la base de datos
	updatedCategory, err := s.categoryRepository.UpdateCategory(categoryModel)
	if err != nil {
		slog.Error("Error updating category", "id", id, "error", err)
		return nil, err
	}
	
	// Crear la respuesta usando el mapper
	return s.categoryMapper.CategoryToUpdateResponse(updatedCategory), nil
}

// GetCategories implementa la obtención de todas las categorías
func (s *CategoryServiceImpl) GetCategories() (*product.GetCategoriesResponse, error) {
	// Obtener todas las categorías desde el repositorio
	categories, err := s.categoryRepository.GetCategories()
	if err != nil {
		slog.Error("Error fetching categories", "error", err)
		return nil, err
	}
	
	// Crear la respuesta usando el mapper
	return s.categoryMapper.CategoriesToGetCategoriesResponse(categories), nil
}
package impl

import (
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/ruiborda/ecommerce-product-service/src/dto/category"
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
func (s *CategoryServiceImpl) CreateCategory(createRequest *category.CreateCategoryRequest) (*product.CreateCategoryResponse, error) {
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
func (s *CategoryServiceImpl) UpdateCategory(updateRequest *category.UpdateCategoryRequest) (*product.UpdateCategoryResponse, error) {
	// Validar datos de entrada
	if updateRequest == nil {
		return nil, errors.New("request cannot be nil")
	}

	if updateRequest.Id == "" {
		return nil, errors.New("category id is required")
	}

	if updateRequest.Name == "" {
		return nil, errors.New("category name is required")
	}

	// Verificar si la categoría existe
	existingCategory, err := s.categoryRepository.GetCategoryById(updateRequest.Id)
	if err != nil {
		slog.Error("Error fetching category for update", "id", updateRequest.Id, "error", err)
		return nil, err
	}

	if existingCategory == nil {
		return nil, errors.New("category not found")
	}

	// Actualizar sólo los campos proporcionados en la solicitud
	existingCategory.Name = updateRequest.Name

	// Guardar la categoría actualizada en la base de datos
	updatedCategory, err := s.categoryRepository.UpdateCategory(existingCategory)
	if err != nil {
		slog.Error("Error updating category", "id", updateRequest.Id, "error", err)
		return nil, err
	}

	// Crear la respuesta usando el mapper
	return s.categoryMapper.CategoryToUpdateResponse(updatedCategory), nil
}

// GetCategories implementa la obtención de todas las categorías
// Esta función ya no es necesaria y puede ser eliminada
// GetCategories() (*category.GetCategoriesResponse, error)

// GetAllCategoriesAsArray implementa la obtención de todas las categorías como un array
func (s *CategoryServiceImpl) GetAllCategoriesAsArray() *[]*category.GetCategoriesResponse {
	// Obtener todas las categorías desde el repositorio
	categories, err := s.categoryRepository.GetCategories()
	if err != nil {
		slog.Error("Error fetching categories", "error", err)
		result := make([]*category.GetCategoriesResponse, 0)
		return &result // Devolver puntero a array vacío en caso de error
	}

	// Convertir las categorías a DTOs y devolverlas como array
	response := make([]*category.GetCategoriesResponse, 0, len(categories))

	for _, categoryModel := range categories {
		categoryDTO := &category.GetCategoriesResponse{
			Id:   categoryModel.Id,
			Name: categoryModel.Name,
		}
		response = append(response, categoryDTO)
	}

	return &response
}

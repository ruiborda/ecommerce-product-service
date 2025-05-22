package category

// CreateCategoryRequest DTO para la creación de una categoría
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

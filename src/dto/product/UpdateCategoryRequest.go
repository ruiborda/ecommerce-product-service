package product

// UpdateCategoryRequest DTO para la actualización de una categoría
type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
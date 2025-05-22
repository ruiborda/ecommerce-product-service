package product

// UpdateCategoryRequest DTO para la actualización de una categoría
type UpdateCategoryRequest struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
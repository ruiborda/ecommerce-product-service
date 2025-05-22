package product

// UpdateCategoryResponse DTO para la respuesta de actualización de una categoría
type UpdateCategoryResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
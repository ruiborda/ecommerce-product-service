package product

// CreateCategoryResponse DTO para la respuesta de creación de una categoría
type CreateCategoryResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
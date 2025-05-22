package product

// CategoryDTO representa una categoría en las respuestas
type CategoryDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GetCategoriesResponse DTO para la respuesta del listado de categorías
type GetCategoriesResponse struct {
	Categories []CategoryDTO `json:"categories"`
}
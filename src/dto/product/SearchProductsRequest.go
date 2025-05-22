// filepath: /home/rui/ecommerce/ecommerce-product-service/src/dto/product/SearchProductsRequest.go
package product

// SearchProductsRequest implementa una búsqueda avanzada con filtros similares a Amazon
type SearchProductsRequest struct {
	// Parámetros de paginación
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`

	// Término de búsqueda principal
	Query string `json:"query" form:"query"`

	// Filtros de producto
	CategoryId string  `json:"categoryId" form:"categoryId"`
	PriceMin   float64 `json:"priceMin" form:"priceMin"`
	PriceMax   float64 `json:"priceMax" form:"priceMax"`

	// Ordenamiento
	SortBy        string `json:"sortBy" form:"sortBy"`               // price, name, created_at, etc.
	SortDirection string `json:"sortDirection" form:"sortDirection"` // asc, desc
}

// filepath: /home/rui/ecommerce/ecommerce-product-service/src/dto/product/SearchProductsResponse.go
package product

type SearchProductsResponse struct {
	Id           string  `json:"id"`
	CategoryId   string  `json:"categoryId"`
	CategoryName string  `json:"categoryName"`
	AuthorId     string  `json:"authorId"`
	AuthorName   string  `json:"authorName"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Currency     string  `json:"currency"`
	Discount     float64 `json:"discount"`
	Sku          string  `json:"sku"`
	Stock        int     `json:"stock"`
	FileImage    string  `json:"fileImage"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}

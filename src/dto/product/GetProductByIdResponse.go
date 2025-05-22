// filepath: /home/rui/ecommerce/ecommerce-product-service/src/dto/product/GetProductByIdResponse.go
package product

type GetProductByIdResponse struct {
	Id           string  `json:"id"`
	CategoryId   string  `json:"categoryId"`
	CategoryName string  `json:"categoryName"`
	AuthorId     string  `json:"authorId"`
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

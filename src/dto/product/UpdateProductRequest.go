// filepath: /home/rui/ecommerce/ecommerce-product-service/src/dto/product/UpdateProductRequest.go
package product

type UpdateProductRequest struct {
	Id          string  `json:"id"`
	CategoryId  string  `json:"categoryId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Discount    float64 `json:"discount"`
	Sku         string  `json:"sku"`
	Stock       int     `json:"stock"`
	ImageBase64 string  `json:"imageBase64"`
}

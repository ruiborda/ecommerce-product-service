package product

type CreateProductResponse struct {
	Id          string  `json:"id"`
	CategoryId  string  `json:"categoryId"`
	AuthorId    string  `json:"authorId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Discount    float64 `json:"discount"`
	Sku         string  `json:"sku"`
	Stock       int     `json:"stock"`
	FileImage   string  `json:"fileImage"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

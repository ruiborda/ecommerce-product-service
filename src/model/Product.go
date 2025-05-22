package model

type Product struct {
	Id          string  `json:"id,omitempty"          firestore:"id,omitempty"`
	CategoryId  string  `json:"categoryId,omitempty"  firestore:"categoryId,omitempty"`
	AuthorId    string  `json:"authorId,omitempty"    firestore:"authorId,omitempty"`
	Name        string  `json:"name,omitempty"        firestore:"name,omitempty"`
	Description string  `json:"description,omitempty" firestore:"description,omitempty"`
	Price       float64 `json:"price,omitempty"       firestore:"price,omitempty"`
	Currency    string  `json:"currency,omitempty"    firestore:"currency,omitempty"`
	Discount    float64 `json:"discount,omitempty"    firestore:"discount,omitempty"`
	Sku         string  `json:"sku,omitempty"         firestore:"sku,omitempty"`
	Stock       int     `json:"stock,omitempty"       firestore:"stock,omitempty"`
	FileImage   string  `json:"fileImage,omitempty"   firestore:"fileImage,omitempty"`
	CreatedAt   string  `json:"createdAt,omitempty"   firestore:"createdAt,omitempty"`
	UpdatedAt   string  `json:"updatedAt,omitempty"   firestore:"updatedAt,omitempty"`
}

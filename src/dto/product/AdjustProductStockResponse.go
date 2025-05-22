// filepath: /home/rui/ecommerce/ecommerce-product-service/src/dto/product/AdjustProductStockResponse.go
package product

type AdjustProductStockResponse struct {
	Id            string `json:"id"`
	PreviousStock int    `json:"previousStock"`
	CurrentStock  int    `json:"currentStock"`
}

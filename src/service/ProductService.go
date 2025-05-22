package service

import (
	dto "github.com/ruiborda/ecommerce-product-service/src/dto/common"
	"github.com/ruiborda/ecommerce-product-service/src/dto/product"
)

type ProductService interface {
	// CreateProduct crea un nuevo producto en el sistema
	CreateProduct(createProductRequest *product.CreateProductRequest, authorId string) (*product.CreateProductResponse, error)

	// GetProductById obtiene un producto por su ID
	GetProductById(id string) (*product.GetProductByIdResponse, error)

	// UpdateProduct actualiza un producto existente por su ID
	UpdateProduct(id string, updateProductRequest *product.UpdateProductRequest) (*product.UpdateProductResponse, error)

	// DeleteProduct elimina un producto por su ID
	DeleteProduct(id string) (*product.DeleteProductByIdResponse, error)

	// GetProductsPaginated obtiene una lista paginada de productos
	GetProductsPaginated(pageable *dto.Pageable) (*dto.PaginationResponse[product.GetProductsPaginatedResponse], error)

	// AdjustProductStock ajusta el stock de un producto
	AdjustProductStock(id string, request *product.AdjustProductStockRequest) (*product.AdjustProductStockResponse, error)

	// SearchProducts busca productos con filtros avanzados
	SearchProducts(request *product.SearchProductsRequest) (*dto.PaginationResponse[product.SearchProductsResponse], error)
}

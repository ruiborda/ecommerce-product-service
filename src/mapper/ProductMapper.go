// filepath: /home/rui/ecommerce/ecommerce-product-service/src/mapper/ProductMapper.go
package mapper

import (
	"time"

	"github.com/ruiborda/ecommerce-product-service/src/dto/product"
	"github.com/ruiborda/ecommerce-product-service/src/model"
)

// ProductMapper struct
type ProductMapper struct {
}

// CreateRequestToProduct convierte un CreateProductRequest a un modelo Product
func (m *ProductMapper) CreateRequestToProduct(request *product.CreateProductRequest) *model.Product {
	now := time.Now().Format(time.RFC3339)
	return &model.Product{
		// ID será asignado por el servicio
		// AuthorId será asignado desde el JWT
		CategoryId:  request.CategoryId,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Currency:    request.Currency,
		Discount:    request.Discount,
		Sku:         request.Sku,
		Stock:       request.Stock,
		FileImage:   "",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateRequestToProduct convierte un UpdateProductRequest a un modelo Product parcial
func (m *ProductMapper) UpdateRequestToProduct(request *product.UpdateProductRequest) *model.Product {
	return &model.Product{
		// ID será asignado por el servicio
		CategoryId:  request.CategoryId,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Currency:    request.Currency,
		Discount:    request.Discount,
		Sku:         request.Sku,
		Stock:       request.Stock,
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
}

// ProductToCreateResponse convierte un modelo Product a un CreateProductResponse
func (m *ProductMapper) ProductToCreateResponse(model *model.Product) *product.CreateProductResponse {
	return &product.CreateProductResponse{
		Id:          model.Id,
		CategoryId:  model.CategoryId,
		AuthorId:    model.AuthorId,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		Currency:    model.Currency,
		Discount:    model.Discount,
		Sku:         model.Sku,
		Stock:       model.Stock,
		FileImage:   model.FileImage,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

// ProductToGetByIdResponse convierte un modelo Product a un GetProductByIdResponse básico
func (m *ProductMapper) ProductToGetByIdResponse(model *model.Product) *product.GetProductByIdResponse {
	return &product.GetProductByIdResponse{
		Id:          model.Id,
		CategoryId:  model.CategoryId,
		AuthorId:    model.AuthorId,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		Currency:    model.Currency,
		Discount:    model.Discount,
		Sku:         model.Sku,
		Stock:       model.Stock,
		FileImage:   model.FileImage,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		// CategoryName se agregará en el servicio
	}
}

// ProductToUpdateResponse convierte un modelo Product a un UpdateProductResponse
func (m *ProductMapper) ProductToUpdateResponse(model *model.Product) *product.UpdateProductResponse {
	return &product.UpdateProductResponse{
		Id:          model.Id,
		CategoryId:  model.CategoryId,
		AuthorId:    model.AuthorId,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		Currency:    model.Currency,
		Discount:    model.Discount,
		Sku:         model.Sku,
		Stock:       model.Stock,
		FileImage:   model.FileImage,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

// ProductToGetPaginatedResponse convierte un modelo Product a un GetProductsPaginatedResponse
func (m *ProductMapper) ProductToGetPaginatedResponse(model *model.Product) *product.GetProductsPaginatedResponse {
	return &product.GetProductsPaginatedResponse{
		Id:          model.Id,
		CategoryId:  model.CategoryId,
		AuthorId:    model.AuthorId,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		Currency:    model.Currency,
		Discount:    model.Discount,
		Sku:         model.Sku,
		Stock:       model.Stock,
		FileImage:   model.FileImage,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

// ProductToSearchResponse convierte un modelo Product a un SearchProductsResponse básico
func (m *ProductMapper) ProductToSearchResponse(model *model.Product) *product.SearchProductsResponse {
	return &product.SearchProductsResponse{
		Id:          model.Id,
		CategoryId:  model.CategoryId,
		AuthorId:    model.AuthorId,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		Currency:    model.Currency,
		Discount:    model.Discount,
		Sku:         model.Sku,
		Stock:       model.Stock,
		FileImage:   model.FileImage,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		// CategoryName y AuthorName se agregarán en el servicio
	}
}

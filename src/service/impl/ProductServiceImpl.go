// filepath: /home/rui/ecommerce/ecommerce-product-service/src/service/impl/ProductServiceImpl.go
package impl

import (
	"os"
	"time"

	"github.com/ruiborda/ecommerce-product-service/src/repository/impl"

	"log/slog"

	"github.com/google/uuid"
	dto "github.com/ruiborda/ecommerce-product-service/src/dto/common"
	"github.com/ruiborda/ecommerce-product-service/src/dto/product"
	"github.com/ruiborda/ecommerce-product-service/src/mapper"
	"github.com/ruiborda/ecommerce-product-service/src/model"
	"github.com/ruiborda/ecommerce-product-service/src/repository"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
	r2Repository      repository.R2Repository
	productMapper     *mapper.ProductMapper
}

func NewProductServiceImpl() *ProductServiceImpl {
	return &ProductServiceImpl{
		productRepository: impl.NewProductRepositoryImpl(),
		r2Repository: impl.NewR2RepositoryImpl(
			"ecommerce",
			os.Getenv("R2_ACCOUNT_ID"),
			os.Getenv("R2_ACCESS_KEY"),
			os.Getenv("R2_SECRET_KEY"),
		),
		productMapper: &mapper.ProductMapper{},
	}
}

// CreateProduct implementa la creación de un nuevo producto
func (ps *ProductServiceImpl) CreateProduct(createRequest *product.CreateProductRequest, authorId string) (*product.CreateProductResponse, error) {
	// Generar un ID único para el producto
	productId := uuid.New().String()

	// Crear el modelo de producto usando el mapper y luego asignar el ID
	productModel := ps.productMapper.CreateRequestToProduct(createRequest)
	productModel.Id = productId

	// Asignar el ID del autor extraído del token JWT
	productModel.AuthorId = authorId

	// Procesar la imagen si existe
	if createRequest.ImageBase64 != "" {
		fileName, err := ps.r2Repository.UploadBase64File(&createRequest.ImageBase64)
		if err != nil {
			slog.Error("Error uploading product image", "error", err)
			return nil, err
		}
		productModel.FileImage = fileName
	}

	// Guardar el producto en la base de datos
	createdProduct, err := ps.productRepository.CreateProduct(productModel)
	if err != nil {
		// Si hubo error y se subió una imagen, eliminarla
		if productModel.FileImage != "" {
			_ = ps.r2Repository.DeleteFile(productModel.FileImage)
		}
		slog.Error("Error creating product", "error", err)
		return nil, err
	}

	// Crear la respuesta usando el mapper
	return ps.productMapper.ProductToCreateResponse(createdProduct), nil
}

// GetProductById obtiene los detalles de un producto por su ID
func (ps *ProductServiceImpl) GetProductById(id string) (*product.GetProductByIdResponse, error) {
	// Obtener el producto desde el repositorio
	productModel, err := ps.productRepository.GetProductById(id)
	if err != nil {
		slog.Error("Error getting product", "id", id, "error", err)
		return nil, err
	}

	if productModel == nil {
		return nil, nil
	}

	// Crear la respuesta básica usando el mapper
	response := ps.productMapper.ProductToGetByIdResponse(productModel)

	// Agregar información adicional como el nombre de la categoría
	// En una implementación real, aquí se obtendría el nombre de la categoría desde un servicio
	response.CategoryName = "Categoría Default" // Para este ejemplo usamos un valor por defecto

	return response, nil
}

// UpdateProduct actualiza un producto existente
func (ps *ProductServiceImpl) UpdateProduct(id string, updateRequest *product.UpdateProductRequest) (*product.UpdateProductResponse, error) {
	// Verificar si el producto existe
	existingProduct, err := ps.productRepository.GetProductById(id)
	if err != nil {
		slog.Error("Error getting product for update", "id", id, "error", err)
		return nil, err
	}

	if existingProduct == nil {
		return nil, nil
	}

	// Crear un modelo parcial con los datos de actualización
	updateModel := ps.productMapper.UpdateRequestToProduct(updateRequest)

	// Mantener campos que no deben cambiar
	updateModel.Id = existingProduct.Id
	updateModel.AuthorId = existingProduct.AuthorId
	updateModel.FileImage = existingProduct.FileImage
	updateModel.CreatedAt = existingProduct.CreatedAt

	// Procesar la imagen si se proporcionó una nueva
	if updateRequest.ImageBase64 != "" {
		// Eliminar la imagen anterior si existe
		if updateModel.FileImage != "" {
			_ = ps.r2Repository.DeleteFile(updateModel.FileImage)
		}

		// Subir la nueva imagen
		fileName, err := ps.r2Repository.UploadBase64File(&updateRequest.ImageBase64)
		if err != nil {
			slog.Error("Error uploading updated product image", "error", err)
			return nil, err
		}
		updateModel.FileImage = fileName
	}

	// Guardar los cambios en la base de datos
	updatedProduct, err := ps.productRepository.UpdateProduct(updateModel)
	if err != nil {
		// Si hubo error y se subió una imagen nueva, eliminarla
		if updateRequest.ImageBase64 != "" && updateModel.FileImage != "" {
			_ = ps.r2Repository.DeleteFile(updateModel.FileImage)
		}
		slog.Error("Error updating product", "id", id, "error", err)
		return nil, err
	}

	// Crear y devolver la respuesta usando el mapper
	return ps.productMapper.ProductToUpdateResponse(updatedProduct), nil
}

// DeleteProduct elimina un producto por su ID
func (ps *ProductServiceImpl) DeleteProduct(id string) (*product.DeleteProductByIdResponse, error) {
	// Verificar si el producto existe
	existingProduct, err := ps.productRepository.GetProductById(id)
	if err != nil {
		slog.Error("Error getting product for delete", "id", id, "error", err)
		return nil, err
	}

	if existingProduct == nil {
		return &product.DeleteProductByIdResponse{
			Success: false,
			Message: "Product not found",
		}, nil
	}

	// Eliminar la imagen asociada si existe
	if existingProduct.FileImage != "" {
		if err := ps.r2Repository.DeleteFile(existingProduct.FileImage); err != nil {
			slog.Error("Error deleting product image", "fileName", existingProduct.FileImage, "error", err)
		}
	}

	// Eliminar el producto de la base de datos
	err = ps.productRepository.DeleteProductById(id)
	if err != nil {
		slog.Error("Error deleting product", "id", id, "error", err)
		return nil, err
	}

	// Devolver respuesta exitosa
	return &product.DeleteProductByIdResponse{
		Success: true,
		Message: "Product successfully deleted",
	}, nil
}

// GetProductsPaginated obtiene productos con paginación
func (ps *ProductServiceImpl) GetProductsPaginated(pageable *dto.Pageable) (*dto.PaginationResponse[product.GetProductsPaginatedResponse], error) {
	// Obtener los productos desde el repositorio
	products, err := ps.productRepository.GetProducts()
	if err != nil {
		slog.Error("Error getting products for pagination", "error", err)
		return nil, err
	}

	// Implementar paginación manual (en una implementación real, esto se haría en la base de datos)
	totalElements := len(products)
	startIndex := (pageable.Page - 1) * pageable.Size
	endIndex := startIndex + pageable.Size

	if startIndex >= totalElements {
		startIndex = 0
		endIndex = 0
	}

	if endIndex > totalElements {
		endIndex = totalElements
	}

	var paginatedProducts []*product.GetProductsPaginatedResponse

	// Convertir los productos a DTOs usando el mapper
	for i := startIndex; i < endIndex; i++ {
		p := products[i]
		paginatedProducts = append(paginatedProducts, ps.productMapper.ProductToGetPaginatedResponse(p))
	}

	// Construir la respuesta paginada
	totalPages := (totalElements + pageable.Size - 1) / pageable.Size
	if totalPages == 0 {
		totalPages = 1
	}

	result := &dto.PaginationResponse[product.GetProductsPaginatedResponse]{
		Page: dto.Page{
			CurrentPage:   pageable.Page,
			Size:          pageable.Size,
			TotalElements: totalElements,
			TotalPages:    totalPages,
		},
		Links: dto.PageLinks{
			Self: "",
			Next: "",
			Prev: "",
		},
	}

	if paginatedProducts == nil {
		paginatedProducts = make([]*product.GetProductsPaginatedResponse, 0)
	}

	result.Data = &paginatedProducts

	return result, nil
}

// AdjustProductStock ajusta el stock de un producto
func (ps *ProductServiceImpl) AdjustProductStock(id string, request *product.AdjustProductStockRequest) (*product.AdjustProductStockResponse, error) {
	// Verificar si el producto existe
	existingProduct, err := ps.productRepository.GetProductById(id)
	if err != nil {
		slog.Error("Error getting product for stock adjustment", "id", id, "error", err)
		return nil, err
	}

	if existingProduct == nil {
		return nil, nil
	}

	// Guardar el stock anterior
	previousStock := existingProduct.Stock

	// Actualizar el stock
	existingProduct.Stock = previousStock + request.Quantity

	// Evitar stock negativo
	if existingProduct.Stock < 0 {
		existingProduct.Stock = 0
	}

	// Actualizar la fecha de modificación
	existingProduct.UpdatedAt = time.Now().Format(time.RFC3339)

	// Guardar los cambios
	_, err = ps.productRepository.UpdateProduct(existingProduct)
	if err != nil {
		slog.Error("Error updating product stock", "id", id, "error", err)
		return nil, err
	}

	// Crear respuesta manualmente ya que no tenemos un mapper específico para esto
	return &product.AdjustProductStockResponse{
		Id:            id,
		PreviousStock: previousStock,
		CurrentStock:  existingProduct.Stock,
	}, nil
}

// SearchProducts busca productos con filtros avanzados
func (ps *ProductServiceImpl) SearchProducts(request *product.SearchProductsRequest) (*dto.PaginationResponse[product.SearchProductsResponse], error) {
	// Obtener todos los productos (en una implementación real se usarían filtros en la base de datos)
	allProducts, err := ps.productRepository.GetProducts()
	if err != nil {
		slog.Error("Error getting products for search", "error", err)
		return nil, err
	}

	// Filtrar productos manualmente según los criterios (esto sería más eficiente en la base de datos)
	var filteredProducts []*model.Product

	for _, p := range allProducts {
		include := true

		// Filtrar por query en nombre o descripción
		if request.Query != "" && !containsIgnoreCase(p.Name, request.Query) && !containsIgnoreCase(p.Description, request.Query) {
			include = false
		}

		// Filtrar por categoría
		if include && request.CategoryId != "" && p.CategoryId != request.CategoryId {
			include = false
		}

		// Filtrar por precio mínimo
		if include && request.PriceMin > 0 && p.Price < request.PriceMin {
			include = false
		}

		// Filtrar por precio máximo
		if include && request.PriceMax > 0 && p.Price > request.PriceMax {
			include = false
		}

		if include {
			filteredProducts = append(filteredProducts, p)
		}
	}

	// Ordenar productos (implementación básica)
	sortProducts(filteredProducts, request.SortBy, request.SortDirection)

	// Aplicar paginación
	totalElements := len(filteredProducts)
	startIndex := (request.Page - 1) * request.Size
	endIndex := startIndex + request.Size

	if startIndex >= totalElements {
		startIndex = 0
		endIndex = 0
	}

	if endIndex > totalElements {
		endIndex = totalElements
	}

	var paginatedProducts []*product.SearchProductsResponse

	// Crear las respuestas de productos usando el mapper y añadiendo información adicional
	for i := startIndex; i < endIndex; i++ {
		p := filteredProducts[i]

		// Crear la respuesta básica usando el mapper
		productResponse := ps.productMapper.ProductToSearchResponse(p)

		// Añadir información adicional que no viene del mapper
		productResponse.CategoryName = "Categoría " + p.CategoryId // En una implementación real, se obtendría de un servicio
		productResponse.AuthorName = "Autor " + p.AuthorId         // En una implementación real, se obtendría de un servicio

		paginatedProducts = append(paginatedProducts, productResponse)
	}

	// Construir la respuesta paginada
	totalPages := (totalElements + request.Size - 1) / request.Size
	if totalPages == 0 {
		totalPages = 1
	}

	result := &dto.PaginationResponse[product.SearchProductsResponse]{
		Page: dto.Page{
			CurrentPage:   request.Page,
			Size:          request.Size,
			TotalElements: totalElements,
			TotalPages:    totalPages,
		},
		Links: dto.PageLinks{
			Self: "",
			Next: "",
			Prev: "",
		},
	}

	if paginatedProducts == nil {
		paginatedProducts = make([]*product.SearchProductsResponse, 0)
	}

	result.Data = &paginatedProducts

	return result, nil
}

// Función auxiliar para buscar texto ignorando mayúsculas/minúsculas
func containsIgnoreCase(s, substr string) bool {
	// Implementación simple para ejemplo
	return true // En una implementación real se haría la comparación correctamente
}

// Función auxiliar para ordenar productos
func sortProducts(products []*model.Product, sortBy, sortDirection string) {
	// Implementación simple para ejemplo
	// En una implementación real se ordenarían los productos según los criterios
}

// filepath: /home/rui/ecommerce/ecommerce-product-service/src/controller/ProductController.go
package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ruiborda/ecommerce-product-service/src/dto/auth"
	dto "github.com/ruiborda/ecommerce-product-service/src/dto/common"
	"github.com/ruiborda/ecommerce-product-service/src/dto/product"
	"github.com/ruiborda/ecommerce-product-service/src/service"
	"github.com/ruiborda/ecommerce-product-service/src/service/impl"
	"github.com/ruiborda/go-jwt/src/domain/entity"
	"github.com/ruiborda/go-swagger-generator/src/openapi"
	"github.com/ruiborda/go-swagger-generator/src/openapi_spec/mime"
	"github.com/ruiborda/go-swagger-generator/src/swagger"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		productService: impl.NewProductServiceImpl(),
	}
}

var _ = swagger.Swagger().Path("/api/v1/products").
	Post(func(operation openapi.Operation) {
		operation.Summary("Create a new product").
			OperationID("CreateProduct").
			Tag("ProductController").
			Consume(mime.ApplicationJSON).
			Produces(mime.ApplicationJSON).
			BodyParameter(func(param openapi.Parameter) {
				param.Description("Product object that needs to be added to the system").
					Required(true).
					SchemaFromDTO(&product.CreateProductRequest{})
			}).
			Security("BearerAuth")
	}).Doc()

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var createProductRequest = &product.CreateProductRequest{}

	if err := c.BindJSON(createProductRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extraer el ID del autor del token JWT
	claimsValue, exists := c.Get("jwtClaims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No JWT claims found"})
		return
	}

	// Obtener el ID del usuario (autor) desde el JWT
	var authorId string

	claims, ok := claimsValue.(*entity.JWTClaims[*auth.JwtPrivateClaims])
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JWT claims format"})
		return
	}

	// Obtener el subject (ID de usuario) desde los claims registrados
	authorId = claims.RegisteredClaims.Subject

	if authorId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in token"})
		return
	}

	// Llamar al servicio pasando tanto el DTO como el authorId extraído del JWT
	response, err := pc.productService.CreateProduct(createProductRequest, authorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

var _ = swagger.Swagger().Path("/api/v1/products/{id}").
	Get(func(operation openapi.Operation) {
		operation.Summary("Get product by ID").
			OperationID("GetProductById").
			Tag("ProductController").
			Produces(mime.ApplicationJSON).
			PathParameter("id", func(param openapi.Parameter) {
				param.Description("ID of the product to get").
					Required(true).
					Type("string")
			}).
			Security("BearerAuth")
	}).Doc()

func (pc *ProductController) GetProductById(c *gin.Context) {
	id := c.Param("id")

	response, err := pc.productService.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if response == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, response)
}

var _ = swagger.Swagger().Path("/api/v1/products/{id}").
	Put(func(operation openapi.Operation) {
		operation.Summary("Update an existing product").
			OperationID("UpdateProduct").
			Tag("ProductController").
			Consume(mime.ApplicationJSON).
			Produces(mime.ApplicationJSON).
			PathParameter("id", func(param openapi.Parameter) {
				param.Description("ID of the product to update").
					Required(true).
					Type("string")
			}).
			BodyParameter(func(param openapi.Parameter) {
				param.Description("Product object with updated values").
					Required(true).
					SchemaFromDTO(&product.UpdateProductRequest{})
			}).
			Security("BearerAuth")
	}).Doc()

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updateProductRequest = &product.UpdateProductRequest{}

	if err := c.BindJSON(updateProductRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := pc.productService.UpdateProduct(id, updateProductRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if response == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, response)
}

var _ = swagger.Swagger().Path("/api/v1/products/{id}").
	Delete(func(operation openapi.Operation) {
		operation.Summary("Delete a product").
			OperationID("DeleteProduct").
			Tag("ProductController").
			Produces(mime.ApplicationJSON).
			PathParameter("id", func(param openapi.Parameter) {
				param.Description("ID of the product to delete").
					Required(true).
					Type("string")
			}).
			Security("BearerAuth")
	}).Doc()

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	response, err := pc.productService.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

var _ = swagger.Swagger().Path("/api/v1/products/pages").
	Get(func(operation openapi.Operation) {
		operation.Summary("Get paginated list of products").
			OperationID("GetProductsPaginated").
			Tag("ProductController").
			Produces(mime.ApplicationJSON).
			QueryParameter("page", func(param openapi.Parameter) {
				param.Description("Page number").
					Type("integer").
					Format("int32")
			}).
			QueryParameter("size", func(param openapi.Parameter) {
				param.Description("Page size").
					Type("integer").
					Format("int32")
			}).
			Security("BearerAuth")
	}).Doc()

func (pc *ProductController) GetProductsPaginated(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")
	query := c.DefaultQuery("query", "")

	pageable := dto.NewPageable(pageStr, sizeStr, query)
	response, err := pc.productService.GetProductsPaginated(pageable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

var _ = swagger.Swagger().Path("/api/v1/products/{id}/stock").
	Put(func(operation openapi.Operation) {
		operation.Summary("Adjust product stock").
			OperationID("AdjustProductStock").
			Tag("ProductController").
			Consume(mime.ApplicationJSON).
			Produces(mime.ApplicationJSON).
			PathParameter("id", func(param openapi.Parameter) {
				param.Description("ID of the product to adjust stock").
					Required(true).
					Type("string")
			}).
			BodyParameter(func(param openapi.Parameter) {
				param.Description("Stock adjustment details").
					Required(true).
					SchemaFromDTO(&product.AdjustProductStockRequest{})
			}).
			Security("BearerAuth")
	}).Doc()

func (pc *ProductController) AdjustProductStock(c *gin.Context) {
	id := c.Param("id")
	var adjustStockRequest = &product.AdjustProductStockRequest{}

	if err := c.BindJSON(adjustStockRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := pc.productService.AdjustProductStock(id, adjustStockRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if response == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, response)
}

var _ = swagger.Swagger().Path("/api/v1/products/search").
	Get(func(operation openapi.Operation) {
		operation.Summary("Search products with advanced filters").
			OperationID("SearchProducts").
			Tag("ProductController").
			Produces(mime.ApplicationJSON).
			QueryParameter("query", func(param openapi.Parameter) {
				param.Description("Search query").
					Type("string")
			}).
			QueryParameter("page", func(param openapi.Parameter) {
				param.Description("Page number").
					Type("integer").
					Format("int32")
			}).
			QueryParameter("size", func(param openapi.Parameter) {
				param.Description("Page size").
					Type("integer").
					Format("int32")
			}).
			QueryParameter("categoryId", func(param openapi.Parameter) {
				param.Description("Filter by category ID").
					Type("string")
			}).
			QueryParameter("priceMin", func(param openapi.Parameter) {
				param.Description("Minimum price").
					Type("number").
					Format("float")
			}).
			QueryParameter("priceMax", func(param openapi.Parameter) {
				param.Description("Maximum price").
					Type("number").
					Format("float")
			}).
			QueryParameter("sortBy", func(param openapi.Parameter) {
				param.Description("Field to sort by").
					Type("string")
			}).
			QueryParameter("sortDirection", func(param openapi.Parameter) {
				param.Description("Sort direction (asc or desc)").
					Type("string")
			}).
			Response(200, func(response openapi.Response) {
				response.Description("Successful operation").
					SchemaFromDTO(&product.SearchProductsResponse{})
			}).
			Security("BearerAuth")
	}).Doc()

func (pc *ProductController) SearchProducts(c *gin.Context) {
	query := c.Query("query")
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")
	categoryId := c.Query("categoryId")
	priceMinStr := c.Query("priceMin")
	priceMaxStr := c.Query("priceMax")
	sortBy := c.DefaultQuery("sortBy", "")
	sortDirection := c.DefaultQuery("sortDirection", "asc")

	// Construir la solicitud de búsqueda
	searchRequest := &product.SearchProductsRequest{
		Query: query,
		Page:  1,
		Size:  10,
	}

	// Convertir los parámetros numéricos
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err == nil && page > 0 {
			searchRequest.Page = page
		}
	}

	if sizeStr != "" {
		size, err := strconv.Atoi(sizeStr)
		if err == nil && size > 0 {
			searchRequest.Size = size
		}
	}

	// Agregar filtros adicionales si están presentes
	if categoryId != "" {
		searchRequest.CategoryId = categoryId
	}

	if priceMinStr != "" {
		priceMin, err := strconv.ParseFloat(priceMinStr, 64)
		if err == nil {
			searchRequest.PriceMin = priceMin
		}
	}

	if priceMaxStr != "" {
		priceMax, err := strconv.ParseFloat(priceMaxStr, 64)
		if err == nil {
			searchRequest.PriceMax = priceMax
		}
	}

	// Agregar parámetros de ordenamiento
	searchRequest.SortBy = sortBy
	searchRequest.SortDirection = sortDirection

	// Llamar al servicio de búsqueda
	response, err := pc.productService.SearchProducts(searchRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

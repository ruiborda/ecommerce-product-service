package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruiborda/ecommerce-product-service/src/dto/product"
	"github.com/ruiborda/ecommerce-product-service/src/service"
	"github.com/ruiborda/ecommerce-product-service/src/service/impl"
	"github.com/ruiborda/go-jwt/src/domain/entity"
	"github.com/ruiborda/go-swagger-generator/src/openapi"
	"github.com/ruiborda/go-swagger-generator/src/openapi_spec/mime"
	"github.com/ruiborda/go-swagger-generator/src/swagger"
	"github.com/ruiborda/ecommerce-product-service/src/dto/auth"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryService: impl.NewCategoryServiceImpl(),
	}
}

var _ = swagger.Swagger().Path("/api/v1/categories").
	Post(func(operation openapi.Operation) {
		operation.Summary("Create a new category").
			OperationID("CreateCategory").
			Tag("CategoryController").
			Consume(mime.ApplicationJSON).
			Produces(mime.ApplicationJSON).
			BodyParameter(func(param openapi.Parameter) {
				param.Description("Category object that needs to be added to the system").
					Required(true).
					SchemaFromDTO(&product.CreateCategoryRequest{})
			}).
			Security("BearerAuth")
	}).Doc()

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var createCategoryRequest = &product.CreateCategoryRequest{}

	if err := c.BindJSON(createCategoryRequest); err != nil {
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
	response, err := cc.categoryService.CreateCategory(createCategoryRequest, authorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

var _ = swagger.Swagger().Path("/api/v1/categories/{id}").
	Put(func(operation openapi.Operation) {
		operation.Summary("Update an existing category").
			OperationID("UpdateCategory").
			Tag("CategoryController").
			Consume(mime.ApplicationJSON).
			Produces(mime.ApplicationJSON).
			PathParameter("id", func(param openapi.Parameter) {
				param.Description("ID of the category to update").
					Required(true).
					Type("string")
			}).
			BodyParameter(func(param openapi.Parameter) {
				param.Description("Category object with updated values").
					Required(true).
					SchemaFromDTO(&product.UpdateCategoryRequest{})
			}).
			Security("BearerAuth")
	}).Doc()

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var updateCategoryRequest = &product.UpdateCategoryRequest{}

	if err := c.BindJSON(updateCategoryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar al servicio para actualizar la categoría
	response, err := cc.categoryService.UpdateCategory(id, updateCategoryRequest)
	if err != nil {
		if err.Error() == "category not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

var _ = swagger.Swagger().Path("/api/v1/categories").
	Get(func(operation openapi.Operation) {
		operation.Summary("Get all categories").
			OperationID("GetCategories").
			Tag("CategoryController").
			Produces(mime.ApplicationJSON).
			Response(200, func(response openapi.Response) {
				response.Description("Successful operation").
					SchemaFromDTO(&product.GetCategoriesResponse{})
			}).
			Security("BearerAuth")
	}).Doc()

func (cc *CategoryController) GetCategories(c *gin.Context) {
	// Llamar al servicio para obtener todas las categorías
	response, err := cc.categoryService.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
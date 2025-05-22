package controller

import (
	"net/http"

	"github.com/ruiborda/ecommerce-product-service/src/dto/category"

	"github.com/gin-gonic/gin"
	"github.com/ruiborda/ecommerce-product-service/src/service"
	"github.com/ruiborda/ecommerce-product-service/src/service/impl"
	"github.com/ruiborda/go-swagger-generator/src/openapi"
	"github.com/ruiborda/go-swagger-generator/src/openapi_spec/mime"
	"github.com/ruiborda/go-swagger-generator/src/swagger"
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
					SchemaFromDTO(&category.CreateCategoryRequest{})
			}).
			Security("BearerAuth")
	}).Doc()

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var createCategoryRequest = &category.CreateCategoryRequest{}

	if err := c.BindJSON(createCategoryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar al servicio para crear la categoría
	response, err := cc.categoryService.CreateCategory(createCategoryRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

var _ = swagger.Swagger().Path("/api/v1/categories").
	Put(func(operation openapi.Operation) {
		operation.Summary("Update an existing category").
			OperationID("UpdateCategory").
			Tag("CategoryController").
			Consume(mime.ApplicationJSON).
			Produces(mime.ApplicationJSON).
			BodyParameter(func(param openapi.Parameter) {
				param.Description("Category object with updated values").
					Required(true).
					SchemaFromDTO(&category.UpdateCategoryRequest{})
			}).
			Security("BearerAuth")
	}).Doc()

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	var updateCategoryRequest = &category.UpdateCategoryRequest{}

	if err := c.BindJSON(updateCategoryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar al servicio para actualizar la categoría
	response, err := cc.categoryService.UpdateCategory(updateCategoryRequest)
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
			OperationID("GetAllCategories").
			Tag("CategoryController").
			Produces(mime.ApplicationJSON).
			Response(http.StatusOK, func(response openapi.Response) {
				response.Description("List of all categories").
					SchemaFromDTO(&[]*category.GetCategoriesResponse{})
			}).
			Security("BearerAuth")
	}).Doc()

func (cc *CategoryController) GetCategories(c *gin.Context) {
	// Usar el método que devuelve un puntero a un array
	response := cc.categoryService.GetAllCategoriesAsArray()
	// Desreferenciar el puntero para obtener el array
	c.JSON(http.StatusOK, *response)
}

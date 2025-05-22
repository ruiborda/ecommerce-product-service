package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ruiborda/ecommerce-product-service/src/controller"
	"github.com/ruiborda/ecommerce-user-service/src/middleware"
	"github.com/ruiborda/ecommerce-user-service/src/model"
)

func ApiRouter(router *gin.Engine) {
	productController := controller.NewProductController()
	categoryController := controller.NewCategoryController()

	router.POST(
		"/api/v1/products",
		middleware.RequireJWT(),
		middleware.RequirePermission(model.CreateUser),
		productController.CreateProduct,
	)

	router.GET(
		"/api/v1/products/:id",
		middleware.RequireJWT(),
		middleware.RequirePermission(model.CreateUser),
		productController.GetProductById,
	)

	router.PUT(
		"/api/v1/products/:id",
		middleware.RequireJWT(),
		middleware.RequirePermission(model.CreateUser),
		productController.UpdateProduct,
	)

	router.DELETE(
		"/api/v1/products/:id",
		middleware.RequireJWT(),
		middleware.RequirePermission(model.CreateUser),
		productController.DeleteProduct,
	)

	router.GET(
		"/api/v1/products/pages",
		middleware.RequireJWT(),
		middleware.RequirePermission(model.CreateUser),
		productController.GetProductsPaginated,
	)
	router.PUT(
		"/api/v1/products/:id/stock",
		middleware.RequireJWT(),
		middleware.RequirePermission(model.CreateUser),
		productController.AdjustProductStock,
	)

	router.GET(
		"/api/v1/products/search",
		middleware.RequireJWT(),
		middleware.RequirePermission(model.CreateUser),
		productController.SearchProducts,
	)

	// Rutas de categorías
	router.POST(
		"/api/v1/categories",
		middleware.RequireJWT(),
		middleware.RequirePermission(model.CreateUser),
		categoryController.CreateCategory,
	)
	
	router.PUT(
		"/api/v1/categories",
		middleware.RequireJWT(),
		middleware.RequirePermission(model.CreateUser),
		categoryController.UpdateCategory,
	)
	
	router.GET(
		"/api/v1/categories",
		middleware.RequireJWT(),
		middleware.RequirePermission(model.CreateUser),
		categoryController.GetCategories,
	)
}

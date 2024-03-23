package src

import (
	"backend-test/src/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *Dependency) *gin.Engine {
	r := gin.Default()

	product := r.Group("/products")
	product.POST("", handlers.StoreProduct(app.Service.ProductSVC))
	product.GET("/retrieved", handlers.GetPagincation(app.Service.ProductSVC))

	return r
}

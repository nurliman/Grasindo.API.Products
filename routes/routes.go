package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter function define routes endpoints
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"api": "grasindo"})
	})

	// endpoints "/v1"
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"api": "grasindo", "version": "1"})
		})

		// customers endpoints "/v1/customers"
		products := v1.Group("/products")
		{

		}
	}

	return router
}

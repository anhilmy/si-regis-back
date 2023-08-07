package main

import (
	docs "sireg/rest-api-kegiatan/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", index)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do piing
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} helloword
// @Router / [get]
func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", index)
	router.Run()
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

package kategori

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandler(router *gin.RouterGroup, service Service) {
	res := resource{service}

	router.GET("/", res.GetAll)
}

type resource struct {
	s Service
}

func (r resource) GetAll(c *gin.Context) {
	res, err := r.s.GetAll(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, res)
}

package kategori

import (
	"sireg/rest-api-kegiatan/internal/request"
	"sireg/rest-api-kegiatan/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterHandler(router *gin.RouterGroup, service Service) {
	res := resource{service}

	router.GET("/", res.GetAll)
	router.GET("/:kategoriId", res.Get)
	router.POST("/", res.Create)
	router.PATCH("/", res.Update)
	router.DELETE("/:kategoriId", res.Delete)

}

type resource struct {
	s Service
}

// Get all active kategori godoc
// @Summary get all active kategoris
// @Schemes
// @Tags kategori
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]response.Kategori}
// @Router /kategori/ [get]
func (r resource) GetAll(c *gin.Context) {
	kategoris, err := r.s.GetAll(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	res := response.SuccessResponse{
		Message: "success",
		Data:    kategoris,
	}
	c.JSON(200, res)
}

func (r resource) Get(c *gin.Context) {
	var kategoriId request.PathKategoriID
	if err := c.BindUri(&kategoriId); err != nil {
		// this should cancel the context because BindWith
		return
	}

	kategori, err := r.s.Get(c, kategoriId.KategoriId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	res := response.SuccessResponse{
		Message: "success",
		Data:    kategori,
	}

	c.JSON(200, res)
}
func (r resource) Create(c *gin.Context) {
	var request request.ReqKategori
	if err := c.BindWith(&request, binding.JSON); err != nil {
		// this should cancel the context because BindWith
		return
	}

	kategori, err := r.s.Create(c, &request)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	res := response.SuccessResponse{
		Message: "success",
		Data:    kategori,
	}

	c.JSON(200, res)
}

func (r resource) Update(c *gin.Context) {
	var reqBody request.ReqKategori
	if err := c.BindWith(&reqBody, binding.JSON); err != nil {
		// this should cancel the context because BindWith
		return
	}

	var kategoriId request.PathKategoriID
	if err := c.BindUri(&kategoriId); err != nil {
		return
	}

	kategori, err := r.s.Update(c, &reqBody, kategoriId.KategoriId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	res := response.SuccessResponse{
		Message: "success",
		Data:    kategori,
	}

	c.JSON(200, res)
}

func (r resource) Delete(c *gin.Context) {
	var kategoriId request.PathKategoriID
	if err := c.BindUri(&kategoriId); err != nil {
		return
	}

	err := r.s.Delete(c, kategoriId.KategoriId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	res := response.SuccessResponse{
		Message: "success",
	}

	c.JSON(200, res)
}

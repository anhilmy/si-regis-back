package kategori

import (
	"net/http"
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
	router.PATCH("/:kategoriId", res.Update)
	router.DELETE("/:kategoriId", res.Delete)

}

type resource struct {
	s Service
}

// Get all kategori godoc
// @Summary get all kategoris
// @Schemes
// @Tags kategori
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]response.Kategori}
// @Router /kategori/ [get]
func (r resource) GetAll(c *gin.Context) {
	kategoris, err := r.s.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	res := response.SuccessResponse{
		Message: "success",
		Data:    kategoris,
	}
	c.JSON(http.StatusOK, res)
}

// Get kategori godoc
// @Summary get kategori
// @Schemes
// @Tags kategori
// @Accept json
// @Produce json
// @Param kategoriId path int true "id kategori"
// @Success 200 {object} response.SuccessResponse{data=response.Kategori}
// @Router /kategori/{kategoriId} [get]
func (r resource) Get(c *gin.Context) {
	var kategoriId request.PathKategoriID
	if err := c.BindUri(&kategoriId); err != nil {
		// this should cancel the context because BindWith
		return
	}

	kategori, err := r.s.Get(c, kategoriId.KategoriId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := response.SuccessResponse{
		Message: "success",
		Data:    kategori,
	}

	c.JSON(http.StatusOK, res)
}

// Create kategori godoc
// @Summary create kategori
// @Schemes
// @Tags kategori
// @Accept json
// @Produce json
// @Param request body request.ReqKategori true "create new kategori"
// @Success 200 {object} response.SuccessResponse{data=response.Kategori}
// @Router /kategori/ [post]
func (r resource) Create(c *gin.Context) {
	var request request.ReqKategori
	if err := c.MustBindWith(&request, binding.JSON); err != nil {
		// this should cancel the context because BindWith
		c.JSON(http.StatusBadRequest, gin.H{"message": "400 bad request"})
		return
	}

	kategori, err := r.s.Create(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := response.SuccessResponse{
		Message: "success",
		Data:    kategori,
	}

	c.JSON(http.StatusOK, res)
}

// Update kategori godoc
// @Summary create kategori
// @Schemes
// @Tags kategori
// @Accept json
// @Produce json
// @Param request body request.ReqKategori true "create new kategori"
// @Param kategoriId path int true "id kategori"
// @Success 200 {object} response.SuccessResponse
// @Router /kategori/{kategoriId} [patch]
func (r resource) Update(c *gin.Context) {
	var reqBody request.ReqKategori
	if err := c.MustBindWith(&reqBody, binding.JSON); err != nil {
		// this should cancel the context because BindWith
		c.JSON(http.StatusBadRequest, gin.H{"message": "400 bad request"})
		return
	}

	var kategoriId request.PathKategoriID
	if err := c.BindUri(&kategoriId); err != nil {
		return
	}

	err := r.s.Update(c, &reqBody, kategoriId.KategoriId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := response.SuccessResponse{
		Message: "success",
	}

	c.JSON(http.StatusOK, res)
}

// Delete kategori godoc
// @Summary delete kategori
// @Schemes
// @Tags kategori
// @Accept json
// @Produce json
// @Param kategoriId path int true "id kategori"
// @Success 200 {object} response.SuccessResponse
// @Router /kategori/{kategoriId} [delete]
func (r resource) Delete(c *gin.Context) {
	var kategoriId request.PathKategoriID
	if err := c.BindUri(&kategoriId); err != nil {
		return
	}

	err := r.s.Delete(c, kategoriId.KategoriId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := response.SuccessResponse{
		Message: "success",
	}

	c.JSON(http.StatusOK, res)
}

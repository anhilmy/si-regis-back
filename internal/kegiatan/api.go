package kegiatan

import (
	"sireg/rest-api-kegiatan/internal/request"
	"sireg/rest-api-kegiatan/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterHandler(router *gin.RouterGroup, service Service) {
	res := resource{service}

	router.GET("/", res.GetAll)
	router.GET("/:id", res.Get)
	router.POST("/", res.Create)
	router.PATCH("/:id", res.Update)
	router.DELETE("/:id", res.Delete)

}

type resource struct {
	s Service
}

// Get all kegiatan godoc
// @Summary get all kegiatan
// @Schemes
// @Tags kegiatan
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]response.Kegiatan}
// @Router /kegiatan/ [get]
func (r resource) GetAll(c *gin.Context) {
	kategoris, err := r.s.GetAll(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	res := response.SuccessResponse{
		Message: "success",
		Data:    kategoris,
	}
	c.JSON(200, res)
}

// Get kegiatan godoc
// @Summary get kegiatan
// @Schemes
// @Tags kegiatan
// @Accept json
// @Produce json
// @Param id path int true "id kegiatan"
// @Success 200 {object} response.SuccessResponse{data=response.Kegiatan}
// @Router /kegiatan/{id} [get]
func (r resource) Get(c *gin.Context) {
	var kegiatanId request.PathKegiatanID
	if err := c.BindUri(&kegiatanId); err != nil {
		// this should cancel the context because BindWith
		return
	}

	kegiatan, err := r.s.Get(c, kegiatanId.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := response.SuccessResponse{
		Message: "success",
		Data:    kegiatan,
	}

	c.JSON(200, res)
}

// Create kegiatan godoc
// @Summary create kegiatan
// @Schemes
// @Tags kegiatan
// @Accept json
// @Produce json
// @Param request body request.ReqKegiatan true "create new kegiatan"
// @Success 200 {object} response.SuccessResponse{data=response.Kegiatan}
// @Router /kegiatan/ [post]
func (r resource) Create(c *gin.Context) {
	var request request.ReqKegiatan
	if err := c.MustBindWith(&request, binding.JSON); err != nil {
		// this should cancel the context because BindWith
		return
	}

	kegiatan, err := r.s.Create(c, &request)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := response.SuccessResponse{
		Message: "success",
		Data:    kegiatan,
	}

	c.JSON(200, res)
}

// Update kegiatan godoc
// @Summary create kegiatan
// @Schemes
// @Tags kegiatan
// @Accept json
// @Produce json
// @Param request body request.ReqKegiatan true "create new kegiatan"
// @Param id path int true "id kegiatan"
// @Success 200 {object} response.SuccessResponse{data=response.Kegiatan}
// @Router /kegiatan/{id} [patch]
func (r resource) Update(c *gin.Context) {
	var reqBody request.ReqKegiatan
	if err := c.MustBindWith(&reqBody, binding.JSON); err != nil {
		// this should cancel the context because BindWith
		return
	}

	var kegiatanId request.PathKegiatanID
	if err := c.BindUri(&kegiatanId); err != nil {
		return
	}

	kegiatan, err := r.s.Update(c, &reqBody, kegiatanId.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := response.SuccessResponse{
		Message: "success",
		Data:    kegiatan,
	}

	c.JSON(200, res)
}

// Delete kegiatan godoc
// @Summary delete kegiatan
// @Schemes
// @Tags kegiatan
// @Accept json
// @Produce json
// @Param id path int true "id kegiatan"
// @Success 200 {object} response.SuccessResponse
// @Router /kegiatan/{id} [delete]
func (r resource) Delete(c *gin.Context) {
	var kegiatanId request.PathKegiatanID
	if err := c.BindUri(&kegiatanId); err != nil {
		return
	}

	err := r.s.Delete(c, kegiatanId.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := response.SuccessResponse{
		Message: "success",
	}

	c.JSON(200, res)
}

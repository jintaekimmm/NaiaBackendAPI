package controllers

import (
	"github.com/99-66/NaiaBackendApi/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WINApi struct {
	WINService services.WINService
}

func ProvideWINApi(w services.WINService) WINApi {
	return WINApi{WINService: w}
}

// List godoc
// @Summary 이슈 단어목록 API
// @Description 현재시간 기준 6시간 전까지의 상위 이슈 단어 30개를 반환한다
// @Tags WhatIssueNow
// @Accept application/json
// @Produce application/json
// @Success 200 {array} models.WINListResponse
// @Failure 500 {object} config.APIError
// @Router /list [get]
func (w *WINApi) List(c *gin.Context) {
	winResp, err := w.WINService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": winResp,
	})
}

// FindWordToTagPercent godoc
// @Summary  단어별 태그 점유율 API
// @Description 특정 단어의 발생지(태그) 점유율을 반환한다
// @Tags WhatIssueNow
// @Accept application/json
// @Produce application/json
// @Param word path string true "Word"
// @Success 200 {object} models.WINTagResponse
// @Failure 500 {object} config.APIError
// @Router /tag/{word} [get]
func (w *WINApi) FindWordToTagPercent(c *gin.Context) {
	p := c.Param("word")
	tagResp, err := w.WINService.FindWordToTagPercent(p)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": tagResp,
	})
}
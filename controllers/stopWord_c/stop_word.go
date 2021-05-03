/*
Unused file
*/

package stopWord_c

import (
	"github.com/99-66/NaiaBackendApi/libs"
	"github.com/99-66/NaiaBackendApi/models/win_m"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetWords 집계시에 제외되는 불용어를 등록한다
// List godoc
// @Summary 불용어 단어등록 API
// @Description ElasticSearch 집계시에 제외되는 불용어 단어를 등록한다
// @Tags StopWords
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string
// @Failure 400 {object} libs.APIError
func SetWords(c *gin.Context) {
	var wStopWord win_m.WStopWord
	err := c.ShouldBindJSON(&wStopWord)
	if err != nil {
		libs.ErrResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	err = wStopWord.Set()
	if err != nil {
		libs.ErrResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": wStopWord.Word,
	})
}

// GetStopWords 집계시에 제외되는 불용어를 목록을 반환한다
// List godoc
// @Summary 불용어 단어목록 API
// @Description ElasticSearch 집계시에 제외되는 불용어 목록을 반환한다
// @Tags StopWords
// @Accept application/json
// @Produce application/json
// @Success 200 {array} string
// @Failure 400 {object} libs.APIError
// @Router /stopwords [get]
func GetStopWords(c *gin.Context) {
	var wStopWord win_m.WStopWord
	words, err := wStopWord.List()
	if err != nil {
		libs.ErrResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": words,
	})
}

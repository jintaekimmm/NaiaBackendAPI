/*
Unused file
*/

package controllers

import (
	"github.com/99-66/NaiaBackendApi/libs"
	"github.com/99-66/NaiaBackendApi/models/win"
	"github.com/99-66/NaiaBackendApi/services"
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
// @Failure 400 {object} responses.Error
func SetWords(c *gin.Context) {
	var stopWord win.StopWord
	err := c.ShouldBindJSON(&stopWord)
	if err != nil {
		libs.ErrResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	err = stopWord.Set()
	if err != nil {
		libs.ErrResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": stopWord.Word,
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
// @Failure 400 {object} responses.Error
// @Router /stopwords [get]
func GetStopWords(c *gin.Context) {
	words, err := services.StopWords()
	if err != nil {
		libs.ErrResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": words,
	})
}

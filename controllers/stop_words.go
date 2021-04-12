/*
Unused file
*/

package controllers

import (
	"github.com/99-66/NaiaBackendApi/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Body struct {
	Word string `json:"word"`
}

// SetStopWords 불용어를 등록한다
func SetStopWords(key, stopWord string) error {
	result := config.REDIS.SAdd(config.CTX, key, stopWord)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

// SetWords 집계시에 제외되는 불용어를 등록한다
// List godoc
// @Summary 불용어 단어등록 API
// @Description ElasticSearch 집계시에 제외되는 불용어 단어를 등록한다
// @Tags StopWords
// @Accept application/json
// @Produce application/json
// @Success 200 {string} Body.Word
// @Failure 400 {object} config.APIError
func SetWords(c *gin.Context) {
	var body Body
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
		return
	}

	err = SetStopWords(os.Getenv("REDIS_KEY"), body.Word)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": body.Word,
	})

}

package controllers

import (
	"github.com/99-66/NaiaBackendApi/libs"
	"github.com/99-66/NaiaBackendApi/services"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

// List godoc
// @Summary 이슈 단어목록 API
// @Description 현재시간 기준 3시간 전까지의 상위 이슈 단어 30개를 반환한다
// @Tags WhatIssueNow
// @Accept application/json
// @Produce application/json
// @Param count query int false "Word list count by count"
// @Param f query string false "Word Filtering by f"
// @Success 200 {array} win.List
// @Failure 500 {object} responses.Error
// @Router /list [get]
func List(c *gin.Context) {
	p := c.Query("count")
	filter := c.Query("f")

	count, err := strconv.Atoi(p)
	if err != nil {
		count = 30
	}
	count = int(math.Abs(float64(count)))
	if count > 100 {
		count = 30
	}

	resp, err := services.WordList(count, filter)

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	})
}

// WordToTagPercent godoc
// @Summary  단어별 태그 점유율 API
// @Description 특정 단어의 발생지(태그) 점유율을 반환한다
// @Tags WhatIssueNow
// @Accept application/json
// @Produce application/json
// @Param word path string true "Word"
// @Success 200 {object} win.Tag
// @Failure 500 {object} responses.Error
// @Router /tag/w/{word} [get]
func WordToTagPercent(c *gin.Context) {
	p := c.Param("word")
	resp, err := services.WordToTagPercent(p)

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	})
}

// ListForWordCloud godoc
// @Summary 워드 클라우드를 위한 이슈 단어목록 API
// @Description 현재시간 기준 3시간 전까지의 상위 이슈 단어 60개를 반환한다
// @Tags WhatIssueNow
// @Accept application/json
// @Produce application/json
// @Param count query int false "Word list count by count"
// @Param f query string false "Word Filtering by f"
// @Param count query int false "워드클라우드 단어 개수(최대 100개)"
// @Success 200 {array} win.WordCloud
// @Failure 500 {object} responses.Error
// @Router /wordcloud [get]
func ListForWordCloud(c *gin.Context) {
	p := c.Query("count")
	filter := c.Query("f")

	count, err := strconv.Atoi(p)
	if err != nil {
		count = 60
	}
	count = int(math.Abs(float64(count)))
	if count > 100 {
		count = 60
	}

	resp, err := services.WordCloud(count, filter)

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	})
}

// WordToFindRelated godoc
// @Summary  특정 단어와 연관된 단어 목록 API
// @Description 특정 단어와 관련된 다른 단어들을 반환한다
// @Tags WhatIssueNow
// @Accept application/json
// @Produce application/json
// @Param word path string true "Word"
// @Success 200 {object} win.WordsResponse
// @Failure 500 {object} responses.Error
// @Router /related/w/{word} [get]
func WordToFindRelated(c *gin.Context) {
	p := c.Param("word")

	wordsResp, err := services.RelatedWords(p)

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": wordsResp,
	})
}

// WordToFindRelatedTweets godoc
// @Summary  특정 단어의 최근 트윗 목록 API
// @Description 특정 단어의 관련된 최근 트윗 목록 100개를 반환한다
// @Tags WhatIssueNow
// @Accept application/json
// @Produce application/json
// @Param word path string true "Word"
// @Success 200 {object} win.RelatedTweets
// @Failure 500 {object} responses.Error
// @Router /related/list/{word} [get]
func WordToFindRelatedTweets(c *gin.Context) {
	p := c.Param("word")

	tweetsResp, err := services.RelatedTweets(p)

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": tweetsResp,
	})
}

// TagCount godoc
// @Summary 태그별 단어 수 API
// @Description 현재시간 기준 3시간 전까지의 태그별 단어 수를 반환한다
// @Tags WhatIssueNow
// @Accept application/json
// @Produce application/json
// @Success 200 {array} win.TagCount
// @Failure 500 {object} responses.Error
// @Router /tag/count [get]
func TagCount(c *gin.Context) {
	resp, err := services.WordToTagCounts()

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	})
}

// WordToCount godoc
// @Summary 일주일간 수집한 단어의 수 API
// @Description 현재시간 기준 7일전까지의 수집한 단어의 수를 일별로 반환한다
// @Tags WhatIssueNow
// @Accept application/json
// @Produce application/json
// @Success 200 {array} win.WordCount
// @Failure 500 {object} responses.Error
// @Router /word/count [get]
func WordToCount(c *gin.Context) {
	resp, err := services.WeeklyWordCount()

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	})
}

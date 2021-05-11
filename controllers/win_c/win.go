package win_c

import (
	"github.com/99-66/NaiaBackendApi/libs"
	"github.com/99-66/NaiaBackendApi/models/win_m"
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
// @Success 200 {array} win_m.WList
// @Failure 500 {object} libs.APIError
// @Router /list [get]
func List(c *gin.Context) {
	var wList win_m.WList
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

	resp, err := wList.List(count, filter)

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
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
// @Success 200 {object} win_m.WTag
// @Failure 500 {object} libs.APIError
// @Router /tag/{word} [get]
func WordToTagPercent(c *gin.Context) {
	p := c.Param("word")
	var wTag win_m.WTag
	resp, err := wTag.WordToTagPercent(p)

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
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
// @Success 200 {array} win_m.WordCloudList
// @Failure 500 {object} libs.APIError
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

	var wWordList win_m.WWordCloud
	resp, err := wWordList.List(count, filter)

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
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
// @Success 200 {object} win_m.WordsResponse
// @Failure 500 {object} libs.APIError
// @Router /related/{word} [get]
func WordToFindRelated(c *gin.Context) {
	p := c.Param("word")

	var wRelatedWords win_m.WRelatedWords
	wordsResp, err := wRelatedWords.RelatedWords(p)

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
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
// @Success 200 {object} win_m.RTweets
// @Failure 500 {object} libs.APIError
// @Router /related/list/{word} [get]
func WordToFindRelatedTweets(c *gin.Context) {
	p := c.Param("word")

	var wRelatedTweets win_m.WRelatedTweets
	tweetsResp, err := wRelatedTweets.RelatedTweets(p)

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": tweetsResp,
	})
}

// TagCount godoc
// @Summary 태그별 단어 수 API
// @Description 현재시간 기준 3시간 전까지의 태그별 단어 수를 반환한다
// @Tags WhatIssueNow
// @Accept application/json
// @Produce application/json
// @Success 200 {array} win_m.WTagCount
// @Failure 500 {object} libs.APIError
// @Router /tag/count [get]
func TagCount(c *gin.Context) {
	var tagCount win_m.WTagCount

	resp, err := tagCount.Counts()

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": resp,
	})
}

// WordToCount godoc
// @Summary 일주일간 수집한 단어의 수 API
// @Description 현재시간 기준 7일전까지의 수집한 단어의 수를 일별로 반환한다
// @Tags WhatIssueNow
// @Accept application/json
// @Produce application/json
// @Success 200 {array} win_m.WWordCount
// @Failure 500 {object} libs.APIError
// @Router /word/count [get]
func WordToCount(c *gin.Context) {
	var wordCount win_m.WWordCount

	resp, err := wordCount.Counts()

	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": resp,
	})
}

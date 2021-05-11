package win_r

import (
	"github.com/99-66/NaiaBackendApi/controllers/stopWord_c"
	"github.com/99-66/NaiaBackendApi/controllers/win_c"
	"github.com/gin-gonic/gin"
)

// InitWINRoutes WhatIssueNow API Routes
func InitWINRoutes(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	{
		v1 := api.Group("/1")
		{
			v1.GET("/list", win_c.List)
			v1.GET("/wordcloud", win_c.ListForWordCloud)
			v1.GET("/tag/w/:word", win_c.WordToTagPercent)
			v1.GET("/tag/count", win_c.TagCount)
			v1.GET("/word/count", win_c.WordToCount)
			v1.GET("/stopwords", stopWord_c.GetStopWords)
			//v1.POST("/stopwords", stopWord_c.SetWords)
			related := v1.Group("/related")
			{
				related.GET("/w/:word", win_c.WordToFindRelated)
				related.GET("/list/:word", win_c.WordToFindRelatedTweets)
			}
		}
	}

	return r
}

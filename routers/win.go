package routers

import (
	"github.com/99-66/NaiaBackendApi/controllers"
	"github.com/gin-gonic/gin"
)

// initWINRoutes WhatIssueNow API Routes
func initWINRoutes(v1 *gin.RouterGroup) {
	v1.GET("/list", controllers.List)
	v1.GET("/wordcloud", controllers.ListForWordCloud)
	v1.GET("/word/count", controllers.WordToCount)
	v1.GET("/stopwords", controllers.GetStopWords)
	//v1.POST("/stopwords", stopWord_c.SetWords)

	tag := v1.Group("/tag")
	{
		tag.GET("/w/:word", controllers.WordToTagPercent)
		tag.GET("/count", controllers.TagCount)
	}

	related := v1.Group("/related")
	{
		related.GET("/w/:word", controllers.WordToFindRelated)
		related.GET("/list/:word", controllers.WordToFindRelatedTweets)
	}
}

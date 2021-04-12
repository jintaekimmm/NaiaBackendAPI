package routes

import (
	"github.com/99-66/NaiaBackendApi/controllers"
	"github.com/gin-gonic/gin"
)

// WINRoute WhatIssueNow API Routes
func WINRoute(r *gin.Engine, win controllers.WINApi) *gin.Engine {
	api := r.Group("/api")
	{
		v1 := api.Group("/1")
		{
			v1.GET("/list", win.List)
			v1.GET("/tag/:word", win.FindWordToTagPercent)
			v1.GET("/stopwords", win.GetStopWords)
		}
	}

	return r
}

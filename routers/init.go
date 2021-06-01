package routers

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) *gin.Engine {
	initSwaggerRoutes(r)

	api := r.Group("/api")
	{
		v1 := api.Group("/1")
		{
			initWINRoutes(v1)
		}
	}

	return r
}

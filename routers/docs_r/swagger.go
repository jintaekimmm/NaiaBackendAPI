package docs_r

import (
	_ "github.com/99-66/NaiaBackendApi/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwaggerRoutes(r *gin.Engine) {
	// Swagger settings
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

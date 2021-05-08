package docs_r

import (
	_ "github.com/99-66/NaiaBackendApi/docs"
	"github.com/gin-gonic/gin"
)

func InitSwaggerRoutes(r *gin.Engine) {
	// Swagger settings
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

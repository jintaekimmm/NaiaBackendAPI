package routers

import (
	_ "github.com/99-66/NaiaBackendApi/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initSwaggerRoutes(r *gin.Engine) {
	// Swagger settings
	// SWAGGER 변수가 설정되어 있다면 스웨거는 사용하지 못한다(Return 404)
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER"))
}

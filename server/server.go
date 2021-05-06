package server

import (
	"github.com/99-66/NaiaBackendApi/routers/docs_r"
	"github.com/99-66/NaiaBackendApi/routers/win_r"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	return RunAPIWithMiddleware(address)
}

func RunAPIWithMiddleware(address string) error {
	r := gin.Default()

	// Set CORS Middlewares
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	r.Use(cors.New(corsConf))

	// Set Routes
	win_r.InitWINRoutes(r)
	docs_r.InitSwaggerRoutes(r)

	return r.Run(address)
}

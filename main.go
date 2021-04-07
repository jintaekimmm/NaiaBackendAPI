package main

import (
	"github.com/99-66/NaiaBackendApi/config"
	_ "github.com/99-66/NaiaBackendApi/docs"
	"github.com/99-66/NaiaBackendApi/routes"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

// @title WhatIssueNow API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/1
// @query.collection.format multi

// @title WhatIssueNow API
// @description WhatIssueNow Service API
// @schemes http https

type dbHandlers struct {
	es *elasticsearch.Client
}

func main() {
	var err error
	var dbHandler dbHandlers

	dbHandler.es, err = config.InitElasticSearch()
	if err != nil {
		panic(err)
	}

	r := initRoutes(dbHandler)
	log.Fatal(r.Run())
}

func initRoutes(dbHandler dbHandlers) *gin.Engine {
	r := gin.Default()

	// CORS allows all origins
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	r.Use(cors.New(conf))

	// Project WhatIssueNow routes
	winApi := initWINApi(dbHandler.es)
	routes.WINRoute(r, winApi)

	// Swagger settings
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

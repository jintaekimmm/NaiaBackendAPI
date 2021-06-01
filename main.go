package main

import (
	"fmt"
	"github.com/99-66/NaiaBackendApi/repositories"
	"github.com/99-66/NaiaBackendApi/server"
	"log"
	"os"
)

// @title WhatIssueNow API
// @version 1.2
// @description What Issue Now? Service Api Docs

// @contact.name Jintae, kim
// @contact.url http://whatissuenow.com
// @contact.email 6199@outlook.kr

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.whatissuenow.com
// @BasePath /api/1
// @query.collection.format multi

// @title WhatIssueNow API
// @description WhatIssueNow Service API
// @schemes http https
func main() {
	repositories.Init()
	if repositories.Connections.DB != nil {
		db := repositories.Connections.DB
		sqlDB, ok := db.DB()
		if ok != nil {
			defer sqlDB.Close()
		}
	}

	addr := "0.0.0.0"
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", addr, port)

	log.Fatal(server.RunAPI(address))
}

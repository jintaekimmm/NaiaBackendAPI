package main

import (
	"github.com/99-66/NaiaBackendApi/controllers"
	"github.com/99-66/NaiaBackendApi/repositories"
	"github.com/99-66/NaiaBackendApi/services"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/google/wire"
)

func initWINApi(es *elasticsearch.Client) controllers.WINApi {
	wire.Build(
		repositories.ProvideWINRepository,
		services.ProvideWINService,
		controllers.ProvideWINApi)

	return controllers.WINApi{}
}
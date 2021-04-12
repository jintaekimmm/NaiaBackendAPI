package main

import (
	"github.com/99-66/NaiaBackendApi/controllers"
	"github.com/99-66/NaiaBackendApi/repositories"
	"github.com/99-66/NaiaBackendApi/services"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

func initWINApi(es *elasticsearch.Client, redis *redis.Client) controllers.WINApi {
	wire.Build(
		repositories.ProvideWINRepository,
		services.ProvideWINService,
		controllers.ProvideWINApi)

	return controllers.WINApi{}
}

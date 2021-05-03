package repositories

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type connections struct {
	DB    *gorm.DB
	ES    *elasticsearch.Client
	REDIS *redis.Client
}

var Connections connections

// ElasticSearch Meta Information
type esMeta struct {
	Index string
}

var ESMeta esMeta

func Init() {
	//initDB()
	initRedis()
	initElasticSearch()
}

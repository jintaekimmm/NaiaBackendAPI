package repositories

import (
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT"`
	DB       int    `env:"REDIS_DB"`
	Password string `env:"REDIS_PASSWORD"`
}

func newRedis(config RedisConfig) (*redis.Client, error) {
	dsn := fmt.Sprintf("%s:%d", config.Host, config.Port)

	r := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: config.Password,
		DB:       config.DB,
	})

	return r, nil
}

func initRedis() {
	redisConfig := RedisConfig{}
	if err := env.Parse(&redisConfig); err != nil {
		panic(errors.New("could not load redis configuration"))
	}

	r, err := newRedis(redisConfig)
	if err != nil {
		panic(err)
	}

	Connections.REDIS = r
}

package repositories

import (
	"errors"
	"github.com/caarlos0/env"
	elastic "github.com/elastic/go-elasticsearch/v7"
)

type ElasticSearchConfig struct {
	Host     []string `env:"ELS_HOST" envSeparator:","`
	User     string   `env:"ELS_USER"`
	Password string   `env:"ELS_PASSWORD"`
	Index    string   `env:"ELS_INDEX"`
}

func newElasticSearch(config ElasticSearchConfig) (*elastic.Client, error) {
	cfg := elastic.Config{
		Addresses: config.Host,
		Username:  config.User,
		Password:  config.Password,
	}

	es, err := elastic.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return es, nil
}

func initElasticSearch() {
	elsConfig := ElasticSearchConfig{}
	if err := env.Parse(&elsConfig); err != nil {
		panic(errors.New("cloud not load elasticsearch environment variables"))
	}

	ESMeta.Index = elsConfig.Index

	es, err := newElasticSearch(elsConfig)
	if err != nil {
		panic(err)
	}

	Connections.ES = es
}

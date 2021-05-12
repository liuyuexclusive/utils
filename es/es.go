package es

import (
	"github.com/liuyuexclusive/utils/config"

	elastic "github.com/olivere/elastic/v7"
)

func Client() (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL(config.MustGet().ElasticURL), elastic.SetSniff(false))
}

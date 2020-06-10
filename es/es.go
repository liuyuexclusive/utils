package es

import (
	"github.com/liuyuexclusive/utils/appconfig"

	elastic "github.com/olivere/elastic/v7"
)

func Client() (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL(appconfig.MustGet().ElasticURL), elastic.SetSniff(false))
}

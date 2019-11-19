package elasticutil

import (
	"github.com/liuyuexclusive/utils/appconfigutil"

	"github.com/olivere/elastic/v7"
)

func Client() (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL(appconfigutil.MustGet().ElasticURL), elastic.SetSniff(false))
}

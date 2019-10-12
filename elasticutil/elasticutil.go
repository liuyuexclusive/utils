package elasticutil

import (
	"future/util/configutil"

	"github.com/olivere/elastic/v7"
)

func Client() (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL(configutil.MustGet().ElasticURL), elastic.SetSniff(false))
}

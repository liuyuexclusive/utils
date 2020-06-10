package log

import (
	"fmt"

	"github.com/liuyuexclusive/utils/appconfig"
	"github.com/liuyuexclusive/utils/es"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	elogrus "gopkg.in/sohlich/elogrus.v7"
)

// LogToES logrus to elastic
func LogToES(index string) error {
	client, err := es.Client()
	if err != nil {
		logrus.Error(err)
		return err
	}
	hook, err := elogrus.NewElasticHook(client, appconfig.MustGet().ElasticURL, logrus.InfoLevel, index)
	if err != nil {
		return fmt.Errorf("fail of log to elastic : %w", err)
	}
	logrus.AddHook(hook)
	gin.DefaultWriter = logrus.StandardLogger().Writer()
	logrus.Info("seccessfully of log to elastic")
	return nil
}

package logutil

import (
	"fmt"

	"github.com/liuyuexclusive/utils/configutil"
	"github.com/liuyuexclusive/utils/elasticutil"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	elogrus "gopkg.in/sohlich/elogrus.v7"
)

// LogToElastic logrus to elastic
func LogToES(index string) error {
	client, err := elasticutil.Client()
	if err != nil {
		logrus.Error(err)
		return err
	}
	hook, err := elogrus.NewElasticHook(client, configutil.MustGet().ElasticURL, logrus.InfoLevel, index)
	if err != nil {
		return fmt.Errorf("写入elistic日志失败: %w", err)
	}
	logrus.AddHook(hook)
	gin.DefaultWriter = logrus.StandardLogger().Writer()
	logrus.Info("写入elistic日志成功!")
	return nil
}

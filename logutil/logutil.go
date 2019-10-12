package logutil

import (
	"fmt"
	"time"
	"utils/configutil"
	"utils/elasticutil"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	elogrus "gopkg.in/sohlich/elogrus.v7"
)

// LogToElastic logrus to elastic
func LogToElastic() error {
	client, err := elasticutil.Client()
	if err != nil {
		log.Error(err)
	}
	hook, err := elogrus.NewElasticHook(client, configutil.MustGet().ElasticURL, logrus.InfoLevel, "log-future-"+time.Now().Format("20060102"))
	if err != nil {
		return fmt.Errorf("写入elistic日志失败: %w", err)
	}
	logrus.AddHook(hook)
	gin.DefaultWriter = logrus.StandardLogger().Writer()
	logrus.Info("写入elistic日志成功!")
	return nil
}

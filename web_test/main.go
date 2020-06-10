package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liuyuexclusive/utils/web"
	"github.com/sirupsen/logrus"
)

type start struct{}

func (s *start) Start(engine *gin.Engine) {

}

func main() {
	if err := web.Startup("go.micro.api.basic", new(start), func(options *web.Options) {
		options.IsLogToES = false
		options.IsTrace = false
		options.IsMonitor = false
		options.IsRateLimite = false
	}); err != nil {
		logrus.Fatal(err)
	}
}

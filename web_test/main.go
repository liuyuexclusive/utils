package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liuyuexclusive/utils/webutil"
	"github.com/sirupsen/logrus"
)

type start struct{}

func (s *start) Start(engine *gin.Engine) {

}

func main() {
	if err := webutil.Startup("go.micro.api.basic", new(start), func(options *webutil.Options) {
		options.IsLogToES = false
		options.IsTrace = false
		options.IsMonitor = false
		options.IsRateLimite = false
	}); err != nil {
		logrus.Fatal(err)
	}
}

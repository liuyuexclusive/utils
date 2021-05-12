package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuexclusive/utils/log"
	"go.uber.org/zap/zapcore"
)

func main() {
	log.Init("./log", zapcore.DebugLevel)

	engine := gin.Default()

	log.Gin(engine)

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	engine.Run(":8080")
}

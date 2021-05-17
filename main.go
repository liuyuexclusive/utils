package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/yuexclusive/utils/docs"
	"github.com/yuexclusive/utils/logger"
	"github.com/yuexclusive/utils/web"
)

// Auth godoc
// @Summary
// @Description
// @Tags 获取token
// @Accept  json
// @Produce  json
// @Success 200 {string} string "output"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /test/hello [get]
func Hello(c *gin.Context) {
	c.String(http.StatusOK, "good luck")
}

func main() {
	logger.Init()
	engine := gin.Default()

	web.Prometheus(engine)
	web.Swagger(engine)

	engine.Use(web.AllowOrigin())
	engine.Use(web.Tracer())
	engine.Use(web.RateLimite(time.Second, 5))

	engine.GET("/test/hello", Hello)

	engine.Run(":8080")
}

package web

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/yuexclusive/utils/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	p "github.com/prometheus/client_golang/prometheus/promhttp"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/yuexclusive/utils/config"
)

const contextTracerKey = "Tracer-context"

var l = logger.Single()

// sf sampling frequency
var sf = 100

func init() {
	rand.Seed(time.Now().Unix())
}

// SetSamplingFrequency 设置采样频率
// 0 <= n <= 100
func SetSamplingFrequency(n int) {
	sf = n
}

// Prometheus
func Prometheus(engine *gin.Engine) {
	path := "/metrics"
	engine.GET(path, func(c *gin.Context) {
		p.Handler().ServeHTTP(c.Writer, c.Request)
	})
	l.Info("open prometheus", zap.String("path", path))
}

// Swagger
func Swagger(engine *gin.Engine) {
	var url, path string
	if name := config.MustGet().Name; name == "" {
		url = fmt.Sprintf("http://%s:%s/swagger/doc.json", config.MustGet().IP, config.MustGet().Port)
		path = "/swagger/*any"
	} else {
		url = fmt.Sprintf("http://%s:%s/%s/swagger/doc.json", config.MustGet().IP, config.MustGet().Port, config.MustGet().Name) // The url pointing to API definition
		path = fmt.Sprintf("/%s/swagger/*any", name)
	}
	engine.GET(path, ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(url)))

	l.Info("open swagger", zap.String("url", url), zap.String("path", path))
}

// AllowOrigin
func AllowOrigin() gin.HandlerFunc {
	l.Info("open allow origin")
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET,OPTIONS")
		c.Header("Access-Control-Max-Age", "86400") // 缓存请求信息 单位为秒
		// c.Header("Access-Control-Allow-Credentials", "false")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// RateLimite
func RateLimite(duration time.Duration, capacity int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(duration, capacity)
	l.Info("open rate limit", zap.String("duration", duration.String()), zap.Int64("capacity", capacity))
	return func(c *gin.Context) {
		available := bucket.TakeAvailable(1)
		if available <= 0 {
			c.JSON(http.StatusBadRequest, "visit too frequently, please try again later")
			c.Abort()
		}
	}
}

// contextWithSpan
func contextWithSpan(c *gin.Context) (ctx context.Context) {
	v, exist := c.Get(contextTracerKey)
	if exist {
		if r, ok := v.(context.Context); ok {
			ctx = r
			return
		}
	}
	ctx = context.Background()
	return
}

// Tracer
func Tracer() gin.HandlerFunc {
	l.Info("open trace")

	return func(c *gin.Context) {
		md := make(map[string]string)
		spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		sp := opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
		defer sp.Finish()

		if err := opentracing.GlobalTracer().Inject(sp.Context(),
			opentracing.TextMap,
			opentracing.TextMapCarrier(md)); err != nil {
			l.Fatal(err.Error())
		}

		ctx := context.TODO()
		ctx = opentracing.ContextWithSpan(ctx, sp)
		// ctx = metadata.NewContext(ctx, md)
		c.Set(contextTracerKey, ctx)

		c.Next()

		statusCode := c.Writer.Status()
		ext.HTTPStatusCode.Set(sp, uint16(statusCode))
		ext.HTTPMethod.Set(sp, c.Request.Method)
		ext.HTTPUrl.Set(sp, c.Request.URL.EscapedPath())
		if statusCode >= http.StatusInternalServerError {
			ext.Error.Set(sp, true)
		} else if rand.Intn(100) > sf {
			ext.SamplingPriority.Set(sp, 0)
		}
	}
}

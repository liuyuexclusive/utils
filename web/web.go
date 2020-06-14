package web

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/liuyuexclusive/utils/appconfig"
	"github.com/liuyuexclusive/utils/log"
	"github.com/liuyuexclusive/utils/trace"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/broker/nats"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
	"github.com/sirupsen/logrus"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/micro/go-micro/v2/metadata"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	p "github.com/prometheus/client_golang/prometheus/promhttp"
)

func Prometheus(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == fmt.Sprintf("/%s/metrics", name) {
			p.Handler().ServeHTTP(c.Writer, c.Request)
		}
		c.Next()
	}
}

// UseSwagger UseSwagger
func UseSwagger(path string, url string, router *gin.Engine) {
	router.GET(path, ginSwagger.CustomWrapHandler(&ginSwagger.Config{
		URL: url,
	}, swaggerFiles.Handler))
}

func AllowOrigin() gin.HandlerFunc {
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

// RateLimite 限流
func RateLimite() gin.HandlerFunc {
	bucket := ratelimit.NewBucket(time.Millisecond*5, 1000)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) <= 0 {
			c.JSON(400, "网络繁忙，请稍候再试")
			c.Abort()
		}
	}
}

func ReadBody(c *gin.Context, data interface{}) bool {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		Bad(c, err)
		return false
	}
	if err := json.Unmarshal(bytes, data); err != nil {
		Bad(c, err)
		return false
	}
	return true
}

func Bad(c *gin.Context, err error) {
	if err != nil {
		logrus.Error(err)
		c.JSON(400, err)
	}
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

type Starter interface {
	Start(engine *gin.Engine)
}

type Options struct {
	IsOpenSwagger bool
	// 是否记录日志到ES 默认为false
	IsLogToES bool
	// 是否使用opentrace(jaeger)
	IsTrace bool
	//是否允许跨域 默认为false,因为micro api默认做了跨域处理
	IsAllowOrigin bool
	//是否限流 默认为true
	IsRateLimite bool
	//端口 默认为空
	Port string
	//是否监控
	IsMonitor bool
}

type Option func(ops *Options)

func Startup(name string, starter Starter, opts ...Option) error {
	options := &Options{
		IsOpenSwagger: false,
		IsLogToES:     false,
		IsTrace:       false,
		IsAllowOrigin: false,
		IsRateLimite:  false,
		Port:          "",
	}

	for _, opt := range opts {
		opt(options)
	}

	if options.IsLogToES {
		log.LogToES(name)
	}

	if name == "" {
		return errors.New("请输入服务名称")
	}

	config := appconfig.MustGet()
	client.DefaultClient.Init(client.Broker(nats.NewBroker(broker.Addrs(config.NatsAddress))))

	webOptions := []web.Option{
		web.Name(name),
		web.Version("latest"),
		web.Registry(etcd.NewRegistry(registry.Addrs(config.ETCDAddress))),
		web.RegisterTTL(time.Second * 30),
		web.RegisterInterval(time.Second * 15),
	}
	if options.Port != "" {
		webOptions = append(webOptions, web.Address(options.Port))
	}

	service := web.NewService(
		webOptions...,
	)

	if err := service.Init(); err != nil {
		return fmt.Errorf("服务初始化失败:%w", err)
	}

	router := gin.Default()

	if options.IsAllowOrigin {
		router.Use(
			AllowOrigin(),
		)
		logrus.Infoln("开启跨域")
	}

	if options.IsRateLimite {
		router.Use(
			RateLimite(),
		)
		logrus.Infoln("开启限流")
	}

	if options.IsTrace {
		_, closer, err := trace.NewTracer(name, appconfig.MustGet().JaegerAddress)

		if err != nil {
			logrus.Fatal(err)
			return nil
		}
		defer closer.Close()

		router.Use(TracerWrapper)
		logrus.Infoln("开启链路追踪")
	}

	head := strings.TrimPrefix(name, "go.micro.api.")

	if options.IsOpenSwagger {
		var swaggerPath, swaggerURL string
		swaggerPath = fmt.Sprintf("/%s/swagger/*any", head)
		swaggerURL = fmt.Sprintf("http://%s:%s/%s/swagger/doc.json", config.HostIP, config.APIPort, head)

		UseSwagger(swaggerPath, swaggerURL, router)
	}

	if options.IsMonitor {
		router.Use(
			Prometheus(head),
		)
	}

	starter.Start(router)

	service.Handle("/", router)

	// run service
	if err := service.Run(); err != nil {
		return fmt.Errorf("服务运行错误:%w", err)
	}

	return nil
}

const contextTracerKey = "Tracer-context"

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

// TracerWrapper tracer 中间件
func TracerWrapper(c *gin.Context) {
	md := make(map[string]string)
	spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sp := opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
	defer sp.Finish()

	if err := opentracing.GlobalTracer().Inject(sp.Context(),
		opentracing.TextMap,
		opentracing.TextMapCarrier(md)); err != nil {
		logrus.Info(err)
	}

	ctx := context.TODO()
	ctx = opentracing.ContextWithSpan(ctx, sp)
	ctx = metadata.NewContext(ctx, md)
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

// ContextWithSpan 返回context
func ContextWithSpan(c *gin.Context) (ctx context.Context) {
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

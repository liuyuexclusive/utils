package webutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"

	"github.com/liuyuexclusive/utils/appconfigutil"
	"github.com/liuyuexclusive/utils/logutil"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
	"github.com/sirupsen/logrus"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

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

// func Validate() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")
// 		if token == "" {
// 			token = c.Query("token")
// 		}
// 		res, err := user.NewUserService("go.micro.srv.basic", client.DefaultClient).Validate(context.TODO(), &user.ValidateRequest{Token: token})
// 		if err != nil {
// 			c.JSON(401, err.Error())
// 			c.Abort()
// 		}
// 		c.Set("username", res.Name)
// 	}
// }

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
	//是否记录日志到ES 默认为false
	IsLogToES bool
	//是否允许跨域 默认为true
	IsAllowOrigin bool
	//是否限流 默认为true
	IsRateLimite bool
	//端口 默认为空
	Port string
}

type Option func(ops *Options)

func Startup(name string, starter Starter, opts ...Option) error {
	options := &Options{
		IsLogToES:     false,
		IsAllowOrigin: true,
		IsRateLimite:  true,
		Port:          "",
	}

	for _, opt := range opts {
		opt(options)
	}

	if options.IsLogToES {
		logutil.LogToES(name)
	}

	if name == "" {
		return errors.New("请输入服务名称")
	}

	config := appconfigutil.MustGet()
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
	}

	if options.IsRateLimite {
		router.Use(
			RateLimite(),
		)
	}

	var swaggerPath, swaggerURL string
	if options.Port != "" {
		swaggerPath = "/swagger/*any"
		swaggerURL = fmt.Sprintf("http://%s%s/swagger/doc.json", config.HostIP, options.Port)
	} else {
		head := strings.TrimPrefix(name, "go.micro.web.")
		swaggerPath = fmt.Sprintf("/%s/swagger/*any", head)
		swaggerURL = fmt.Sprintf("http://%s:%s/%s/swagger/doc.json", config.HostIP, config.APIPort, head)
	}

	UseSwagger(swaggerPath, swaggerURL, router)

	starter.Start(router)

	service.Handle("/", router)

	// run service
	if err := service.Run(); err != nil {
		return fmt.Errorf("服务运行错误:%w", err)
	}

	return nil
}

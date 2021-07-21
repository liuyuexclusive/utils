package redis

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/yuexclusive/utils/logger"

	"github.com/go-redis/redis"
)

const Nil = redis.Nil

var clientMapLock sync.Mutex
var clientMap = make(map[ClientName]*redis.Client)
var configMap = make(map[ClientName]*Config)
var sugar = logger.Single().Sugar()

// Client 根据name获取客户端连接
func Client(name ClientName) *redis.Client {
	return clientMap[name]
}

// InitClient 初始化连接
func InitClient(config *Config) (*redis.Client, error) {
	if v, ok := clientMap[config.ClientName]; ok && v != nil {
		if _, err := v.Ping().Result(); err == nil {
			return clientMap[config.ClientName], nil
		}
	}

	clientMapLock.Lock()
	defer clientMapLock.Unlock()

	var err error
	var client *redis.Client
	client, err = connect(config)
	if err != nil {
		return nil, err
	}

	if client != nil {
		clientMap[config.ClientName] = client
		configMap[config.ClientName] = config
	}

	return client, err
}

// init
func init() {
	go monitoring()
}

// monitoring 重连监控,无限循环，检查redis客服端是否断开连接，如果断开重新连接
func monitoring() {
	sugar.Info("redis 默认开启重连监控")
	defer func() {
		if err := recover(); err != nil {
			sugar.Errorf("redis connection aliveness sniffer stopped: %v", err)
			return
		}
		sugar.Error("redis connection aliveness sniffer stopped")
	}()

	for {
		// 先休眠30秒
		time.Sleep(30 * time.Second)
		c := ping()
		if len(c) <= 0 {
			continue
		}

		sugar.Infof("redis异常断开，正在尝试重连~~~~~")
		reconnect(c)
	}
}

// getExpirationConn 获取过期连接
func ping() []ClientName {
	r := make([]ClientName, 0, 1)
	clientMapLock.Lock()
	defer clientMapLock.Unlock()

	for _, v := range configMap {
		if c, ok := clientMap[v.ClientName]; ok {
			if _, err := c.Ping().Result(); err == nil {
				continue
			}
		}
		r = append(r, v.ClientName)
	}

	return r
}

// connect 建立redis连接
func connect(config *Config) (*redis.Client, error) {
	opt := &redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	}
	client := redis.NewClient(opt)

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	// 设置client name(app name)，方便定位问题
	if config.AppName != "" {
		if err := client.Process(redis.NewStringCmd("client", "setname", fmt.Sprintf("%s:%s", config.ClientName, config.AppName))); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// reconnect 重连
func reconnect(c []ClientName) {
	for _, v := range c {
		config := configMap[v]
		_, err := InitClient(config)
		if err != nil {
			b, _ := json.Marshal(config)
			sugar.Errorf("reconnect redis重连失败,config=%v,err=%+v", string(b), err)
		} else {
			sugar.Infof("reconnect redis重连成功...clientName=%s", config.ClientName)
		}
	}
}

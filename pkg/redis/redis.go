package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-hub/config"
	"go-hub/pkg/logger"
	"sync"
)

type Redis struct {
	rClient *redis.Client
	Context context.Context
}

var (
	once   sync.Once
	Client *Redis
)

func Connect(c config.RedisConfig) {
	once.Do(func() {
		host := fmt.Sprintf("%v:%v", c.Host, c.Port)
		Client = NewClient(host, c.Username, c.Password, c.Database)
	})
}

func NewClient(address string, username string, password string, db int) *Redis {
	rds := &Redis{}
	rds.Context = context.Background()

	// 使用 redis 库里的 NewClient 初始化连接
	rds.rClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	// 测试一下连接
	err := rds.Ping()
	logger.LogIf(err)

	return rds
}

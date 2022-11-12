package boot

import (
	"go-hub/config"
	"go-hub/pkg/redis"
)

func SetupRedis() {
	redis.Connect(config.Cfg.Redis)
}

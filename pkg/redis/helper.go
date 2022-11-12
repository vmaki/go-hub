package redis

import (
	"github.com/go-redis/redis/v8"
	"go-hub/pkg/logger"
	"time"
)

func (rds Redis) Ping() error {
	_, err := rds.rClient.Ping(rds.Context).Result()
	return err
}

func (rds Redis) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.rClient.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}

	return true
}

func (rds Redis) Get(key string) string {
	result, err := rds.rClient.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}

		return ""
	}

	return result
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds Redis) Has(key string) bool {
	_, err := rds.rClient.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}

		return false
	}

	return true
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds Redis) Del(keys ...string) bool {
	if err := rds.rClient.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}

	return true
}

func (rds Redis) Incr(key string) bool {
	if err := rds.rClient.Incr(rds.Context, key).Err(); err != nil {
		logger.ErrorString("Redis", "Incr", err.Error())
		return false
	}

	return true
}

func (rds Redis) IncrBy(key string, value int64) bool {
	if err := rds.rClient.IncrBy(rds.Context, key, value).Err(); err != nil {
		logger.ErrorString("Redis", "IncrBy", err.Error())
		return false
	}

	return true
}

func (rds Redis) Decr(key string) bool {
	if err := rds.rClient.Decr(rds.Context, key).Err(); err != nil {
		logger.ErrorString("Redis", "Decr", err.Error())
		return false
	}

	return true
}

func (rds Redis) DecrBy(key string, value int64) bool {
	if err := rds.rClient.DecrBy(rds.Context, key, value).Err(); err != nil {
		logger.ErrorString("Redis", "DecrBy", err.Error())
		return false
	}

	return true
}

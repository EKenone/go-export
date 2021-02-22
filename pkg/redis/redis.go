package redis

import (
	"github.com/go-redis/redis/v8"
	"go-export/internal/conf"
	"runtime"
)

func GetClient() *redis.Client {
	addr := conf.Conf.Redis.Host + ":" + conf.Conf.Redis.Port
	return redis.NewClient(&redis.Options{
		Addr:         addr,
		Username:     conf.Conf.Redis.User,
		Password:     conf.Conf.Redis.Pwd,
		MinIdleConns: runtime.NumCPU(),
	})
}

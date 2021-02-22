package redis

import (
	"github.com/go-redis/redis/v8"
	"go-export/internal/conf"
	"runtime"
	"sync"
)

var (
	cli  *redis.Client
	once = new(sync.Once)
)

func GetClient() *redis.Client {
	once.Do(func() {
		addr := conf.Conf.Redis.Host + ":" + conf.Conf.Redis.Port
		cli = redis.NewClient(&redis.Options{
			Addr:         addr,
			Username:     conf.Conf.Redis.User,
			Password:     conf.Conf.Redis.Pwd,
			MinIdleConns: runtime.NumCPU(),
		})
	})

	return cli

}

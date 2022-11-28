package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	service.ServiceConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	Redis      redis.RedisConf
}

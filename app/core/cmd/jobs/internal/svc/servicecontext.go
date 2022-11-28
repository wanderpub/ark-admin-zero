package svc

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/config"
	"ark-admin-zero/app/core/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config              //配置文件
	RedisClient       *redis.Redis               //redis
	TaskTemplateModel model.SysTaskTemplateModel //消息模版
	TaskAccountModel  model.SysTaskAccountModel  //发消息的帐号
	TaskJobModel      model.SysTaskJobModel      //计划任务
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})

	return &ServiceContext{
		Config:            c,
		RedisClient:       redisClient,
		TaskTemplateModel: model.NewSysTaskTemplateModel(conn, c.CacheRedis),
		TaskAccountModel:  model.NewSysTaskAccountModel(conn, c.CacheRedis),
		TaskJobModel:      model.NewSysTaskJobModel(conn, c.CacheRedis),
	}
}

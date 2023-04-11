package lib

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	Log         *zap.SugaredLogger
	EmailCfg    = &EmailCfgParam{}
	RedisCfg    = &RedisCfgParam{}
	RedisClient *redis.Client
	DBCfg       = &DBCfgParam{}
	DB          *sqlx.DB
)

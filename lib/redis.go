package lib

import "github.com/redis/go-redis/v9"

// Redis配置
type RedisCfgParam struct {
	Host     string `toml:"host"`     // 地址 127.0.0.1:6379
	Password string `toml:"password"` // 密码
	DB       int    `toml:"db"`       // 默认连接库
}

// 获取Redis Client
func (rc *RedisCfgParam) NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     rc.Host,
		Password: rc.Password,
		DB:       rc.DB,
	})
}

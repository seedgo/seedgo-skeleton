package common

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func ParseRedisConf() {
	if viper.IsSet("redis") {
		rdb = redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.host"),
			Password: viper.GetString("redis.password"), // 没有密码，默认值
			DB:       viper.GetInt("redis.db"),          // 默认DB 0
		})
	}
}

func RedisDB() *redis.Client {
	return rdb
}

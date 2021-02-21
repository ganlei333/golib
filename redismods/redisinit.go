package redismods

import (
	"ctock_info/config"
	"fmt"

	"github.com/letsfire/redigo/v2"
)

var RedisConn *redigo.Client

func RedisInit() (err error) {
	var Redisconfig = new(RedisConf)
	Redisconfig.RedisAddr = config.SerConf.RedisConf.RedisAddr

	Redisconfig.RedisDialDatabase = config.SerConf.RedisConf.RedisDialDatabase

	Redisconfig.RedisDialPassword = config.SerConf.RedisConf.RedisDialPassword

	Redisconfig.RedisIdleTimeout = config.SerConf.RedisConf.RedisIdleTimeout

	Redisconfig.RedisMaxActive = config.SerConf.RedisConf.RedisMaxActive

	Redisconfig.RedisMaxIdle = config.SerConf.RedisConf.RedisMaxIdle
	if config.SerConf.RedisConf.RedisMods == 0 {
		RedisConn, err = NewRedisAlone(Redisconfig)
		if err != nil {
			return
		}
	} else if config.SerConf.RedisConf.RedisMods == 1 {
		RedisConn, err = NewRedisSentinel(Redisconfig)
		if err != nil {
			return
		}
	} else if config.SerConf.RedisConf.RedisMods == 2 {
		RedisConn, err = NewRedisSentinel(Redisconfig)
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("Redis配置模式错误!")
	}
	err = Ping(RedisConn)
	return
}

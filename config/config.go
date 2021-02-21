package config

import (
	"github.com/Unknwon/goconfig"
)

var c = new(goconfig.ConfigFile)

func InitConfig() {
	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		panic(err)
	}
	c = cfg
	SerConf = new(ServerConf)
	SerConf.HTTPLogSwitch = cfg.MustBool(goconfig.DEFAULT_SECTION, "HttpLog", false)
	SerConf.HTTPPortValue = cfg.MustValue(goconfig.DEFAULT_SECTION, "HttpPort", "8080")
	SerConf.Token = cfg.MustValue(goconfig.DEFAULT_SECTION, "Token", "c03cd6a0de721e1857ead9d3933c4f30")
	SerConf.UserID = cfg.MustValue(goconfig.DEFAULT_SECTION, "UserID", "641130")
	SerConf.RedisConf.RedisAddr = cfg.MustValue(goconfig.DEFAULT_SECTION, "redis_addr", "127.0.0.1:6379")
	SerConf.RedisConf.RedisDialDatabase = cfg.MustInt(goconfig.DEFAULT_SECTION, "redis_dialdatabase", 0)
	SerConf.RedisConf.RedisDialPassword = cfg.MustValue(goconfig.DEFAULT_SECTION, "redis_dialpassword", "")
	SerConf.RedisConf.RedisIdleTimeout = cfg.MustInt(goconfig.DEFAULT_SECTION, "redis_idletimeout", 0)
	SerConf.RedisConf.RedisMaxActive = cfg.MustInt(goconfig.DEFAULT_SECTION, "redis_maxactive", 0)
	SerConf.RedisConf.RedisMaxIdle = cfg.MustInt(goconfig.DEFAULT_SECTION, "redis_maxidle", 0)
	SerConf.RedisConf.RedisMods = cfg.MustInt(goconfig.DEFAULT_SECTION, "redis_mods", 0)
	SerConf.LogConf.LogCompress = cfg.MustBool(goconfig.DEFAULT_SECTION, "LogCompress", false)
	SerConf.LogConf.LogFilename = cfg.MustValue(goconfig.DEFAULT_SECTION, "LogFilename", "logs/test.log")
	SerConf.LogConf.LogLevel = cfg.MustValue(goconfig.DEFAULT_SECTION, "LogLevel", "0")
	SerConf.LogConf.LogMaxAge = cfg.MustInt(goconfig.DEFAULT_SECTION, "LogMaxAge", 30)
	SerConf.LogConf.LogMaxBackups = cfg.MustInt(goconfig.DEFAULT_SECTION, "LogMaxBackups", 5)
	SerConf.LogConf.LogMaxSize = cfg.MustInt(goconfig.DEFAULT_SECTION, "LogMaxSize", 1)

}

func SetConfig(key, vue string) {
	c.SetValue(goconfig.DEFAULT_SECTION, key, vue)

}

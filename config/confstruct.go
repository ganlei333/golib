package config

var (
	SerConf *ServerConf
)

type ServerConf struct {
	HTTPLogSwitch bool
	HTTPPortValue string
	DebugLog      bool
	Token         string
	UserID        string
	LogConf       LogConf
	RedisConf     RedisConf
}

type LogConf struct {
	LogFilename   string
	LogMaxSize    int
	LogMaxBackups int
	LogMaxAge     int
	LogCompress   bool
	LogLevel      string
}

type RedisConf struct {
	RedisMods         int
	RedisAddr         string
	RedisMaxActive    int
	RedisMaxIdle      int
	RedisIdleTimeout  int
	RedisDialPassword string
	RedisDialDatabase int
}

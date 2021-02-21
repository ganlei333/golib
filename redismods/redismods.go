package redismods

import (
	"bytes"
	"encoding/json"

	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/letsfire/redigo/v2"

	"github.com/letsfire/redigo/v2/mode/alone"
	"github.com/letsfire/redigo/v2/mode/cluster"
	"github.com/letsfire/redigo/v2/mode/sentinel"
)

//RedisConf Redis连接池配置结构体
type RedisConf struct {
	RedisAddr         string
	RedisMaxActive    int
	RedisMaxIdle      int
	RedisIdleTimeout  int
	RedisDialPassword string
	RedisDialDatabase int
}

//NewRedisAlone 创建单实例redis连接池
func NewRedisAlone(rediss *RedisConf) (redispool *redigo.Client, err error) {

	var aloneMode = alone.New(
		alone.Addr(rediss.RedisAddr),
		alone.PoolOpts(
			redigo.MaxActive(rediss.RedisMaxActive), // 最大连接数，默认0无限制
			redigo.MaxIdle(rediss.RedisMaxIdle),     // 最多保持空闲连接数，默认2*runtime.GOMAXPROCS(0)
			redigo.Wait(false),                      // 连接耗尽时是否等待，默认false
			redigo.IdleTimeout(time.Duration(rediss.RedisIdleTimeout)*time.Second),
			// 空闲连接超时时间，默认0不超时
			redigo.MaxConnLifetime(0), // 连接的生命周期，默认0不失效
			redigo.TestOnBorrow(nil),  // 空间连接取出后检测是否健康，默认nil
		),
		alone.DialOpts(
			redis.DialReadTimeout(time.Second),           // 读取超时，默认time.Second
			redis.DialWriteTimeout(time.Second),          // 写入超时，默认time.Second
			redis.DialConnectTimeout(time.Second),        // 连接超时，默认500*time.Millisecond
			redis.DialPassword(rediss.RedisDialPassword), // 鉴权密码，默认空
			redis.DialDatabase(rediss.RedisDialDatabase), // 数据库号，默认0
			redis.DialKeepAlive(time.Minute*5),           // 默认5*time.Minute
			redis.DialNetDial(nil),                       // 自定义dial，默认nil
			redis.DialUseTLS(false),                      // 是否用TLS，默认false
			redis.DialTLSSkipVerify(false),               // 服务器证书校验，默认false
			redis.DialTLSConfig(nil),                     // 默认nil，详见tls.Config
		),
	)
	redispool = redigo.New(aloneMode)
	err = Ping(redispool)

	return

	// res, err := instance.String(func(c redis.Conn) (res interface{}, err error) {
	// 	return c.Do("ECHO", echoStr)
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// } else if res != echoStr {
	// 	log.Fatalf("unexpected result, expect = %s, but = %s", echoStr, res)
	// }
}

//NewRedisSentinel 创建哨兵redis连接
func NewRedisSentinel(rediss *RedisConf) (redispool *redigo.Client, err error) {

	server := strings.Split(rediss.RedisAddr, ";")

	var sentinelMode = sentinel.New(
		sentinel.Addrs(server),
		sentinel.PoolOpts(
			redigo.MaxActive(rediss.RedisMaxActive),                                // 最大连接数，默认0无限制
			redigo.MaxIdle(rediss.RedisMaxIdle),                                    // 最多保持空闲连接数，默认2*runtime.GOMAXPROCS(0)
			redigo.Wait(false),                                                     // 连接耗尽时是否等待，默认false
			redigo.IdleTimeout(time.Duration(rediss.RedisIdleTimeout)*time.Second), // 空闲连接超时时间，默认0不超时
			redigo.MaxConnLifetime(0),                                              // 连接的生命周期，默认0不失效
			redigo.TestOnBorrow(nil),                                               // 空间连接取出后检测是否健康，默认nil
		),
		sentinel.DialOpts(
			redis.DialReadTimeout(time.Second),           // 读取超时，默认time.Second
			redis.DialWriteTimeout(time.Second),          // 写入超时，默认time.Second
			redis.DialConnectTimeout(time.Second),        // 连接超时，默认500*time.Millisecond
			redis.DialPassword(rediss.RedisDialPassword), // 鉴权密码，默认空
			redis.DialDatabase(rediss.RedisDialDatabase), // 数据库号，默认0
			redis.DialKeepAlive(time.Minute*5),           // 默认5*time.Minute
			redis.DialNetDial(nil),                       // 自定义dial，默认nil
			redis.DialUseTLS(false),                      // 是否用TLS，默认false
			redis.DialTLSSkipVerify(false),               // 服务器证书校验，默认false
			redis.DialTLSConfig(nil),                     // 默认nil，详见tls.Config
		),
	)
	redispool = redigo.New(sentinelMode)
	err = Ping(redispool)
	return

	// res, err := instance.String(func(c redis.Conn) (res interface{}, err error) {
	// 	return c.Do("ECHO", echoStr)
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// } else if res != echoStr {
	// 	log.Fatalf("unexpected result, expect = %s, but = %s", echoStr, res)
	// }
}

//NewRedisCluster 创建redis集群连接
func NewRedisCluster(rediss *RedisConf) (redispool *redigo.Client, err error) {
	server := strings.Split(rediss.RedisAddr, ";")
	var clusterMode = cluster.New(
		cluster.Nodes(server),
		cluster.PoolOpts(
			redigo.MaxActive(rediss.RedisMaxActive),                                // 最大连接数，默认0无限制
			redigo.MaxIdle(rediss.RedisMaxIdle),                                    // 最多保持空闲连接数，默认2*runtime.GOMAXPROCS(0)
			redigo.Wait(false),                                                     // 连接耗尽时是否等待，默认false
			redigo.IdleTimeout(time.Duration(rediss.RedisIdleTimeout)*time.Second), // 空闲连接超时时间，默认0不超时
			redigo.MaxConnLifetime(0),                                              // 连接的生命周期，默认0不失效
			redigo.TestOnBorrow(nil),                                               // 空间连接取出后检测是否健康，默认nil
		),
		cluster.DialOpts(
			redis.DialReadTimeout(time.Second),           // 读取超时，默认time.Second
			redis.DialWriteTimeout(time.Second),          // 写入超时，默认time.Second
			redis.DialConnectTimeout(time.Second),        // 连接超时，默认500*time.Millisecond
			redis.DialPassword(rediss.RedisDialPassword), // 鉴权密码，默认空
			redis.DialDatabase(rediss.RedisDialDatabase), // 数据库号，默认0
			redis.DialKeepAlive(time.Minute*5),           // 默认5*time.Minute
			redis.DialNetDial(nil),                       // 自定义dial，默认nil
			redis.DialUseTLS(false),                      // 是否用TLS，默认false
			redis.DialTLSSkipVerify(false),               // 服务器证书校验，默认false
			redis.DialTLSConfig(nil),                     // 默认nil，详见tls.Config
		),
	)
	redispool = redigo.New(clusterMode)
	err = Ping(redispool)
	return

	// res, err := instance.String(func(c redis.Conn) (res interface{}, err error) {
	// 	return c.Do("ECHO", echoStr)
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// } else if res != echoStr {
	// 	log.Fatalf("unexpected result, expect = %s, but = %s", echoStr, res)
	// }
}

func GetRedisData(key string, num int, instance *redigo.Client) (data [][]byte, err error) {
	return instance.ByteSlices(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("LRANGE", key, num, num+1)
	})
}

func GetRedisNum(key string, instance *redigo.Client) (data int, err error) {
	return instance.Int(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("LLEN", key)
	})
}

func SedCh(key string, instance *redigo.Client, ch chan [][]string) {
	i := 0
	for {
		time.Sleep(time.Second * 1)
		num, err1 := GetRedisNum(key, instance)
		if err1 != nil {
			//fmt.Println(key, "num,err1:=GetRedisNum(key,instance)", err1)
			time.Sleep(time.Second * 1)
			continue
		}
		if i >= num {
			//fmt.Println(key,"redis 数据下标",num,"程序已经处理到下标",i)
			time.Sleep(time.Second * 1)
			continue
		}
		data, err := GetRedisData(key, i, instance)

		if err != nil || len(data) != 2 {
			//fmt.Println(key, "data,err:=GetRedisData(key,i,instance)", err)
			time.Sleep(time.Second * 1)
			continue
		}
		if !bytes.EqualFold(data[0], data[1]) {
			var c [][]string
			var b []string
			var a []string
			json.Unmarshal(data[0], &b)
			json.Unmarshal(data[1], &a)
			c = append(c, a, b)
			//fmt.Println(a)
			//fmt.Println(b)
			//fmt.Println(c)
			ch <- c

		}
		i++
	}
}

//GetRedisDataNum获取最近的一条Redis记录
func GetRedisDataNum(key string, instance *redigo.Client) (data []string, err error) {
	bydata, err := instance.Bytes(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("LINDEX", key, -1)
	})
	//var a []string
	err = json.Unmarshal(bydata, &data)
	return
}

//GetRedisDataNum获取最近的一条Redis记录
func Ping(instance *redigo.Client) (err error) {
	_, err = instance.Bytes(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("ping")
	})
	//var a []string
	if err != nil {
		return
	}
	return
}

//GetRedisDataNum获取最近的一条Redis记录
func Hget(instance *redigo.Client, res string) (data []string, err error) {
	data, err = instance.Strings(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("hget", res)
	})
	return
}

//GetRedisDataNum获取最近的一条Redis记录
func Blpop(instance *redigo.Client, res string) (data string, err error) {
	data, err = instance.String(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("hget", res, time.Second*3)
	})
	return
}

func Set(instance *redigo.Client, key string, dd []byte) (data string, err error) {
	return instance.String(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("set", key, dd)
	})
}

func Get(instance *redigo.Client, key string) (data []byte, err error) {
	return instance.Bytes(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("get", key)
	})
}

func Keys(instance *redigo.Client, key string) (data []string, err error) {
	return instance.Strings(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("Keys", key)
	})
}

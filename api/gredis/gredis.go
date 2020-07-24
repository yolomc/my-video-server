package gredis

import (
	"time"

	"github.com/go-playground/log"
	"github.com/gomodule/redigo/redis"
	"github.com/yolomc/my-video-server/api/config"
)

var RedisCli *redis.Pool

// Setup Initialize the Redis instance
func Setup() {
	RedisCli = &redis.Pool{
		MaxIdle:     config.RedisMaxIdle,
		MaxActive:   config.RedisMaxActive,
		IdleTimeout: config.RedisIdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisHost,
				redis.DialConnectTimeout(time.Duration(3000)*time.Millisecond))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return
}

func IsExists(key string) (bool, error) {
	redisConn := RedisCli.Get()
	defer redisConn.Close()
	is_key_exit, err := redis.Bool(redisConn.Do("EXISTS", key))
	if err != nil {
		log.Error("pkg.gredis.IsExists Do error is:", err)
		return false, err
	}
	return is_key_exit, nil
}

func Set(key, value string) error {
	redisConn := RedisCli.Get()
	defer redisConn.Close()

	_, err := redisConn.Do("SET", key, value, "EX", config.RedisTTL)
	if err != nil {
		log.Error("pkg.gredis.Set Do error is:", err)
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	redisConn := RedisCli.Get()
	defer redisConn.Close()

	value, err := redis.String(redisConn.Do("GET", key))
	if err != nil {
		log.Error("pkg.gredis.Get Do error is:", err)
		return "", err
	}
	return value, nil
}

func Del(key string) error {
	redisConn := RedisCli.Get()
	defer redisConn.Close()

	_, err := redisConn.Do("DEL", key)
	if err != nil {
		log.Error("pkg.gredis.Del Do error is:", err)
		return err
	}
	return nil
}

// Set redis hash表的使用
// func SetBit(key string, field, value interface{}) {
// 	redisConn := RedisCli.Get()
// 	defer redisConn.Close()
// 	_, err := redisConn.Do("SETBIT", key, field, value)
// 	if err != nil {
// 		log.Error("pkg.gredis.SETBIT Do error is:", err)
// 		return
// 	}
// 	return
// }

// GetBit redis hash表的使用
// func GetBit(key string, field interface{}) (interface{}, error) {
// 	redisConn := RedisCli.Get()
// 	defer redisConn.Close()
// 	reply, err := redisConn.Do("GetBit", key, field)
// 	if err != nil {
// 		log.Error("pkg.gredis.HGet Do error is:", err)
// 		return nil, err
// 	}
// 	return reply, nil
// }

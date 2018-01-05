package store

import (
	"fmt"
	"self/commons/g"
	"time"

	"github.com/garyburd/redigo/redis"
)

var RedisPool *redis.Pool

func initRedisPool() {
	//var err error
	cfg := g.Conf()

	config := cfg.Redis
	RedisPool = &redis.Pool{

		Dial: func() (redis.Conn, error) {

			conn, err := redis.Dial("tcp", config.Address, redis.DialPassword(config.Password), redis.DialDatabase(config.DB))
			if err != nil {
				fmt.Println("redispool connect error!")
				return nil, err
			}

			return conn, nil
		},
		MaxIdle:     config.PoolSize / 2,
		MaxActive:   config.PoolSize,
		IdleTimeout: 300 * time.Second,
	}

	return
}

func GetRedisConn() redis.Conn {

	return RedisPool.Get()
}

func closeRedis() {

}

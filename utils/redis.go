package utils

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/gomodule/redigo/redis"

	"fmt"
	"strconv"
)

var (
	redisClient *redis.Pool
)

func SetRedisClient(client *redis.Pool) {
	redisClient = client
}

func GetRedisClient() *redis.Pool {
	return redisClient
}

func GetRedisConn() redis.Conn {
	if redisClient == nil {
		logs.Error("GetRedisConn redisClient nil")
		return nil
	}

	conn := redisClient.Get()
	if conn.Err() != nil {
		logs.Error("GetRedisConn get nil", conn.Err())
		conn = redisClient.Get()
		if conn.Err() != nil {
			logs.Error("second GetRedisConn get nil", conn.Err())
		}
	}
	return conn
}

func LockByRedis(key string, ttl int64) bool {
	client := redisClient.Get()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)
	_, err := redis.String(client.Do("set", key, 1, "NX", "EX", ttl))
	if err == redis.ErrNil {
		// The lock was not successful, it already exists.
		return false
	}

	if err != nil {
		// redis fail
		return false
	}
	logs.Info("rebuild key passed ", key)
	return true
}

func ZCardByRedis(key string) int {
	client := redisClient.Get()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)
	num, _ := client.Do("ZCARD", key)
	n, _ := strconv.Atoi(fmt.Sprint(num))
	return n
}

func ZCountByRedis(key string, start, end int64) int {
	client := redisClient.Get()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)
	num, _ := client.Do("ZCOUNT", key, start, end)
	n, _ := strconv.Atoi(fmt.Sprint(num))
	return n
}

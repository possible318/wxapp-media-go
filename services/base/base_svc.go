package base

import (
	"encoding/json"
	"media/models/base"
	"media/utils"
	"time"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/gomodule/redigo/redis"
)

var baseService *Service

type Service struct {
}

func GetService() *Service {
	if baseService == nil {
		baseService = new(Service)
	}
	return baseService
}

func (f *Service) ExistInCache(key string, ctx *base.AppContext) bool {
	client := utils.GetRedisConn()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)

	if ctx.Force > 0 {
		return false
	}

	if ctx.IsReview {
		key += ":review"
	}

	exists, _ := redis.Bool(client.Do("EXISTS", key))
	return exists
}

func (f *Service) PutToCache(key string, data interface{}, expire int, ctx *base.AppContext) {
	client := utils.GetRedisConn()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)

	if ctx.IsReview {
		key += ":review"
	}

	if data != nil {
		bytes, err := json.Marshal(data)
		if err != nil {
			logs.Error("base.Service, PutToCache, error, ", key, data, err)
		} else {
			_, _ = client.Do("SET", key, bytes, "EX", expire)
		}
	}
}

func (f *Service) GetCache(key string, data interface{}) error {
	client := utils.GetRedisConn()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)
	bytes, err := redis.Bytes(client.Do("GET", key))
	if err == nil {
		re := json.Unmarshal(bytes, data)
		return re
	}
	return err
}

func (f *Service) SetCache(key string, data interface{}, expire int) {
	client := utils.GetRedisConn()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)

	if data != nil {
		bytes, err := json.Marshal(data)
		if err != nil {
			logs.Error("base.Service, setCache, error, ", key, data, err)
		} else {
			_, _ = client.Do("SET", key, bytes, "EX", expire)
		}
	}
}

func (f *Service) GetFromCache(key string, data interface{}, ctx *base.AppContext) error {
	client := utils.GetRedisConn()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)

	if ctx.IsReview {
		key += ":review"
	}

	bytes, err := redis.Bytes(client.Do("GET", key))

	if err != nil {
		// 失败进行一次重试
		bytes, err = redis.Bytes(client.Do("GET", key))
		if err != nil {
			logs.Error("second base.Service, GetFromCache, error, ", key, err, client.Err())
		} else {
			_ = json.Unmarshal(bytes, data)
		}
	} else {
		_ = json.Unmarshal(bytes, data)
	}
	return err
}

func (f *Service) SetRebuildTime(key string, rebuildAt int64) {
	key += ":rebuild"
	client := utils.GetRedisConn()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)
	_, _ = client.Do("SET", key, rebuildAt, "EX", rebuildAt-time.Now().Unix())
}

func (f *Service) IsNeedRebuildCache(key string) bool {
	client := utils.GetRedisConn()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)

	rebuildAt, err := redis.Int64(client.Do("GET", key))
	if err != nil || rebuildAt < time.Now().Unix() {
		return true
	}
	return false
}

func (f *Service) BatchGetFromCache(keys []string) map[string]*base.CacheData {
	client := utils.GetRedisConn()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)

	for _, key := range keys {
		_ = client.Send("GET", key)
	}

	res := make(map[string]*base.CacheData)
	err := client.Flush()
	if err != nil {
		logs.Info("base.Service, BatchGetFromCache, error, ", keys)
		return res
	}

	for _, key := range keys {
		bytes, err := redis.Bytes(client.Receive())
		if err != nil {
			continue
		}

		data := new(base.CacheData)
		_ = json.Unmarshal(bytes, data)
		res[key] = data
	}
	return res
}

func (f *Service) MGetFromCache(keys []string, data []base.CacheData) {
	client := utils.GetRedisConn()
	defer func(client redis.Conn) {
		_ = client.Close()
	}(client)

	args := make([]interface{}, 0, len(keys))
	for _, k := range keys {
		args = append(args, k)
	}
	byteSlices, err := redis.ByteSlices(client.Do("MGET", args...))
	if err != nil {
		logs.Info("base.Service, MGetFromCache, error, ", err)
		return
	}

	for i, v := range byteSlices {
		if len(v) > 0 {
			_ = json.Unmarshal(v, &data[i])
		} else {
			logs.Debug("empty key:", keys[i])
		}
	}
}

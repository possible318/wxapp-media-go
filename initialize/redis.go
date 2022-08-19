package initialize

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
	_ "media/routers"
	"media/utils"
	"time"
)

func RegisterRedis() {
	logs.Info("Lib:Redis Init, Start!")
	client := initRedis("redis")
	utils.SetRedisClient(client)
	logs.Info("Lib:Redis Init, Finish!")
}

func initRedis(name string) *redis.Pool {
	logs.Info("Lib:Redis Init, Start!", name)
	Host, _ := web.AppConfig.String(name + ".host")
	Port, _ := web.AppConfig.String(name + ".port")
	DB, _ := web.AppConfig.String(name + ".db")
	maxIdle := web.AppConfig.DefaultInt(name+".max_idle", 10)
	maxActive := web.AppConfig.DefaultInt(name+".max_active", 10)

	client := &redis.Pool{
		MaxIdle:   maxIdle,
		MaxActive: maxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", Host+":"+Port,
				redis.DialConnectTimeout(time.Millisecond*500),
				redis.DialReadTimeout(time.Millisecond*500),
				redis.DialWriteTimeout(time.Millisecond*500),
			)
			if err != nil {
				return nil, err
			}
			if DB != "" {
				_, err := c.Do("SELECT", DB)
				if err != nil {
					return nil, err
				}
			}
			return c, nil
		},
	}
	logs.Info("Lib:Redis Init, Finish!", name)
	return client
}

package main

import (
	"media/initialize"
	_ "media/routers"
	"net/http"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

// 初始化
func init() {
	logs.Info("Library Init Start")
	// 注册mysql
	initialize.RegisterDatabase()
	// 注册redis
	initialize.RegisterRedis()
	// 注册日志服务
	initialize.RegisterLog()
	logs.Info("Library Init End")
}

func main() {
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.EnableErrorsRender = true
		web.BConfig.EnableErrorsShow = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		go func() {
			_ = http.ListenAndServe(":8030", nil)
		}()
	}

	web.BConfig.RecoverPanic = true
	web.BConfig.WebConfig.AutoRender = false
	web.Run()
}

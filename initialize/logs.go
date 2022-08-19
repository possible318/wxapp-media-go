package initialize

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/server/web"
)

func RegisterLog() {
	path, _ := web.AppConfig.String("log.path")
	cfg := `{"filename":"` + path + `","separate":["error","info","warning","debug"]}`

	logs.Info("Lib:Log Init," + path + " Start!")

	err := logs.SetLogger(logs.AdapterMultiFile, cfg)
	if err != nil {
		logs.Error("Lib:Log Init, " + path + " Failed!")
	}
	logs.EnableFuncCallDepth(true)
	logs.Info("Lib:Log Init, " + path + " Finish!")
	if web.BConfig.RunMode != "dev" {
		_ = logs.SetLogger(logs.AdapterConsole)
	}
	web.BConfig.Log.AccessLogs = true
}

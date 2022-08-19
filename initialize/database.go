package initialize

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

func RegisterDatabase() {
	logs.Info("Lib:Database Init, Start!")
	initMysql()
	logs.Info("Lib:Database Init, Finish!")
}

func initMysql() bool {
	logs.Info("Lib:Database Mysql Init, Start")
	dbs := GetDBList()
	for _, db := range dbs {
		if !initMysqlDB(db) {
			return false
		}
	}
	logs.Info("Lib:Database Mysql Init, Finish")
	runMode, _ := web.AppConfig.String("RunMode")
	if runMode == "dev" || runMode == "test" {
		orm.Debug = true
	}
	return true
}

func initMysqlDB(name string) bool {
	logs.Info("Lib:Database Mysql Init, " + name + " Start")

	dbDriver, _ := web.AppConfig.String("db." + name + ".driver")
	dbUser, _ := web.AppConfig.String("db." + name + ".user")
	dbPwd, _ := web.AppConfig.String("db." + name + ".pwd")
	dbHost, _ := web.AppConfig.String("db." + name + ".host")
	dbPort, _ := web.AppConfig.String("db." + name + ".port")
	dbName, _ := web.AppConfig.String("db." + name + ".name")

	link := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	err := orm.RegisterDataBase(name, dbDriver, link)
	if err != nil {
		logs.Error("Lib:Database MySQL Init," + name + " Failed!")
		return false
	}
	logs.Info("Lib:Database MySQL Init," + name + " Finish!")
	return true
}

func GetDBList() []string {
	return []string{"default"}
}

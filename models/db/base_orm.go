package db

import "github.com/beego/beego/v2/client/orm"

func GetConnection() orm.Ormer {
	conn := orm.NewOrm()
	return conn
}

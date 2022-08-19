package db

import (
	"github.com/beego/beego/v2/client/orm"
)

func init() {
	orm.RegisterModel(new(BlPhoto))
}

type BlPhoto struct {
	ID       int    `orm:"column(id)" json:"id"`
	ItemID   string `orm:"column(item_id)" json:"item_id"`
	Platform string `json:"platform"`
	Text     string `json:"text"`
	Pid      string `json:"pid"`
	Src      string `json:"src"`
	AddTime  string `json:"add_time"`
}

func (f BlPhoto) TableName() string {
	return "bl_photo"
}

func (f BlPhoto) GetQuery() orm.QuerySeter {
	conn := GetConnection()
	return conn.QueryTable(f.TableName())
}

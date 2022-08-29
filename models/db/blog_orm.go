package db

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
)

func init() {
	orm.RegisterModel(new(Blog))
}

type Blog struct {
	ID       int    `orm:"column(id)" json:"id"`
	Platform string `json:"platform"`
	ItemID   string `orm:"column(item_id)" json:"item_id"`
	Text     string `json:"text"`
	Pid      string `json:"pid"`
	Src      string `json:"src"`
	AddTime  string `json:"add_time"`
	ShowType int    `json:"show_type"`
	Deleted  int    `json:"deleted"`
	Index    int    `json:"index"`
	Height   string `json:"height"`
	Width    string `json:"width"`
	Star     int    `json:"star"`
	Step     int    `json:"step"`
	Status   int    `json:"status"`
}

func (f Blog) TableName() string {
	return "blog"
}

func (f Blog) GetQuery() orm.QuerySeter {
	conn := GetConnection()
	return conn.QueryTable(f.TableName())
}

func (f Blog) GetById(id int) (*Blog, error) {
	r := new(Blog)
	err := GetConnection().QueryTable(f.TableName()).Filter("Id", id).One(r)
	if err != nil {
		logs.Error(err)
	}
	return r, err
}

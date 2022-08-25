package db

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(Same))
}

type Same struct {
	Id int
	A  string
	B  string
}

func (f Same) TableName() string {
	return "Same"
}

func (f Same) GetQuery() orm.QuerySeter {
	conn := GetConnection()
	return conn.QueryTable(f.TableName())
}

func (f Same) GetSameMsg() []orm.Params {
	sql := `
select
	b.id as b_id,
    b.platform as b_pf,
    b.src as b_src,
    c.id as c_id,
    c.platform as c_pf,
    c.src as c_src
from same as a
    inner join blog as b on a.a = b.pid
    inner join blog as c on a.b = c.pid`
	var itemList []orm.Params
	conn := GetConnection()
	_, _ = conn.Raw(sql).Values(&itemList)
	return itemList
}

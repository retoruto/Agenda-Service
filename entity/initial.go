package entity

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"

)

var engine *xorm.Engine

func init() {
	orm, err1 := xorm.NewEngine("sqlite3", "./agenda.db")
	checkErr(err1)

	//自动检测和创建表,自动转换varchar字段类型到text字段类型
	err1 = orm.Sync2(new(UserTable), new(MeetingTable))
	checkErr(err1)
	engine = orm
	//engine.ShowSQL(true)
	//名称映射规则,主要负责结构体名称到表名和结构体field到表字段的名称映射
	//engine.SetMapper(core.GonicMapper{})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

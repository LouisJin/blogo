package service

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var sql orm.Ormer

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	if err := orm.RegisterDataBase("default", "sqlite3", "db/data.db"); err != nil {
		logs.Error("数据库初始化错误！", err)
		return
	}
	if location, err := time.LoadLocation("Asia/Shanghai"); err == nil {
		orm.SetDataBaseTZ("default", location)
	}
	sql = orm.NewOrm()
}

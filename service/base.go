package service

import (
	"blogo/g"
	"context"
	"fmt"
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var sql orm.Ormer
var mCache cache.Cache

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

	var err error
	mCache, err = cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		logs.Error("内存缓存初始化错误！", err)
		return
	}
}

const (
	CacheComment = iota
	CacheThumbsup
)

/**
设置缓存，防止ip多次操作
bool true 设置缓存成功  false 表示失败
error 表示缓存出问题了
*/
func SetCache(cacheType int, ip string) (bool, error) {
	var expireTime time.Duration = 1
	switch cacheType {
	case CacheComment:
		expireTime = time.Duration(g.GlobalConfig.Policy.CommentInterval)
	case CacheThumbsup:
		expireTime = time.Duration(g.GlobalConfig.Policy.ThumbsupInterval)
	}
	key := fmt.Sprintf("%d-%s", cacheType, ip)
	exist, err := mCache.IsExist(context.TODO(), key)
	if err != nil {
		return false, err
	}
	if exist {
		return false, nil
	}
	err2 := mCache.Put(context.TODO(), key, 1, expireTime*time.Minute)
	if err2 != nil {
		return false, err2
	}
	return true, nil
}

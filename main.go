package main

import (
	"blogo/g"
	_ "blogo/routers"
	context2 "context"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	logs.EnableFuncCallDepth(true)
}

var filter beego.FilterFunc = func(ctx *context.Context) {
	session := ctx.Input.CruSession.Get(context2.Background(), g.UserToken)
	if session == nil {
		logs.Warn("未登陆，跳转到登陆界面")
		ctx.Redirect(200, "/admin")
	}
}

func main() {
	// session设置
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 86400
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true
	// 拦截器
	beego.InsertFilter("/admin/api/*", beego.BeforeRouter, filter)
	beego.Run()
}

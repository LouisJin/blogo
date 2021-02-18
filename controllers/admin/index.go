package admin

import (
	"blogo/g"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/satori/go.uuid"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	session := c.GetSession(g.UserToken)
	if session == nil {
		c.TplName = "login.tpl"
	} else {
		c.TplName = "admin.tpl"
	}
}

func (c *MainController) Login() {
	webResult := g.GetDefaultWebResult()
	username := c.GetString("username")
	if username == "" {
		webResult.Code = g.Failure
		webResult.Msg = "用户名不能为空"
		c.Data["json"] = webResult
		c.ServeJSON(true)
		return
	}
	password := c.GetString("password")
	if password == "" {
		webResult.Code = g.Failure
		webResult.Msg = "密码不能为空"
		c.Data["json"] = webResult
		c.ServeJSON(true)
		return
	}
	encryPassword := g.EncryptPassword(password)
	if username == g.GlobalConfig.Admin.Username && encryPassword == g.GlobalConfig.Admin.Password {
		c.SetSession(g.UserToken, uuid.NewV4())
		webResult.Code = g.Ok
		webResult.Msg = "登陆成功"
		c.Data["json"] = webResult
		c.ServeJSON(true)
		return
	} else {
		webResult.Code = g.Failure
		webResult.Msg = "用户名或密码不正确"
		c.Data["json"] = webResult
		c.ServeJSON(true)
	}
}

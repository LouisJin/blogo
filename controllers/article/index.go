package article

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	id, _ := c.GetInt(":id")
	article, b := articleService.Query(id)
	if b {
		c.Data["article"] = article
		c.TplName = "article.tpl"
	} else {
		c.Abort("404")
	}
}

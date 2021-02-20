package controllers

import (
	"blogo/g"
	"blogo/models"
	"blogo/service"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

var articleService = new(service.ArticleService)
var articleGroupService = new(service.ArticleGroupService)

func (c *MainController) Get() {
	articleGroups, b := articleGroupService.List(&models.PageDto{Page: 0, Size: 99})
	if b {
		c.Data["articleGroups"] = articleGroups
	}
	c.Data["globalConfig"] = g.GlobalConfig
	c.TplName = "index.tpl"
}

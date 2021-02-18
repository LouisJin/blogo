package controllers

import (
	"blogo/g"
	"blogo/models"
	"blogo/service"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

var articleService = new(service.ArticleService)
var articleGroupService = new(service.ArticleGroupService)

func (c *MainController) Get() {
	articleDto := new(models.ArticleDto)
	err := c.ParseForm(articleDto)
	if err != nil {
		logs.Error("获取 articleDto 参数失败", err)
	}

	articleGroups, b := articleGroupService.List(&models.PageDto{Page: 0, Size: 99})
	if b {
		c.Data["articleGroups"] = articleGroups
	}
	c.Data["globalConfig"] = g.GlobalConfig

	articles, b := articleService.List(articleDto)
	if b {
		c.Data["articles"] = articles
	}
	c.Data["groupId"] = articleDto.GroupId
	c.Data["title"] = articleDto.Title
	c.TplName = "index.tpl"
}

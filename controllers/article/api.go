package article

import (
	"blogo/g"
	"blogo/models"
	"blogo/service"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type ApiController struct {
	beego.Controller
}

var articleService = new(service.ArticleService)

func (c *ApiController) ArticleList() {
	webResult := g.GetDefaultWebResult()
	articleDto := new(models.ArticleDto)
	err := c.ParseForm(articleDto)
	if err != nil {
		logs.Error("获取 articleDto 参数失败", err)
	}
	if articleDto.Page != 0 {
		articleDto.Page -= 1
	}
	list, count, b := articleService.List(articleDto)
	if b {
		webResult.Code = g.Ok
		data := make(map[string]interface{})
		data["list"] = list
		data["count"] = count
		webResult.Data = data
	} else {
		webResult.Code = g.Failure
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

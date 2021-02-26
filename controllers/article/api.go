package article

import (
	"blogo/g"
	"blogo/models"
	"blogo/service"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
)

type ApiController struct {
	beego.Controller
}

var articleService = new(service.ArticleService)
var articleCommentService = new(service.ArticleCommentService)

func (c *ApiController) ArticleList() {
	webResult := g.GetDefaultWebResult()
	articleDto := new(models.ArticleDto)
	c.ParseForm(articleDto)
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

func (c *ApiController) ArticleCommentAdd() {
	webResult := g.GetDefaultWebResult()
	articleCommentDto := new(models.ArticleCommentDto)
	c.ParseForm(articleCommentDto)
	valid := validation.Validation{}
	ok, _ := valid.Valid(articleCommentDto)
	if !ok {
		webResult.Msg = "获取 articleCommentDto 参数失败"
		c.Data["json"] = webResult
		c.ServeJSON(true)
		return
	}
	session := c.GetSession(g.UserToken)
	if session != nil {
		articleCommentDto.IsAdmin = 1
	}
	ok, err := service.SetCache(service.CacheComment, c.Ctx.Input.IP())
	if err != nil {
		webResult.Msg = fmt.Sprintf("设置评论缓存失败 ip: %s", c.Ctx.Input.IP())
		c.Data["json"] = webResult
		c.ServeJSON(true)
		return
	}
	if ok {
		_, b := articleCommentService.Insert(articleCommentDto)
		if b {
			webResult.Code = g.Ok
			webResult.Msg = "添加文章评论成功"
			articleService.UpdateCommentNum(articleCommentDto.ArticleId)
		} else {
			webResult.Code = g.Failure
			webResult.Msg = "添加文章评论失败"
		}
	} else {
		webResult.Msg = "您评论频率过高,请稍后评论!"
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

func (c *ApiController) ArticleThumbsup() {
	webResult := g.GetDefaultWebResult()
	id, err := c.GetInt("id")
	if err != nil {
		webResult.Msg = "获取 文章id 失败"
		c.Data["json"] = webResult
		c.ServeJSON(true)
		return
	}
	ok, err := service.SetCache(service.CacheThumbsup, c.Ctx.Input.IP())
	if err != nil {
		webResult.Msg = fmt.Sprintf("设置点赞缓存失败 ip: %s", c.Ctx.Input.IP())
		c.Data["json"] = webResult
		c.ServeJSON(true)
		return
	}
	if ok {
		_, b := articleService.IncreaseThumbsup(id)
		if b {
			webResult.Code = g.Ok
			webResult.Msg = "文章点赞成功"
		} else {
			webResult.Code = g.Failure
			webResult.Msg = "文章点赞失败"
		}
	} else {
		webResult.Msg = "您点赞频率过高,请稍后点赞!"
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

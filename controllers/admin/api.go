package admin

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
var articleGroupService = new(service.ArticleGroupService)
var articleCommentService = new(service.ArticleCommentService)

func (c *ApiController) ArticleAdd() {
	webResult := g.GetDefaultWebResult()
	articleDto := new(models.ArticleDto)
	err := c.ParseForm(articleDto)
	if err != nil {
		logs.Error("获取 articleDto 参数失败", err)
	}
	_, b := articleService.Insert(articleDto)
	if b {
		webResult.Code = g.Ok
		webResult.Msg = "添加文章成功"
	} else {
		webResult.Code = g.Failure
		webResult.Msg = "添加文章失败"
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

func (c *ApiController) ArticleEdit() {
	webResult := g.GetDefaultWebResult()
	articleDto := new(models.ArticleDto)
	err := c.ParseForm(articleDto)
	if err != nil {
		logs.Error("获取 articleDto 参数失败", err)
	}
	_, b := articleService.Update(articleDto)
	if b {
		webResult.Code = g.Ok
		webResult.Msg = "编辑文章成功"
	} else {
		webResult.Code = g.Failure
		webResult.Msg = "编辑文章失败"
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

func (c *ApiController) ArticleDel() {
	webResult := g.GetDefaultWebResult()
	id, _ := c.GetInt(":id")
	_, b := articleService.Delete(id)
	if b {
		webResult.Code = g.Ok
		webResult.Msg = "删除文章成功"
	} else {
		webResult.Code = g.Failure
		webResult.Msg = "删除文章失败"
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

func (c *ApiController) ArticleGroupAdd() {
	webResult := g.GetDefaultWebResult()
	articleGroupDto := new(models.ArticleGroupDto)
	err := c.ParseForm(articleGroupDto)
	if err != nil {
		logs.Error("获取 articleGroupDto 参数失败", err)
	}
	_, b := articleGroupService.Insert(articleGroupDto)
	if b {
		webResult.Code = g.Ok
		webResult.Msg = "添加文章分类成功"
	} else {
		webResult.Code = g.Failure
		webResult.Msg = "添加文章分类失败"
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

func (c *ApiController) ArticleGroupEdit() {
	webResult := g.GetDefaultWebResult()
	articleGroupDto := new(models.ArticleGroupDto)
	err := c.ParseForm(articleGroupDto)
	if err != nil {
		logs.Error("获取 articleGroupDto 参数失败", err)
	}
	_, b := articleGroupService.Update(articleGroupDto)
	if b {
		webResult.Code = g.Ok
		webResult.Msg = "编辑文章分类成功"
	} else {
		webResult.Code = g.Failure
		webResult.Msg = "编辑文章分类失败"
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

func (c *ApiController) ArticleGroupDel() {
	webResult := g.GetDefaultWebResult()
	id, _ := c.GetInt(":id")
	_, b := articleGroupService.Delete(id)
	if b {
		webResult.Code = g.Ok
		webResult.Msg = "删除文章分类成功"
	} else {
		webResult.Code = g.Failure
		webResult.Msg = "删除文章分类失败"
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

func (c *ApiController) ArticleCommentAdd() {
	webResult := g.GetDefaultWebResult()
	articleCommentDto := new(models.ArticleCommentDto)
	err := c.ParseForm(articleCommentDto)
	if err != nil {
		logs.Error("获取 articleCommentDto 参数失败", err)
	}
	_, b := articleCommentService.Insert(articleCommentDto)
	if b {
		webResult.Code = g.Ok
		webResult.Msg = "添加文章评论成功"
	} else {
		webResult.Code = g.Failure
		webResult.Msg = "添加文章评论失败"
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

func (c *ApiController) ArticleCommentDel() {
	webResult := g.GetDefaultWebResult()
	id, _ := c.GetInt(":id")
	_, b := articleCommentService.Delete(id)
	if b {
		webResult.Code = g.Ok
		webResult.Msg = "删除文章评论成功"
	} else {
		webResult.Code = g.Failure
		webResult.Msg = "删除文章评论失败"
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

func (c *ApiController) Logout() {
	err := c.DelSession(g.UserToken)
	if err == nil {
		c.Redirect("/admin", 200)
	}
}

func (c *ApiController) ConfigEdit() {
	webResult := g.GetDefaultWebResult()
	globalConfig := new(g.Config)
	err := c.ParseForm(globalConfig)
	if err != nil {
		logs.Error("获取 globalConfig 参数失败", err)
	}
	if globalConfig.Admin.Password != "" {
		globalConfig.Admin.Password = g.EncryptPassword(globalConfig.Admin.Password)
	}
	ok := g.SetGlobalConfig(globalConfig)
	if ok {
		webResult.Code = g.Ok
		webResult.Msg = "编辑配置成功"
	} else {
		webResult.Code = g.Failure
		webResult.Msg = "编辑配置失败"
	}
	c.Data["json"] = webResult
	c.ServeJSON(true)
}

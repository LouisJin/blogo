package routers

import (
	"blogo/controllers"
	"blogo/controllers/admin"
	"blogo/controllers/article"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/admin", &admin.MainController{})
	beego.Router("/admin/login", &admin.MainController{}, "post:Login")
	beego.Router("/admin/api/article", &admin.ApiController{}, "post:ArticleAdd")
	beego.Router("/admin/api/article", &admin.ApiController{}, "put:ArticleEdit")
	beego.Router("/admin/api/article/:id([0-9]+)", &admin.ApiController{}, "delete:ArticleDel")
	beego.Router("/admin/api/articleGroup", &admin.ApiController{}, "post:ArticleGroupAdd")
	beego.Router("/admin/api/articleGroup", &admin.ApiController{}, "put:ArticleGroupEdit")
	beego.Router("/admin/api/articleGroup/:id([0-9]+)", &admin.ApiController{}, "delete:ArticleGroupDel")
	beego.Router("/admin/api/articleComment", &admin.ApiController{}, "post:ArticleCommentAdd")
	beego.Router("/admin/api/articleComment/:id([0-9]+)", &admin.ApiController{}, "delete:ArticleCommentDel")
	beego.Router("/admin/api/logout", &admin.ApiController{}, "*:Logout")
	beego.Router("/admin/api/configEdit", &admin.ApiController{}, "put:ConfigEdit")

	beego.Router("/article/:id", &article.MainController{})
	beego.Router("/article/api/list", &article.ApiController{}, "get:ArticleList")
}

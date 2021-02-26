package models

type ArticleCommentDto struct {
	ArticleId int    `form:"articleId" valid:"Required"`
	Content   string `form:"content" valid:"Required"`
	ParentId  int    `form:"parentId"`
	IsAdmin   int    `form:"isAdmin"`
	PageDto
}

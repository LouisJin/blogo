package models

type ArticleCommentDto struct {
	ArticleId int    `form:"articleId"`
	Content   string `form:"content"`
	ParentId  int    `form:"parentId"`
	IsAdmin   int    `form:"isAdmin"`
	PageDto
}

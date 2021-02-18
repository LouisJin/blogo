package models

type ArticleDto struct {
	Id        int    `form:"id"`
	GroupId   int    `form:"groupId"`
	Title     string `form:"title"`
	Content   string `form:"content"`
	IsComment int    `form:"isComment"`
	PageDto
}

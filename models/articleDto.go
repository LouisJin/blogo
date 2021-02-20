package models

import "time"

type ArticleDto struct {
	Id          int `form:"id"`
	GroupId     int `form:"groupId"`
	GroupName   string
	Title       string `form:"title"`
	Content     string `form:"content"`
	Q           string `form:"q"` // 搜索字段
	IsComment   int    `form:"isComment"`
	ThumbsupNum int
	CommentNum  int
	CreateTime  time.Time
	UpdateTime  time.Time
	PageDto
}

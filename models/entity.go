package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

/**
* 文章
 */
type Article struct {
	Id          int
	Group       *ArticleGroup `orm:"rel(one)"`
	Title       string
	Content     string `json:"Content,omitempty"`
	ThumbsupNum int
	CommentNum  int
	CreateTime  time.Time
	UpdateTime  time.Time
	IsComment   int
	IsDelete    int
}

/**
* 文章分组
 */
type ArticleGroup struct {
	Id         int
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
}

/**
* 文章评论
 */
type ArticleComment struct {
	Id         int
	ArticleId  int
	Content    string
	ParentId   int
	IsAdmin    int
	CreateTime time.Time
	IsDelete   int
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(Article), new(ArticleGroup), new(ArticleComment))
}

package service

import (
	"blogo/models"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

type ArticleCommentService struct {
}

func (service *ArticleCommentService) Insert(articleCommentDto *models.ArticleCommentDto) (int64, bool) {
	one := models.ArticleComment{
		ArticleId:  articleCommentDto.ArticleId,
		Content:    articleCommentDto.Content,
		ParentId:   articleCommentDto.ParentId,
		IsAdmin:    articleCommentDto.IsAdmin,
		CreateTime: time.Now(),
	}
	insert, err := sql.Insert(&one)
	if err != nil {
		logs.Error("数据库插入 ArticleComment 失败！", err)
		return -1, false
	}
	return insert, true
}

func (service *ArticleCommentService) List(pageDto *models.PageDto) ([]models.ArticleComment, bool) {
	qs := sql.QueryTable(models.ArticleComment{})
	qs = qs.Limit(models.GetLimit(pageDto.Page, pageDto.Size))
	qs = qs.Filter("isDelete", false)

	var sqlData []models.ArticleComment
	_, err := qs.All(&sqlData)
	if err != nil {
		logs.Error("数据库查询 ArticleComment 列表失败！", err)
		return sqlData, false
	}
	return sqlData, true
}

func (service *ArticleCommentService) Delete(id int) (int64, bool) {
	one := models.ArticleComment{Id: id}
	err := sql.Read(&one)
	if err != nil {
		logs.Error("数据库删除 ArticleComment 失败！找不到id: %d", id)
		return -1, false
	}
	one.IsDelete = 1
	delete, err := sql.Update(&one)
	if err != nil {
		logs.Error("数据库删除 ArticleComment 失败！%d, %s", id, err)
		return -1, false
	}
	return delete, true
}

package service

import (
	"blogo/models"
	"github.com/beego/beego/v2/core/logs"
	"strings"
	"time"
)

type ArticleService struct {
}

func (service *ArticleService) Insert(articleDto *models.ArticleDto) (int64, bool) {
	one := models.Article{
		GroupId:    articleDto.GroupId,
		Title:      articleDto.Title,
		Content:    articleDto.Content,
		CreateTime: time.Now(),
	}
	insert, err := sql.Insert(&one)
	if err != nil {
		logs.Error("数据库插入 Article 失败！", err)
		return -1, false
	}
	return insert, true
}

func (service *ArticleService) Update(articleDto *models.ArticleDto) (int64, bool) {
	one := models.Article{Id: articleDto.Id}
	err := sql.Read(&one)
	if err != nil {
		logs.Error("数据库更新 Article 失败！找不到id: %d", articleDto.Id)
		return -1, false
	}
	if articleDto.Title != "" {
		one.Title = articleDto.Title
	}
	if articleDto.Content != "" {
		one.Content = articleDto.Content
	}
	if articleDto.GroupId != 0 {
		one.GroupId = articleDto.GroupId
	}
	if articleDto.IsComment == 1 {
		one.IsComment = 1
	} else if articleDto.IsComment == -1 {
		one.IsComment = 0
	}
	one.UpdateTime = time.Now()
	update, err := sql.Update(&one)
	if err != nil {
		logs.Error("数据库更新 Article 失败！", err)
		return -1, false
	}
	return update, true
}

func (service *ArticleService) Query(id int) (models.Article, bool) {
	one := models.Article{Id: id}
	err := sql.Read(&one)
	if err != nil {
		logs.Error("数据库查询 Article 失败！id:%d, %s", id, err)
		return one, false
	}
	return one, true
}

func (service *ArticleService) List(articleDto *models.ArticleDto) ([]models.Article, bool) {
	qs := sql.QueryTable(models.Article{})
	qs = qs.Limit(models.GetLimit(articleDto.Page, articleDto.Size))
	qs = qs.Filter("isDelete", false)

	if strings.Trim(articleDto.Title, "") != "" {
		qs = qs.Filter("title__icontains", articleDto.Title)
	}
	if articleDto.GroupId != 0 {
		qs = qs.Filter("groupId", articleDto.GroupId)
	}
	var sqlData []models.Article
	_, err := qs.All(&sqlData, "Id", "GroupId", "Title", "ThumbsupNum", "CommentNum", "CreateTime", "UpdateTime", "IsComment")
	if err != nil {
		logs.Error("数据库查询 Article 列表失败！", err)
		return sqlData, false
	}
	return sqlData, true
}

func (service *ArticleService) Delete(id int) (int64, bool) {
	one := models.Article{Id: id}
	err := sql.Read(&one)
	if err != nil {
		logs.Error("数据库删除 Article 失败！找不到id: %d", id)
		return -1, false
	}
	one.UpdateTime = time.Now()
	one.IsDelete = 1
	delete, err := sql.Update(&one)
	if err != nil {
		logs.Error("数据库删除 Article 失败！%d, %s", id, err)
		return -1, false
	}
	return delete, true
}

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
		Group: &models.ArticleGroup{
			Id: articleDto.GroupId,
		},
		Title:      articleDto.Title,
		Content:    articleDto.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
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
		one.Group.Id = articleDto.GroupId
	}
	if articleDto.IsComment == 1 {
		one.IsComment = 1
	} else if articleDto.IsComment == -1 {
		one.IsComment = 0
	}
	if articleDto.CommentNum != 0 {
		one.CommentNum = articleDto.CommentNum
	}
	if articleDto.ThumbsupNum != 0 {
		one.ThumbsupNum = articleDto.ThumbsupNum
	}
	one.UpdateTime = time.Now()
	update, err := sql.Update(&one)
	if err != nil {
		logs.Error("数据库更新 Article 失败！", err)
		return -1, false
	}
	return update, true
}

func (service *ArticleService) UpdateCommentNum(id int) (int64, bool) {
	one := models.Article{Id: id}
	err := sql.Read(&one)
	if err != nil {
		logs.Error("数据库更新 Article 评论数失败！找不到id: %d", id)
		return -1, false
	}
	count, err := sql.QueryTable(models.ArticleComment{}).Filter("isDelete", 0).Filter("articleId", id).Count()
	if err != nil {
		logs.Error("数据库查询 ArticleComment 总数失败！", err)
		return -1, false
	}
	one.CommentNum = int(count)
	update, err := sql.Update(&one)
	if err != nil {
		logs.Error("数据库更新 Article 评论数失败！", err)
		return -1, false
	}
	return update, true
}

func (service *ArticleService) IncreaseThumbsup(id int) (int64, bool) {
	one := models.Article{Id: id}
	err := sql.Read(&one)
	if err != nil {
		logs.Error("数据库更新 Article 点赞数失败！找不到id: %d", id)
		return -1, false
	}
	one.ThumbsupNum += 1
	update, err := sql.Update(&one)
	if err != nil {
		logs.Error("数据库更新 Article 点赞数失败！", err)
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

func (service *ArticleService) List(articleDto *models.ArticleDto) ([]models.Article, int64, bool) {
	qs := sql.QueryTable(models.Article{})
	qs = qs.Limit(models.GetLimit(articleDto.Page, articleDto.Size))
	qs = qs.Filter("isDelete", false).OrderBy("-updateTime").RelatedSel()

	if strings.Trim(articleDto.Q, "") != "" {
		qs = qs.Filter("title__icontains", articleDto.Q)
	}
	if articleDto.GroupId != 0 {
		qs = qs.Filter("group", articleDto.GroupId)
	}
	var sqlData []models.Article
	_, err := qs.All(&sqlData, "Id", "Group", "Title", "ThumbsupNum", "CommentNum", "CreateTime", "UpdateTime", "IsComment")
	if err != nil {
		logs.Error("数据库查询 Article 列表失败！", err)
		return sqlData, 0, false
	}
	count, _ := qs.Count()
	return sqlData, count, true
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

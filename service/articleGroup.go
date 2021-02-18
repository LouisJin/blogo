package service

import (
	"blogo/models"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

type ArticleGroupService struct {
}

func (service *ArticleGroupService) Insert(articleGroupDto *models.ArticleGroupDto) (int64, bool) {
	one := models.ArticleGroup{
		Name:       articleGroupDto.Name,
		CreateTime: time.Now(),
	}
	insert, err := sql.Insert(&one)
	if err != nil {
		logs.Error("数据库插入 ArticleGroup 失败！", err)
		return -1, false
	}
	return insert, true
}

func (service *ArticleGroupService) Update(articleGroupDto *models.ArticleGroupDto) (int64, bool) {
	one := models.ArticleGroup{Id: articleGroupDto.Id}
	err := sql.Read(&one)
	if err != nil {
		logs.Error("数据库更新 ArticleGroup 失败！找不到id: %d", articleGroupDto.Id)
		return -1, false
	}
	one.Name = articleGroupDto.Name
	one.UpdateTime = time.Now()
	update, err := sql.Update(&one)
	if err != nil {
		logs.Error("数据库更新 ArticleGroup 失败！", err)
		return -1, false
	}
	return update, true
}

func (service *ArticleGroupService) List(pageDto *models.PageDto) ([]models.ArticleGroup, bool) {
	qs := sql.QueryTable(models.ArticleGroup{})
	qs = qs.Limit(models.GetLimit(pageDto.Page, pageDto.Size))

	var sqlData []models.ArticleGroup
	_, err := qs.All(&sqlData)
	if err != nil {
		logs.Error("数据库查询 ArticleGroup 列表失败！", err)
		return sqlData, false
	}
	return sqlData, true
}

func (service *ArticleGroupService) Delete(id int) (int64, bool) {
	one := models.ArticleGroup{Id: id}
	delete, err := sql.Delete(&one)
	if err != nil {
		logs.Error("数据库删除 ArticleGroup 失败！%d, %s", id, err)
		return -1, false
	}
	return delete, true
}

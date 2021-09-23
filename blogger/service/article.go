package service

import (
	"gitlab.renrenchongdian.com/studygo/blogger/dao/db"
	"gitlab.renrenchongdian.com/studygo/blogger/model"
)

// 获取文章和对应分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleList) <= 0 {
		return
	}

	// 2.获取文章对应分类（多个）
	categoryIds := getCategoryIds(articleList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}

	// 3.聚合
	for _, article := range articleList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if category.CategoryId == categoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

// 根据多个文章的id获取多个分类的id
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	for _, articleInfo := range articleInfoList {
		categoryId := articleInfo.CategoryId
		for _, id := range ids {
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

// 根据分类id，获取该类文章及他们的分类信息
func GetArticleRecordById(articleId int64, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleList, err := db.GetArticleListByCategoryId(articleId, pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleList) <= 0 {
		return
	}

	// 2.获取文章对应分类（多个）
	categoryIds := getCategoryIds(articleList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}

	// 3.聚合
	for _, article := range articleList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if category.CategoryId == categoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

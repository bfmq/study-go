package db

import (
	"fmt"

	"gitlab.renrenchongdian.com/studygo/blogger/model"
)

// 添加文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	if article == nil {
		return
	}

	sqlStr := `insert into 
				article
					(content,summary,title,username,category_id,view_count,comment_count) 
				values 
					(?,?,?,?,?,?,?)`

	result, err := DB.Exec(sqlStr, article.Content, article.Summary, article.Title, article.UserName,
		article.ArticleInfo.CategoryId, article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		fmt.Printf("InsertCategory failed,err:%s\n", err)
		return
	}
	articleId, err = result.LastInsertId()
	return
}

// 获取单个文章分类
func GetArticleById(articleId int64) (article *model.ArticleDetail, err error) {
	sqlStr := "select id,category_name,category_no from article where id=?"
	err = DB.Get(article, sqlStr, articleId)
	return
}

// 获取多个文章分类
func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum <= 0 || pageSize <= 0 {
		return
	}
	sqlStr := `select 
					id,summary,title,username,category_id,view_count,comment_count,create_time 
				from 
					article 
				where 
					status = 1 
				order by 
					create_time desc
				limit
					?,?`
	err = DB.Select(&articleList, sqlStr, pageNum, pageSize)
	return
}

// 获取单个文章详情
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		return
	}

	sqlStr := `select
					id,summary,title,username,category_id,view_count,comment_count,create_time,content
				from
					article
				where
					id = ?
				and
					status = 1`
	err = DB.Get(articleDetail, sqlStr, articleId)
	return
}

// 获取多个文章详情
func GetArticleListByCategoryId(categoryId int64, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum <= 0 || pageSize <= 0 {
		return
	}
	sqlStr := `select 
					id,summary,title,username,category_id,view_count,comment_count,create_time 
				from 
					article 
				where 
					status = 1
				and
					category_id = ?
				order by 
					create_time desc
				limit
					?,?`
	err = DB.Select(&articleList, sqlStr, categoryId, pageNum, pageSize)
	return
}

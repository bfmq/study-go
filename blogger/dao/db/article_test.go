package db

import (
	"fmt"
	"testing"
	"time"

	"gitlab.renrenchongdian.com/studygo/blogger/model"
)

func init() {
	dns := "root:root@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// t *testing.T
func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 1
	article.ArticleInfo.CreateTime = time.Now()
	article.Content = "单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章单元测试文章"
	article.ArticleInfo.Title = "单元测试文章"
	article.ArticleInfo.Summary = "单测"
	article.ArticleInfo.ViewCount = 1
	articleId, err := InsertArticle(article)
	if err != nil {
		return
	}
	t.Logf("articleId is %d\n", articleId)
}

// 获取多个文章分类
func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		fmt.Printf(",err:%s\n", err)
		return
	}
	t.Logf("articleList is %d\n", len(articleList))
}

// 获取单个文章详情
func TestGetArticleDetail(t *testing.T) {
	articleDetail, err := GetArticleDetail(1)
	if err != nil {
		fmt.Printf(",err:%s\n", err)
		return
	}
	t.Logf("articleList is %#v\n", articleDetail)
}

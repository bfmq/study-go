package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.renrenchongdian.com/studygo/blogger/service"
)

func IndexHandle(c *gin.Context) {
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
	}

	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":          http.StatusOK,
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

func CategoryListHandle(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
	}
	articleRecordList, err := service.GetArticleRecordById(categoryId, 0, 15)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
	}

	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":          http.StatusOK,
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.renrenchongdian.com/studygo/blogger/model"
)

// 添加分类
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlStr := "insert into category (category_name,category_no) values (?,?)"
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		fmt.Printf("InsertCategory failed,err:%s\n", err)
		return
	}
	categoryId, err = result.LastInsertId()
	return
}

// 获取单个文章分类
func GetCategoryById(categoryId int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlStr := "select id,category_name,category_no from category where id=?"
	err = DB.Get(category, sqlStr, categoryId)
	return
}

// 获取多个文章分类
func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	sqlStr, args, err := sqlx.In("select id,category_name,category_no from category where id in (?)", categoryIds)
	if err != nil {
		fmt.Printf("GetCategoryList failed,err:%s\n", err)
		return
	}
	err = DB.Select(&categoryList, sqlStr, args...)
	return
}

// 获取所有文章分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sqlStr := "select id,category_name,category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlStr)
	return
}

package db

import (
	"testing"
)

func init() {
	dns := "root:root@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds = []int64{1, 2, 3}
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("id:%d,category:%#v", v.CategoryId, v)
	}

}

func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("id:%d,category:%#v", v.CategoryId, v)
	}
}

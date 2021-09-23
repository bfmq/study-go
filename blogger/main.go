package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.renrenchongdian.com/studygo/blogger/controller"
	"gitlab.renrenchongdian.com/studygo/blogger/dao/db"
)

func main() {
	router := gin.Default()
	dns := "root:root@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}

	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.CategoryListHandle)

	router.Run(":8000")
}

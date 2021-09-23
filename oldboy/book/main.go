package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func bookListHandler(c *gin.Context) {
	bookList, err := queryAllBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": bookList,
	})
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	{
		r.GET("/book/list", bookListHandler)
	}
	_ = r.Run(":8000")
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func IPCodeMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 要求3：获取客户端地址、本次返回状态码，并打印日志到控制台
		ip := c.Request.RemoteAddr
		code := c.Writer.Status()
		log.Println("客户端地址---------------->", ip)
		log.Println("给客户端的返回码---------------->", code)
	}
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(IPCodeMiddle())

	{
		// 要求4：提供healthz接口
		router.GET("/healthz/", healthzGetHandle)
	}

	router.Run(":8000")
}

func healthzGetHandle(c *gin.Context) {
	// 要求1：将request header封回response header
	for k, v := range c.Request.Header {
		c.Header(k, v[0])
	}

	// 要求2：获取当前系统env相关环境信息封回response header
	version := getVersion()
	c.Header("VERSION", version)

	c.JSON(http.StatusOK, gin.H{
		// 要求4：healthz接口返200码
		"code": http.StatusOK,
	})
}

// 写代码时没有VERSION/Version/version这几个环境变量，可能系统变量名不太一样吧
func getVersion() string {
	version := os.Getenv("VERSION")
	if version == "" {
		os.Setenv("VERSION", "test version")
		version = os.Getenv("VERSION")
	}
	return version
}

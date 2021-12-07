package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func IPCodeMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
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
		router.GET("/healthz/", healthzGetHandle)
		router.GET("/delay/", delayGetHandle)
	}

	router.Run(":8000")
}

func healthzGetHandle(c *gin.Context) {
	for k, v := range c.Request.Header {
		c.Header(k, v[0])
	}

	version := getVersion()
	c.Header("VERSION", version)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func getVersion() string {
	version := os.Getenv("VERSION")
	if version == "" {
		os.Setenv("VERSION", "test version")
		version = os.Getenv("VERSION")
	}
	return version
}

// 要求1：随机等待1-2000ms返回
func delayGetHandle(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	sleepTime := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
	c.JSON(http.StatusOK, gin.H{
		"code":      http.StatusOK,
		"sleepTime": sleepTime,
	})
}

package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/chenjiandongx/ginprom"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(ginprom.PromMiddleware(nil))

	{
		router.GET("/healthz/", healthzGetHandle)
		// 要求2：增加metrics监控
		router.GET("/metrics/", ginprom.PromHandler(promhttp.Handler()))
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

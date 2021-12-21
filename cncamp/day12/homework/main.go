package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	jaeger "github.com/uber/jaeger-client-go"
)

const ginServerName = "httpserver"

func main() {
	tracer, closer := jaeger.NewTracer(
		ginServerName,
		jaeger.NewConstSampler(true),
		jaeger.NewInMemoryReporter(),
	)
	defer closer.Close()

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(ginprom.PromMiddleware(nil))
	// 注册链路追踪中间件
	router.Use(ginhttp.Middleware(tracer))

	{
		router.GET("/healthz/", healthzGetHandle)
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

func delayGetHandle(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	sleepTime := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
	c.JSON(http.StatusOK, gin.H{
		"code":      http.StatusOK,
		"sleepTime": sleepTime,
	})
}

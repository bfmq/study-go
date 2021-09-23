package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "err",
		})
		c.Abort()
	}
}

func main() {
	r := gin.Default()
	{
		r.GET("/login", func(c *gin.Context) {
			c.SetCookie("abc", "123", 60, "/", "127.0.0.1", false, false)
			c.String(200, "Login success")
		})
		r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"home": "data",
			})
		})
	}

	r.Run(":8000")
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.LoadHTMLFiles("login.html", "index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username := c.PostForm("username")
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "***"
		}
		//输出json结果给调用方
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":  "ok",
			"username": username,
			"password": password,
		})
	})
	r.Run(":9999")
}

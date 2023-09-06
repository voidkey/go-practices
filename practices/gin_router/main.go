package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
/*
Gin框架中的路由使用的是httprouter这个库。

其基本原理就是构造一个路由地址的前缀树。
*/
func router(r *gin.Engine){
	//普通路由
	r.GET("/index", func(c *gin.Context) {...})
	r.GET("/login", func(c *gin.Context) {...})
	r.POST("/login", func(c *gin.Context) {...})

	//可以匹配所有请求方法的Any方法
	r.Any("/test", func(c *gin.Context) {...})
	//为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码，下面的代码为没有匹配到路由的请求都返回views/404.html页面。
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	r.Run(":9999")
}

func routerGroup(r *gin.Engine){
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {...})
		userGroup.GET("/login", func(c *gin.Context) {...})
		userGroup.POST("/login", func(c *gin.Context) {...})

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		shopGroup.GET("/cart", func(c *gin.Context) {...})
		shopGroup.POST("/checkout", func(c *gin.Context) {...})
	}
	shopGroup = r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		shopGroup.GET("/cart", func(c *gin.Context) {...})
		shopGroup.POST("/checkout", func(c *gin.Context) {...})
		// 嵌套路由组
		xx := shopGroup.Group("xx")
		xx.GET("/oo", func(c *gin.Context) {...})
	}
	r.Run(":9999")
}

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	router(r)
}

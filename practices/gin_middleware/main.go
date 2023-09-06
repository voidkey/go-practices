package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// StatCost 是一个统计耗时请求耗时的中间件
// Gin中的中间件必须是一个gin.HandlerFunc类型
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("StatCost In")
		start := time.Now()
		c.Set("name", "ZHUOFEI") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		fmt.Println(cost)
		fmt.Println("StatCost Out")
	}
}

func indexHandler(c *gin.Context) {
	fmt.Println("index")
	c.JSON(200, gin.H{
		"msg": "index",
	})
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck {
			//存放具体的逻辑
			//是否登录的判断
			//if 是登录用户
			//c.Next()
			//else
			//c.Abort()
		} else {
			c.Next()
		}

	}
}

func main() {
	/*
		gin.Default()默认使用了Logger和Recovery中间件，其中：

		Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
		Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
		如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

		当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。
	*/
	r := gin.New()
	// 注册一个全局中间件
	//在gin框架中，我们可以为每个路由添加任意数量的中间件。
	r.Use(StatCost(), authMiddleware(true))
	r.GET("/index", StatCost(), indexHandler)

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	// 给/test2路由单独注册中间件（可注册多个）
	r.GET("/test2", StatCost(), func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	//为路由组注册中间件
	//写法1
	shopGroup := r.Group("/shop", StatCost())
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		...
	}
	//写法2
	shopGroup = r.Group("/shop")
	shopGroup.Use(StatCost())
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		...
	}
	r.Run(":9999")

}

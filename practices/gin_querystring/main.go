package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器发送的请求携带的querystring参数
		//queryStr := c.Query("query")
		//queryStr := c.DefaultQuery("query", "somebody") //取不到就用默认值
		queryStr, ok := c.GetQuery("query")
		if !ok {
			queryStr = "somebody"
		}

		age := c.Query("age")

		c.JSON(http.StatusOK, gin.H{
			"name": queryStr,
			"age":  age,
		}) //返回的json按照字典序排序
	})
	r.Run(":9999")
}

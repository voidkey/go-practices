package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name    string
	Message string
	Age     int
}

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法1 使用map
		//data := map[string]interface{}{
		//	"name":    "ZHUO",
		//	"message": "HAHAHA",
		//	"age":     100,
		//}
		data := gin.H{"name": "FEI", "message": "HAHAHA", "age": 18}
		//方法2 使用struct
		c.JSON(http.StatusOK, data)
	})

	r.GET("/json2", func(c *gin.Context) {
		data := User{
			Name:    "XUE",
			Message: "HEHEHE",
			Age:     99,
		}
		c.JSON(200, data)
	})

	err := r.Run(":9999")
	if err != nil {
		return
	}
}

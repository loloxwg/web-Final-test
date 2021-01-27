package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	engine:= gin.Default()
	engine.Use(RequestInfos())
	//query解析
	engine.GET("/query", func(context *gin.Context) {
		fmt.Println("中间件的使用方法")
		context.JSON(404,map[string]interface{}{
			"code":1,
			"msg":context.FullPath(),
		})
	})
	engine.GET("/hello", func(context *gin.Context) {
		//todo
	})
	engine.Run(":9001")
}

//打印请求信息的中间件
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		path:=context.FullPath()
		method:=context.Request.Method
		fmt.Println("请求path：",path)
		fmt.Println("请求method：",method)
		fmt.Println("状态码：",context.Writer.Status())

		context.Next()//

		fmt.Println("状态码：",context.Writer.Status())
	}
}
package router

import (
	"gin_demo/app/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	v1 := r.Group("v1/user")
	{
		v1.POST("/signup", controller.UserSignUp)
		v1.GET("/info/get", controller.GetUserInfo)
	}
	r.Run(":8099")
}

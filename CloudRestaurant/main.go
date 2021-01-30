package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tools"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := tools.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	app := gin.Default()
	registerRouter(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine)  {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
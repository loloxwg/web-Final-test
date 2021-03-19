package controller

import "github.com/gin-gonic/gin"

type MessageCategoryController struct {
}

func (fcc *MessageCategoryController) Router(engine *gin.Engine) {
	engine.GET("/api/message_category", fcc.messageCategory)

}

func (fcc *MessageCategoryController) messageCategory(ctx *gin.Context) {
	//调用service功能获取文章信息
}

package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 0
	FAILED  int = 1
)

//普通成功返回
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    SUCCESS,
		"message": "成功",
		"data":    v,
	})
}
//普通操作失败
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  v,
	})
}

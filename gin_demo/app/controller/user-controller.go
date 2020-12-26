package controller

import (
	"fmt"
	"gin_demo/app/jsonapi"
	"gin_demo/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserSignUp(c *gin.Context) {
	respMap := make(map[string]interface{})

	params := new(jsonapi.UserSignUpReq)

	params.Account = c.PostForm("account")
	params.Passwd = c.PostForm("passwd")
	params.RightPw = c.PostForm("right_pw")

	if err := service.UserSrvCli.SignUpSrv(params); err != nil {
		fmt.Println(err)
		respMap["message"] = "注册失败！"
		respMap["code"] = -2001
		c.JSON(http.StatusOK, respMap)
	} else {
		respMap["message"] = "OK"
		respMap["code"] = 0
		c.JSON(http.StatusOK, respMap)
	}
}

func GetUserInfo(c *gin.Context) {
	respMap := make(map[string]interface{})

	uid := c.PostForm("uid")
	userid, _ := strconv.ParseInt(uid, 10, 64)

	resp := service.UserSrvCli.GetUserSrv(userid)
	respMap["data"] = resp
	respMap["message"] = "OK"
	respMap["code"] = 0
	c.JSON(http.StatusOK, respMap)

}

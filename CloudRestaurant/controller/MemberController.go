package controller

import (
	"CloudRestaurant/param"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)
	engine.OPTIONS("/api/login_sms",mc.smslogin)
}

// http://localhost:8090/api/sendcode?phone=15922758321
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//发送验证码
	phone,exist:=context.GetQuery("phone")
	if !exist {
		tool.Failed(context,"参数解析失败")
		//context.JSON(200,map[string]interface{}{
		//	"code":0,
		//	"msg":"参数解析失败",
		//})
		return
	}
	ms:=service.MemberService{}
	isSend:=ms.SendCode(phone)
	if isSend {
		tool.Success(context,"发送成功")
		return
	}
	tool.Failed(context,"发送失败")
}
//手机号+短信 登录的方法
func (mc *MemberController)smslogin(context *gin.Context)  {
	var smsLoginParam param.SmsLoginParam
	err:=tool.Decode(context.Request.Body,&smsLoginParam)
	if err!=nil{
		tool.Failed(context,"参数解析失败")
		return
	}
	//完成手机+验证码登录
	us:= service.MemberService{}
	member:=us.SmsLogin(smsLoginParam)
	if member !=nil {
		tool.Success(context,member)
		return
	}
	tool.Failed(context,"登陆失败")

}
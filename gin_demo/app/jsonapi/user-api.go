package jsonapi

type UserSignUpReq struct {
	Account string `form:"account" json:"account"`
	Passwd  string `form:"passwd" json:"passwd"`
	RightPw string `form:"right_pw" json:"right_pw"`
}

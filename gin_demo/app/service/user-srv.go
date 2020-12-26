package service

import (
	"crypto/md5"
	"fmt"
	"gin_demo/app/jsonapi"
	"gin_demo/model/t_user"
	"time"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u UserService) SignUpSrv(req *jsonapi.UserSignUpReq) error {
	fmt.Println(req)
	sdata := []byte(req.Account + time.December.String())
	slat := md5.Sum(sdata)
	saltstr := fmt.Sprintf("%x", slat)

	pdata := []byte(req.RightPw + saltstr)
	pw := md5.Sum(pdata)

	userInfo := &t_user.Entity{
		Uid:        1,
		Account:    req.Account,
		Salt:       saltstr,
		Passwd:     fmt.Sprintf("%x", pw),
		Nickname:   "bonner",
		Sign:       "......",
		CreateTime: time.Now(),
	}

	if err := t_user.CreateOne(userInfo); err != nil {
		return err
	}

	return nil
}

func (u UserService) GetUserSrv(uid int64) []t_user.Entity {

	resp := t_user.GetOne(uid)
	return resp
}

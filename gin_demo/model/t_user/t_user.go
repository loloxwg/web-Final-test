package t_user

import (
	"fmt"
	"gin_demo/util"
)

func CreateOne(info *Entity) error {
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into t_user(uid,account,salt,passwd,nickname,sign,create_time) values(?,?,?,?,?,?,?)", info.Uid, info.Account, info.Salt, info.Passwd, info.Nickname, info.Sign, info.CreateTime)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func GetOne(uid int64) (resp []Entity) {
	db, err := util.OpenDB()
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return nil
	}

	err = db.Select(&resp, "select uid,account,salt,passwd,nickname,sign from t_user where uid=? ;", uid)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil
	}
	db.Close()
	return resp
}

func Modify(info *Entity) int64 {
	db, err := util.OpenDB()
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return 0
	}

	res, err := db.Exec("update t_user set passwd=?,nickname=?,sign=? where uid=?", info.Passwd, info.Nickname, info.Sign, info.Uid)
	if err != nil {
		fmt.Println(" exec failed, ", err)
		return 0
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
		return 0
	}

	db.Close()
	return row
}

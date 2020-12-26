package t_user

import "time"

type Entity struct {
	Id         int64     `db:"id"`
	Uid        int64     `db:"uid"`
	Account    string    `db:"account"`
	Salt       string    `db:"salt"`
	Passwd     string    `db:"passwd"`
	Nickname   string    `db:"nickname"`
	Sign       string    `db:"sign"`
	CreateTime time.Time `db:"create_time"`
	ModifyTime time.Time `db:"modify_time"`
}

package model

type MessageCategory struct {
	//类别id
	Id int64 `xorm:"pk autoincr" json:"id"`
	//类别标题
	Title string `xorm:"varchar(20)" json:"title"`
	//描述
	Description string `xorm:"varchar(30)" json:"description"`
	//图片
	ImageUrl string `xorm:"varchar(255)" json:"image_url"`
	//链接
	LinkUrl string `xorm:"varchar(255)" json:"link_url"`
	//日期
	PostDate string `xorm:"varchar(30)" json:"post_date"`
	//文章内容
	PostContent string `xorm:"varchar(255)" json:"post_content"`
	//服务状态
	IsInServing bool `json:"is_in_serving"`
}

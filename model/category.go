package model

type Category struct {
	ID    int64  `json:"cid" gorm:"primaryKey;column:cid"`
	Cname string `json:"cname" gorm:"column:cname"`
}

func (Category) TableName() string {

	return "t_category"
}

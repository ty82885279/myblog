package model

type User struct {
	UserID   int64  `gorm:"primaryKey;column:user_id" json:"userId"`
	UserName string `gorm:"column:user_name" json:"userName"`
	Password string `gorm:"column:psw" json:"-"`
}

func (User) TableName() string {
	return "t_user"
}

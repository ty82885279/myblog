package model

type Comment struct {
	ID          int64      `json:"id" gorm:"column:id"`
	ArticleID   int64      `json:"articleId" gorm:"column:article_id"`
	FromUserID  int64      `json:"-" gorm:"column:from_userid"`
	IsDel       int64      `json:"isDel" gorm:"column:is_del"`
	ReplycCount int64      `json:"replyCount" gorm:"column:reply_nums"`
	Content     string     `json:"content" gorm:"column:content"`
	CreatedAt   TimeNormal `json:"createTime" gorm:"column:create_time"`
	UpdatedAt   TimeNormal `json:"-" gorm:"column:update_time"`
	User        User       `json:"fromUser" gorm:"foreignkey:FromUserID"`
}

func (c Comment) TableName() string {
	return "t_comment"
}

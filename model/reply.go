package model

type Reply struct {
	ID         int64      `json:"id" gorm:"column:id"`
	ArticleID  int64      `json:"articleId" gorm:"column:article_id"`
	CommentID  int64      `json:"commentId" gorm:"column:comment_id"`
	ReplyID    int64      `json:"replyId" gorm:"column:reply_id"`
	FromUserID int64      `json:"-" gorm:"column:from_userid"`
	ToUserID   int64      `json:"-" gorm:"column:to_userid"`
	Isdel      int64      `json:"isDel" gorm:"column:is_del"`
	Content    string     `json:"content" gorm:"column:content"`
	CreatedAt  TimeNormal `json:"createTime" gorm:"column:create_time"`
	UpdatedAt  TimeNormal `json:"updateTime" gorm:"column:update_time"`
	FromUser   User       `json:"fromUser" gorm:"foreignkey:FromUserID"`
	ToUser     User       `json:"toUser" gorm:"foreignkey:ToUserID"`
}

func (Reply) TableName() string {
	return "t_reply"
}

package model

// 文章表
type Article struct {
	ID          int64      `json:"id" gorm:"column:id"`
	CID         int64      `json:"-" gorm:"column:cid"`
	Title       string     `json:"title" gorm:"column:title"`
	Content     string     `json:"content" gorm:"column:content"`
	Status      int64      `json:"status" gorm:"column:is_del"`
	CreatedAt   TimeNormal `json:"createTime" gorm:"column:create_time"`
	UpdatedAt   TimeNormal `json:"updateTime" gorm:"column:update_time"`
	Category    Category   `json:"category" gorm:"-"`
	CommentNums int64      `json:"commentNums" gorm:"column:comment_nums"`
	PraiseNums  int64      `json:"praiseNums" gorm:"column:praise_nums"`
	ReadNums    int64      `json:"readNums" gorm:"column:read_nums"`
	IsStar      int64      `json:"isStar" gorm:"-"`
}

func (Article) TableName() string {
	return "t_article"
}

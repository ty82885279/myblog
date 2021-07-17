package model

type ParamSignUp struct {
	UserName   string `json:"username" form:"username" binding:"required"`                      // 用户名 必填
	Password   string `json:"password" form:"password" binding:"required,min=6,max=12"`         // 密码   必填
	RePassword string `json:"repassword" form:"repassword" binding:"required,eqfield=Password"` //确认密码 必填
}

type ParamLogin struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// 文章列表
type ParamArticleList struct {
	CID    int
	UserId int64
	Page   int
	Size   int
}

// 文章详情
type ParamArticleDetail struct {
	UserID    int64
	ArticleID int64
}
type ParamAddArticle struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Cid     int64  `json:"cid" binding:"required"`
}
type ParamArticleSetStart struct {
	ArticleId int64
	UserID    int64
}

// 添加评论参数
type ParamAddComment struct {
	Content   string `json:"content" binding:"required"`
	ArticleID int64  `json:"articleID" binding:"required"`
	UserID    int64
}

// 获取某个文章下的评论
type ParamCommentList struct {
	ArticleID int `json:"cid" binding:"required"`
	Page      int `json:"page" binding:"required"`
	Size      int `json:"size" binding:"required"`
}

// 添加回复
type ParamAddReply struct {
	Content    string `json:"content" binding:"required"`
	ArticleID  int64  `json:"articleID" binding:"required"`
	CommentID  int64  `json:"commentId" binding:"required"`
	ReplyID    int64  `json:"replyId"`
	FromUserID int64
	ToUserID   int64 `json:"toUserId"`
}

// 获取某个评论下的回复
type ParamReplyList struct {
	CommentID int `json:"commentId" binding:"required"`
	Page      int `json:"page" binding:"required"`
	Size      int `json:"size" binding:"required"`
}

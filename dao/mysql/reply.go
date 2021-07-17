package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"myblog/model"
)

// 添加回复
func CreateReply(p *model.ParamAddReply) (err error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("recover---%v\n", err)
		}
	}()
	reply := model.Reply{
		ArticleID:  p.ArticleID,
		CommentID:  p.CommentID,
		FromUserID: p.FromUserID,
		ToUserID:   p.ToUserID,
		Content:    p.Content,
		ReplyID:    p.ReplyID,
	}
	if err = db.Debug().Model(&model.Reply{}).Create(&reply).Error; err != nil {
		panic(err)
	}
	//
	if err = IncreaseCommentNum(p.ArticleID); err != nil {
		zap.L().Error("uodate article's comment count err", zap.Error(err))
		return err
	}
	if err = IncreaseCommentReplyNum(p.CommentID); err != nil {
		zap.L().Error("uodate comment's reply count err", zap.Error(err))
		return err
	}
	return
}

// 获取评论下的回复
func GetReplyListByCommentId(p *model.ParamReplyList) (replyList []*model.Reply, total int64, err error) {

	if err = db.Debug().Model(&model.Reply{}).Preload("FromUser").Preload("ToUser").Where("comment_id = ?", p.CommentID).Count(&total).Limit(p.Size).Offset((p.Page - 1) * p.Size).Order("create_time desc").Find(&replyList).Error; err != nil {
		panic(err)
	}
	return
}

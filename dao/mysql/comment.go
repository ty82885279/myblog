package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"myblog/model"
)

// 添加评论
func CreateComment(p *model.ParamAddComment) (err error) {
	comment := model.Comment{
		ArticleID:  p.ArticleID,
		FromUserID: p.UserID,
		Content:    p.Content,
	}

	// 添加评论
	if err = db.Debug().Model(&model.Comment{}).Create(&comment).Error; err != nil {
		return
	}
	// 更新文章评论数量
	if err = IncreaseCommentNum(p.ArticleID); err != nil {
		zap.L().Error("uodate article's comment count err", zap.Error(err))
		return err
	}
	return
}

// 通过文章id获取评论列表
func GetCommentsByArticleId(p model.ParamCommentList) (comments []*model.Comment, total int64, err error) {

	if err = db.Debug().Model(&model.Comment{}).Preload("User").Where("article_id = ?", p.ArticleID).Count(&total).Limit(p.Size).Offset((p.Page - 1) * p.Size).Order("create_time desc").Find(&comments).Error; err != nil {
		panic(err)
	}
	return
}

// 更新评论回复数+1
func IncreaseCommentReplyNum(commentID int64) (err error) {
	if err = db.Debug().Model(&model.Comment{}).Where("id = ?", commentID).Update("reply_nums", gorm.Expr("reply_nums + ?", 1)).Error; err != nil {
		panic(err)
	}
	return
}

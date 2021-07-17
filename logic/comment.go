package logic

import (
	"myblog/dao/mysql"
	"myblog/model"
)

// 添加评论
func CreateComment(p *model.ParamAddComment) error {
	return mysql.CreateComment(p)
}

// 通过文章id获取评论列表
func GetCommentsByArticleId(p model.ParamCommentList) ([]*model.Comment, int64, error) {

	return mysql.GetCommentsByArticleId(p)
}

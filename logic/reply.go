package logic

import (
	"myblog/dao/mysql"
	"myblog/model"
)

func CreateReply(p *model.ParamAddReply) error {
	return mysql.CreateReply(p)
}
func GetReplyListByCommentId(p *model.ParamReplyList) ([]*model.Reply, int64, error) {

	return mysql.GetReplyListByCommentId(p)

}

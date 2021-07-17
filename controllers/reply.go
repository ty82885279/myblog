package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"myblog/logic"
	"myblog/model"
	"strconv"
)

func CreateReply(c *gin.Context) {

	CurrentUserID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p := &model.ParamAddReply{
		ReplyID:    0,
		FromUserID: CurrentUserID,
	}
	if err = c.ShouldBind(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errMsg(errs.Translate(trans)))
		return
	}
	if err = logic.CreateReply(p); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c)

}

func GetReplyListByCommentId(c *gin.Context) {
	commentId, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	p := model.ParamReplyList{
		CommentID: commentId,
		Page:      page,
		Size:      size,
	}

	replys, total, err := logic.GetReplyListByCommentId(&p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithData(c, gin.H{
		"total": total,
		"list":  replys,
	})
}

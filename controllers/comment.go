package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"myblog/logic"
	"myblog/model"
	"strconv"
)

// 添加评论
func CreateComment(c *gin.Context) {

	CurrentUserID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p := model.ParamAddComment{
		UserID: CurrentUserID,
	}
	if err := c.ShouldBind(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, errMsg(errs.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}

	exist := logic.CheckArticleExist(p.ArticleID)
	if !exist {
		ResponseError(c, CodeArticleNotExist)
		return
	}
	if err := logic.CreateComment(&p); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithMsg(c, CodeSuccess, "添加评论成功")
}

//获取某个文章的评论
func GetCommentsByArticleId(c *gin.Context) {

	articleId, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	p := model.ParamCommentList{
		ArticleID: articleId,
		Page:      page,
		Size:      size,
	}

	list, total, err := logic.GetCommentsByArticleId(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithData(c, gin.H{
		"total": total,
		"list":  list,
	})

}

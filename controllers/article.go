package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"myblog/logic"
	"myblog/model"
	"strconv"
)

// 获取分类下的所有文章
func GetArticlesByCategory(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("id"))
	userId := GetUserIdFromToken(c)
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	p := model.ParamArticleList{
		UserId: userId,
		CID:    cid,
		Page:   page,
		Size:   size,
	}
	list, total := logic.GetArticlesByCategoryId(p)
	ResponseSuccessWithData(c, gin.H{
		"total": total,
		"page":  p.Page,
		"size":  p.Size,
		"list":  list,
	})
}

// 获取文章详情
func GetArticleDetail(c *gin.Context) {
	aid, _ := strconv.Atoi(c.Param("id"))
	userId := GetUserIdFromToken(c)
	p := model.ParamArticleDetail{
		ArticleID: int64(aid),
		UserID:    userId,
	}
	article, err := logic.GetArticleDetail(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccessWithData(c, article)
}

// 添加文章
func CreateArticle(c *gin.Context) {
	p := new(model.ParamAddArticle)
	if err := c.ShouldBind(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, errMsg(errs.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	err := logic.CreateArticle(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c)
}

// 编辑文章
func EditArticle(c *gin.Context) {
	article := model.Article{}
	id, _ := strconv.Atoi(c.Param("id"))
	article.ID = int64(id)
	if err := c.ShouldBind(&article); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errMsg(errs.Translate(trans)))
		return
	}
	if err := logic.EditArticle(&article); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c)
}

// 删除文章
func DeleteArticle(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	if err := logic.DeleteArticle(int64(id)); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c)
}

// 文章点赞
func SetStar(c *gin.Context) {
	userId, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	articleId, _ := strconv.Atoi(c.Param("id"))
	p := model.ParamArticleSetStart{
		ArticleId: int64(articleId),
		UserID:    userId,
	}
	info, err := logic.SetStarArticle(&p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithData(c, gin.H{
		"msg":     info.Msg,
		"starNum": info.StarNum,
		"isStar":  info.IsStar,
	})
}

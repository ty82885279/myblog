package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"myblog/dao/mysql"
	"myblog/logic"
	"myblog/model"
	"net/http"
)

func SignupHandler(c *gin.Context) {
	var u = new(model.ParamSignUp)
	err := c.ShouldBind(u)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			zap.L().Error("ShouldBindJSON Err", zap.Error(err))
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errMsg(errs.Translate(trans)))
		return
	}
	userExist := logic.CheckUserExist(u.UserName)
	if userExist {

		ResponseError(c, CodeUserExist)
		return
	}
	err = logic.CreateUser(u)
	if err != nil {
		return
	}
	ResponseSuccessWithCode(c, CodeSignupSuccess)

}

func LoginHandler(c *gin.Context) {

	u := new(model.ParamLogin)
	err := c.ShouldBind(u)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errMsg(errs.Translate(trans)))
		return
	}
	exist := logic.CheckUserExist(u.UserName)
	if !exist {
		ResponseError(c, CodeUserNotExist)
		return
	}
	AToken, RToken, err := logic.UserLogin(u)
	if err != nil {
		if errors.Is(mysql.ErrorInvalidPassword, err) {
			ResponseError(c, CodeInvalidPassword)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":          CodeLoginSuccess,
		"msg":           CodeLoginSuccess.Msg(),
		"token_access":  AToken,
		"token_refresh": RToken,
	})

}

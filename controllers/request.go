package controllers

import (
	"errors"
	"myblog/pkg/jwt"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登陆")

//  GetCurrentUser  获取当前登陆的用户id
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {

	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	return
}

//获取分页信息
func GetPageInfo(c *gin.Context) (int64, int64) {

	var (
		err        error
		page, size int64
	)

	//获取分页
	page, err = strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		zap.L().Error("PostListHandler.getPage err", zap.Error(err))
		page = 1
	}
	size, err = strconv.ParseInt(c.Query("size"), 10, 64)
	if err != nil {
		size = 10

	}
	return page, size
}

// 通过token获取useID
func GetUserIdFromToken(c *gin.Context) (userId int64) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return 0
	}
	parts := strings.SplitN(authHeader, " ", 2)
	mc, _ := jwt.ParseToken(parts[1])
	return mc.UserID
}

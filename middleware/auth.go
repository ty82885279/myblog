package middleware

import (
	"myblog/controllers"
	"myblog/pkg/jwt"
	"strings"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {

			controllers.ResponseError(c, controllers.CodeTokenNotExist)
			c.Abort()
			zap.L().Error("请求中缺少Auth Token")
			return
		}
		//按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {

			controllers.ResponseError(c, controllers.CodeInvalidToken)
			zap.L().Error("请求头中Token格式错误")
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		//将当前userID保存到请求的上下文c中
		c.Set(controllers.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}

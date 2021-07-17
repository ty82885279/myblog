package routers

import (
	"myblog/controllers"
	"myblog/logger"
	"myblog/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(middleware.Cors())
	router := r.Group("/api/v1")
	{
		// 注册
		router.POST("/user/signup", controllers.SignupHandler)
		// 登录
		router.POST("/user/login", controllers.LoginHandler)

		// 获取所有分类
		router.GET("/category", controllers.CategoryList)

		// 获取分类下所有文章
		router.GET("/category/:id", controllers.GetArticlesByCategory)

		//获取文章详情
		router.GET("/article/:id", controllers.GetArticleDetail)
		// 评论相关
		// 获取某一文章下的评论
		router.GET("/comment/list/:id", controllers.GetCommentsByArticleId)
		router.POST("comment/replyList/:id", controllers.GetReplyListByCommentId)

	}
	// 需要鉴权
	auth := r.Group("/api/v1")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		// 文章相关
		auth.POST("/article", controllers.CreateArticle)
		auth.PUT("/article/:id", controllers.EditArticle)
		auth.DELETE("article/:id", controllers.DeleteArticle)
		auth.POST("article/star/:id", controllers.SetStar)

		// 评论相关
		auth.POST("/comment", controllers.CreateComment)
		auth.POST("/comment/reply", controllers.CreateReply)

	}
	return r
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"myblog/logic"
)

// 获取所有分类
func CategoryList(c *gin.Context) {
	categoryList := logic.CategoryList()
	ResponseSuccessWithData(c, categoryList)
}

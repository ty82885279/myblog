package logic

import (
	"myblog/dao/mysql"
	"myblog/model"
)

func CategoryList() (list []model.Category) {

	return mysql.CategoryList()
}

package mysql

import (
	"myblog/model"
)

func CategoryList() (list []model.Category) {
	list = make([]model.Category, 0, 10)
	err := db.Debug().Table("t_category").Find(&list).Error
	if err != nil {
		panic(err)
	}
	return list
}

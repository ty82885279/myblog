package mysql

import "errors"

var (
	ErrorQuery             = errors.New("查询错误")
	ErrorUserNotExist      = errors.New("用户不存在")
	ErrorInvalidPassword   = errors.New("密码错误")
	ErrorArtiCleIDNotExist = errors.New("文章不存在")
)

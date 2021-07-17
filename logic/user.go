package logic

import (
	"myblog/dao/mysql"
	"myblog/model"
)

func CheckUserExist(userName string) bool {
	return mysql.CheckUserExist(userName)
}

func CreateUser(user *model.ParamSignUp) error {

	return mysql.CreateUser(user)
}

func UserLogin(user *model.ParamLogin) (AToken, RToken string, err error) {

	return mysql.UserLogin(user)
}

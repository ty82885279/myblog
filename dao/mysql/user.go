package mysql

import (
	"go.uber.org/zap"
	"myblog/model"
	"myblog/pkg/jwt"
	"myblog/utils"
)

func CheckUserExist(userName string) bool {

	var u = new(model.User)
	err := db.Debug().Where("user_name = ?", userName).Find(u).Error
	if err != nil {
		panic(err)
	}
	if u.UserID > 0 {
		return true
	}
	return false

}

func CreateUser(user *model.ParamSignUp) (err error) {
	u := model.User{
		UserName: user.UserName,
		Password: utils.EncryptPassword(user.Password),
	}
	err = db.Debug().Create(&u).Error
	if err != nil {
		zap.L().Error("create user err", zap.Error(err))
		return
	}
	return
}

func UserLogin(user *model.ParamLogin) (AToken, RToken string, err error) {
	var u = new(model.User)
	err = db.Debug().Where("user_name = ?", user.UserName).First(&u).Error
	if err != nil {
		zap.L().Error("find user err", zap.Error(err))
		return "", "", ErrorQuery
	}
	if u.Password != utils.EncryptPassword(user.Password) {
		return "", "", ErrorInvalidPassword
	}
	AToken, RToken, err = jwt.GenToken(u.UserID, u.UserName)
	return
}

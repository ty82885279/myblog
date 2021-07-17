package controllers

type ResCode int64

const (
	//common
	CodeSuccess      ResCode = 200
	CodeServerBusy   ResCode = 500
	CodeInvalidParam ResCode = 3000
	//user
	CodeUserExist       ResCode = 402001
	CodeUserNotExist    ResCode = 402002
	CodeInvalidPassword ResCode = 402003
	CodeSignupSuccess   ResCode = 402004
	CodeLoginSuccess    ResCode = 402005
	//token
	CodeInvalidToken   ResCode = 403001
	CodeTokenNotExist  ResCode = 403002
	CodeNeedLogin      ResCode = 403003
	CodeATokenNotExpir ResCode = 403004
	CodeRTokenHasExpir ResCode = 403005
	//article
	CodeAddArticleSuccess    ResCode = 404001
	CodeEditArticleSuccess   ResCode = 404002
	CodeDeleteArticleSuccess ResCode = 404003
	CodeArticleNotExist      ResCode = 404004
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "请求成功",
	CodeServerBusy:   "服务繁忙",
	CodeInvalidParam: "参数错误",

	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeSignupSuccess:   "注册成功",
	CodeLoginSuccess:    "登录成功",
	CodeInvalidToken:    "无效的Token",
	CodeTokenNotExist:   "缺少Token",
	CodeNeedLogin:       "需要登录",
	CodeATokenNotExpir:  "AToken没有过期",
	CodeRTokenHasExpir:  "RToken已经过期，需重新登录",

	CodeAddArticleSuccess:    "添加文章成功",
	CodeEditArticleSuccess:   "编辑文章成功",
	CodeDeleteArticleSuccess: "删除文章成功",
	CodeArticleNotExist:      "该文章不存在",
}

func (r ResCode) Msg() string {
	msg, ok := codeMsgMap[r]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}

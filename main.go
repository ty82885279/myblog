package main

import (
	"fmt"
	"go.uber.org/zap"
	"myblog/controllers"
	"myblog/dao/mysql"
	"myblog/dao/redis"
	"myblog/logger"
	"myblog/routers"
	"myblog/settings"
)

func main() {
	//1. 加载配置
	err := settings.Init()
	if err != nil {
		panic("Init settings failed" + err.Error())
		return
		zap.L().Debug("Init settings failed", zap.Error(err))
		fmt.Printf("错误:%v\n", err)
		return
	}
	//2. 初始化日志
	err = logger.Init(settings.Cfg.LoggerConf, settings.Cfg.Mode)
	if err != nil {
		panic("init log err" + err.Error())
		return
	}
	defer zap.L().Sync() //将缓冲区的日志追加到文件中

	//3.初始化数据库
	err = mysql.Init(settings.Cfg.MysqlConf)
	if err != nil {
		zap.L().Error("init mysql err", zap.Error(err))
		return
	}
	defer mysql.Close()

	// 4.初始化redis
	err = redis.Init(settings.Cfg.RedisConf)
	if err != nil {
		zap.L().Debug("Init redis failed", zap.Error(err))
		return
	}
	defer redis.Close()
	err = controllers.InitTrans("zh")
	if err != nil {
		zap.L().Error("Init validator tran failed", zap.Error(err))
		return
	}
	r := routers.Setup()
	r.Run(fmt.Sprintf(":%d", settings.Cfg.Port))
}

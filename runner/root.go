package main

import (
	"os"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/miacio/varietas/log"
	"github.com/miacio/vishanti/lib"
	"github.com/spf13/viper"
)

// 初始化日志
func initLog() {
	lp := log.LoggerParam{
		Path:       "./log",
		MaxSize:    256,
		MaxBackups: 10,
		MaxAge:     7,
		Compress:   false,
	}
	lib.Log = lp.Default()
}

// 初始化翻译器
func initTrans() {
	lib.Trans = lib.InitValidateTrans(binding.Validator.Engine().(*validator.Validate))
}

// 初始化配置文件
func initConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	runPath, _ := os.Getwd()
	v.AddConfigPath(runPath)

	if err := v.ReadInConfig(); err != nil {
		lib.Log.Errorf("读取配置文件失败: %v", err)
		return
	}
	// 邮箱服务
	if err := v.UnmarshalKey("email", &lib.EmailCfg); err != nil {
		lib.Log.Errorf("邮件配置读取失败: %v", err)
	} else {
		lib.EmailCfg.IsStatus = true
	}
	// Redis服务
	if err := v.UnmarshalKey("redis", &lib.RedisCfg); err != nil {
		lib.Log.Errorf("redis配置读取失败: %v", err)
	} else {
		// 获取redis client
		lib.RedisClient = lib.RedisCfg.NewClient()
	}

	if err := v.UnmarshalKey("mysql", &lib.DBCfg); err != nil {
		lib.Log.Errorf("数据库配置读取失败: %v", err)
		os.Exit(0)
	} else {
		if err := lib.DBCfg.Connect(); err != nil {
			lib.Log.Errorf("数据库连接失败: %v", err)
			os.Exit(0)
		}
	}
}

func start() {
	initLog()
	initTrans()
	initConfig()
}

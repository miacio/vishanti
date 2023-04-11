package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/miacio/varietas/log"
	"github.com/miacio/varietas/web"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/routers"
	"github.com/spf13/viper"
)

func init() {
	// 日志
	lp := log.LoggerParam{
		Path:       "./log",
		MaxSize:    256,
		MaxBackups: 10,
		MaxAge:     7,
		Compress:   false,
	}
	lib.Log = lp.Default()

	lib.Trans = lib.InitValidateTrans(binding.Validator.Engine().(*validator.Validate))

	// 配置文件
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	runPath, _ := os.Getwd()
	v.AddConfigPath(runPath)

	if err := v.ReadInConfig(); err != nil {
		lib.Log.Errorf("读取配置文件失败: %v", err)
		return
	}
	if err := v.UnmarshalKey("email", &lib.EmailCfg); err != nil {
		lib.Log.Errorf("邮件配置读取失败: %v", err)
	} else {
		lib.EmailCfg.IsStatus = true
	}
}

func main() {
	g := web.New(gin.Default())

	// set static folder
	g.Static("/js", "../page/js")
	g.Static("/images", "../page/images")

	// load html folders
	g.LoadHTMLFolders([]string{"../page"}, ".html")

	routers.Register(g)

	g.Use(web.Limit(64))
	g.Run(":8080")
}

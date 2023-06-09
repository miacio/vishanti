package routers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miacio/varietas/web"
	"github.com/miacio/vishanti/lib"
)

// Register 整个路由服务的注册主线
func Register(e *web.Engine) {
	e.Register(routers...)
	e.Prepare()
}

// 所有路由节点注册均写在这里,便于管理所有路由节点
var (
	root             web.Router = (*rootRouters)(nil)            // root 根节点路由
	systemDictionary web.Router = (*systemDictionaryRouter)(nil) // systemDictionary 系统字典节点路由
	systemFile       web.Router = (*systemFileRouter)(nil)       // systemFile 系统文件路由
	email            web.Router = (*emailRouter)(nil)            // email 邮箱节点路由
	user             web.Router = (*userRouter)(nil)             // user 用户节点路由
	circles          web.Router = (*circlesRouter)(nil)          // circles 圈子节点路由

	routers = []web.Router{root, systemDictionary, systemFile, email, user, circles}
)

// rootRouters 根节点路由实现 ↓↓↓
type rootRouters struct{}

type (
	// errorRequest 错误消息请求体
	errorRequest struct {
		Code string `json:"code" form:"code" uri:"code" binding:"required"`
		Msg  string `json:"msg" form:"msg" uri:"msg" binding:"required"`
		Err  string `json:"err" form:"err" uri:"err"`
	}
)

// 根节点路由注册
func (*rootRouters) Execute(e *gin.Engine) {
	pong := func(ctx *gin.Context) {
		lib.ServerResult(ctx, 200, "pong", nil, nil)
	}

	errorHandler := func(ctx *gin.Context) {
		req := errorRequest{}

		code := 500
		errMsg := ""
		if err := ctx.ShouldBind(&req); err != nil {
			errMsg = lib.TransError(err)
		} else {
			errMsg = strings.Join([]string{req.Msg, req.Err}, " <br /> ^1000")
		}

		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"title": code,
			"err":   errMsg,
		})
	}

	e.GET("/", pong)
	e.GET("/ping", pong)
	// error
	e.GET("/error", errorHandler)
	e.POST("/error", errorHandler)

	e.GET("/test/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.HTML(http.StatusOK, name, nil)
	})
}

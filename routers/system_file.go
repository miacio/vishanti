package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/logic"
	"github.com/miacio/vishanti/middlewares"
)

// 文件路由
type systemFileRouter struct{}

func (*systemFileRouter) Execute(e *gin.Engine) {
	fileGroup := e.Group("/file")
	fileGroup.POST("/upload", middlewares.Auth(), logic.SystemFileLogic.Upload) // 文件上传
	fileGroup.GET("/load", logic.SystemFileLogic.Load)                          // 文件读取
}

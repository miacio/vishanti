package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/logic"
)

// emailRouter 邮箱路由
type emailRouter struct{}

func (*emailRouter) Execute(g *gin.Engine) {
	emailGroup := g.Group("/email")
	emailGroup.POST("/sendCheckCode", logic.EmailLogic.SendCheckCode)
}

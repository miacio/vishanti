package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/logic"
	"github.com/miacio/vishanti/middlewares"
)

// circlesRouter 圈子路由
type circlesRouter struct{}

func (*circlesRouter) Execute(e *gin.Engine) {
	circlesGroup := e.Group("/circles", middlewares.Auth())
	circlesGroup.POST("/create", logic.CirclesLogic.Create) // 创建圈子
	circlesGroup.GET("/find", logic.CirclesLogic.Find)      // 查询自己所拥有的圈子
}

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/logic"
	"github.com/miacio/vishanti/middlewares"
)

// 用户路由实现
type userRouter struct{}

func (*userRouter) Execute(e *gin.Engine) {
	user := e.Group("/user")

	user.GET("/token", logic.UserLogic.TokenGet)

	// 邮箱模块
	emailGroup := user.Group("/email")
	{
		emailGroup.POST("/register", logic.UserLogic.EmailRegister)
		emailGroup.POST("/login", logic.UserLogic.EmailLogin)
		emailGroup.POST("/loginPwd", logic.UserLogic.EmailLoginPwd)
	}

	// 用户信息模块
	detailedGroup := user.Group("/detailed", middlewares.Auth())
	{
		detailedGroup.POST("/update", logic.UserLogic.UpdateDetailed)
	}

}

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/logic"
)

// 用户路由实现
type userRouter struct{}

func (*userRouter) Execute(e *gin.Engine) {
	user := e.Group("/user")

	user.GET("/email/register", logic.UserLogic.EmailRegister)
}

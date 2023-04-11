package routers

import (
	"github.com/gin-gonic/gin"
)

// 用户路由实现
type userRouter struct{}

// emailRegisterRequest 邮箱注册请求体
type emailRegisterRequest struct {
}

func (*userRouter) Execute(e *gin.Engine) {
	user := e.Group("/user")

	user.GET("/email/register", func(ctx *gin.Context) {

	})
}

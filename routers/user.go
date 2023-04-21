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
		emailGroup.POST("/register", logic.UserLogic.EmailRegister)   // 用户邮箱注册
		emailGroup.POST("/login", logic.UserLogic.EmailLogin)         // 用户邮箱登录
		emailGroup.POST("/loginPwd", logic.UserLogic.EmailLoginPwd)   // 用户邮箱密码方式登录
		emailGroup.POST("/updatePwd", logic.UserLogic.EmailUpdatePwd) // 用户邮箱修改密码
	}

	// 用户信息模块
	detailedGroup := user.Group("/detailed", middlewares.Auth())
	{
		detailedGroup.POST("/update", logic.UserLogic.UpdateDetailed)       // 修改用户信息接口
		detailedGroup.POST("/updateHeadPic", logic.UserLogic.UpdateHeadPic) // 修改用户头像接口
	}

}

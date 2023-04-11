package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/lib"
)

// emailRouter 邮箱路由
type emailRouter struct {
}

type emailSendRequest struct {
	Email     string `form:"email" uri:"email" json:"email" binding:"required,email"` // 收件人地址
	EmailType string `form:"emailType" uri:"emailType" json:"emailType" binding:"required"`
}

func (*emailRouter) Execute(g *gin.Engine) {
	emailGroup := g.Group("/email")
	emailGroup.POST("/checkCode", func(ctx *gin.Context) {
		req := emailSendRequest{}
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数错误", "err": lib.TransError(err)})
			return
		}
		if err := lib.EmailCfg.Send(req.Email, "这是一次测试", "测试成功了"); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器异常", "err": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "发送成功"})
	})
}

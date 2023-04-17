package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/logic"
	"github.com/miacio/vishanti/middlewares"
)

// 字典路由
type systemDictionaryRouter struct{}

func (*systemDictionaryRouter) Execute(e *gin.Engine) {
	dictionaryGroup := e.Group("/dict")
	dictionaryGroup.GET("/findByGroup", logic.SystemDictionaryLogic.FindByGroup)
	dictionaryGroup.POST("/inserts", middlewares.Auth(), logic.SystemDictionaryLogic.Inserts)
}

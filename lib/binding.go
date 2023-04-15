package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 参数绑定封装,仅只需要判断为false则return即可
func ShouldBind(ctx *gin.Context, obj any) bool {
	if err := ctx.ShouldBind(obj); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误", "err": TransError(err)})
		return false
	}
	return true
}

// 参数绑定封装,仅只需要判断为false则return即可
func ShouldBindJSON(ctx *gin.Context, obj any) bool {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误", "err": TransError(err)})
		return false
	}
	return true
}

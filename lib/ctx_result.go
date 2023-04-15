package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 服务器错误,判断异常是否为空,为空返回true,非空返回false并向客户端发送服务器错误信息
func ServerFail(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": "服务器异常", "err": err.Error()})
		return false
	}
	return true
}

// 服务器错误,判断异常是否为空,为空返回true,非空返回false并向客户端发送服务器错误信息
func ServerFailf(ctx *gin.Context, code int, msg string, err error) bool {
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": msg, "err": err.Error()})
		return false
	}
	return true
}

// 统一响应方法
func ServerResult(ctx *gin.Context, code int, msg string, data any, err error) {
	param := gin.H{
		"code": code,
		"msg":  msg,
	}
	if data != nil {
		param["data"] = data
	}
	if err != nil {
		param["err"] = err.Error()
	}
	ctx.JSON(http.StatusOK, param)
}

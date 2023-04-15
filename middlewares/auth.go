package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/store"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tk := ctx.GetHeader("token")
		if tk == "" {
			lib.ServerFailf(ctx, 500, "非法访问", nil)
			ctx.Abort()
		}
		obj, err := store.UserTokenStore.Get(tk)
		if !lib.ServerFailf(ctx, 500, "token获取失败", err) {
			ctx.Abort()
		}
		ctx.Set("token", obj)
		ctx.Next()
	}
}

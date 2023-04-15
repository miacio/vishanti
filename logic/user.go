package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/store"
)

type userLogic struct{}

type IUserLogic interface {
	TokenGet(ctx *gin.Context)      // 依据token获取登录信息
	EmailRegister(ctx *gin.Context) // EmailRegister 邮箱注册 - 无错误信息时将进行登录操作
}

var UserLogic IUserLogic = (*userLogic)(nil)

// emailRegisterRequest 邮箱注册请求体
type emailRegisterRequest struct {
	Email    string `form:"email" uri:"email" json:"email" binding:"required,email"`                 // 用户邮箱地址
	Code     string `form:"code" uri:"code" json:"code" binding:"required,len=6"`                    // 邮箱验证码
	Uid      string `form:"uid" uri:"uid" json:"uid" binding:"required"`                             // 邮箱验证码的uid
	NickName string `form:"nickName" uri:"nickName" json:"nickName" binding:"required,min=2,max=32"` // 昵称
	Account  string `form:"account" uri:"account" json:"account" binding:"required,min=6,max=32"`    // 账号
	Password string `form:"password" uri:"password" json:"password" binding:"required,min=6,max=32"` // 密码
}

func (*userLogic) TokenGet(ctx *gin.Context) {
	token := ctx.Query("token")

	result, err := store.UserTokenStore.Get(token)
	if !lib.ServerFailf(ctx, 500, "获取失败", err) {
		return
	}

	lib.ServerResult(ctx, 200, "获取成功", result, nil)
}

// EmailRegister 邮箱注册 - 无错误信息时将进行登录操作
func (*userLogic) EmailRegister(ctx *gin.Context) {
	var req emailRegisterRequest
	if !lib.ShouldBind(ctx, &req) {
		return
	}

	err := EmailLogic.CheckCode(req.Email, "register", req.Uid, req.Code)
	if !lib.ServerFailf(ctx, 500, "验证码校验失败", err) {
		return
	}

	userId, err := store.UserStore.EmailRegister(req.Email, req.NickName, req.Account, req.Password)
	if !lib.ServerFailf(ctx, 500, "注册失败", err) {
		return
	}

	tokenKey, err := store.UserTokenStore.LoginSave(userId)
	if !lib.ServerFailf(ctx, 500, "登录失败", err) {
		return
	}

	lib.ServerResult(ctx, 200, "登录成功", tokenKey, nil)
}

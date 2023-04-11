package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/lib"
)

type userLogic struct{}

type IUserLogic interface {
	EmailRegister(ctx *gin.Context) // EmailRegister 邮箱注册
}

var UserLogic IUserLogic = (*userLogic)(nil)

// emailRegisterRequest 邮箱注册请求体
type emailRegisterRequest struct {
	Email    string `form:"email" uri:"email" json:"email" binding:"required,email"`         // 用户邮箱地址
	Code     string `form:"code" uri:"code" json:"code" binding:"required,eq=6"`             // 邮箱验证码
	Uid      string `form:"uid" uri:"uid" json:"uid" binding:"required"`                     // 邮箱验证码的uid
	NickName string `form:"nickName" uri:"nickName" json:"nickName" binding:"required,gt=2"` // 昵称
	Account  string `form:"account" uri:"account" json:"account" binding:"required,gt=6"`    // 账号
	Password string `form:"password" uri:"password" json:"password" binding:"required,gt=6"` // 密码
}

// EmailRegister 邮箱注册
func (*userLogic) EmailRegister(ctx *gin.Context) {
	var req emailRegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数错误", "err": lib.TransError(err)})
		return
	}

	if err := EmailLogic.CheckCode(req.Email, "register", req.Uid, req.Code); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "验证码校验失败", "err": err})
		return
	}

	// lib.DB.
}

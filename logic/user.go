package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
	"github.com/miacio/vishanti/store"
)

type userLogic struct{}

type IUserLogic interface {
	TokenGet(ctx *gin.Context) // 依据token获取登录信息

	EmailRegister(ctx *gin.Context)  // EmailRegister 邮箱注册 - 无错误信息时将进行登录操作
	EmailLogin(ctx *gin.Context)     // EmailLogin 邮箱登录
	EmailLoginPwd(ctx *gin.Context)  // EmailLoginPwd 邮箱登录 密码登录方式
	EmailUpdatePwd(ctx *gin.Context) // EmailUpdatePwd 邮箱修改密码

	UpdateDetailed(ctx *gin.Context) // UpdateDetailed 修改用户信息
	UpdateHeadPic(ctx *gin.Context)  // UpdateHeadPic 修改用户头像
}

var UserLogic IUserLogic = (*userLogic)(nil)

// 依据token获取登录信息
func (*userLogic) TokenGet(ctx *gin.Context) {
	token := ctx.Query("token")

	result, err := store.UserTokenStore.Get(token)
	if !lib.ServerFailf(ctx, 500, "获取失败", err) {
		return
	}

	lib.ServerSuccess(ctx, "获取成功", result)
}

// emailRegisterRequest 邮箱注册请求体
type emailRegisterRequest struct {
	Email    string `form:"email" uri:"email" json:"email" binding:"required,email"`                 // 用户邮箱地址
	Code     string `form:"code" uri:"code" json:"code" binding:"required,len=6"`                    // 邮箱验证码
	Uid      string `form:"uid" uri:"uid" json:"uid" binding:"required"`                             // 邮箱验证码的uid
	NickName string `form:"nickName" uri:"nickName" json:"nickName" binding:"required,min=2,max=32"` // 昵称
	Account  string `form:"account" uri:"account" json:"account" binding:"required,min=6,max=32"`    // 账号
	Password string `form:"password" uri:"password" json:"password" binding:"required,min=6,max=32"` // 密码
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

	lib.ServerSuccess(ctx, "登录成功", tokenKey)
}

// emailLoginRequest 邮箱登录请求体
type emailLoginRequest struct {
	Email string `form:"email" uri:"email" json:"email" binding:"required,email"` // 用户邮箱地址
	Code  string `form:"code" uri:"code" json:"code" binding:"required,len=6"`    // 邮箱验证码
	Uid   string `form:"uid" uri:"uid" json:"uid" binding:"required"`             // 邮箱验证码的uid
}

// EmailLogin 邮箱登录
func (*userLogic) EmailLogin(ctx *gin.Context) {
	var req emailLoginRequest
	if !lib.ShouldBindJSON(ctx, &req) {
		return
	}
	err := EmailLogic.CheckCode(req.Email, "login", req.Uid, req.Code)
	if !lib.ServerFailf(ctx, 500, "验证码校验失败", err) {
		return
	}
	userAccountInfo, err := store.UserStore.FindAccountByEmail(req.Email)
	if !lib.ServerFailf(ctx, 500, "登录失败", err) {
		return
	}
	tokenKey, err := store.UserTokenStore.LoginSave(userAccountInfo.ID)
	if !lib.ServerFailf(ctx, 500, "登录失败", err) {
		return
	}
	lib.ServerSuccess(ctx, "登录成功", tokenKey)
}

// emailLoginPwdRequest 邮箱密码登录请求体
type emailLoginPwdRequest struct {
	Email    string `form:"email" uri:"email" json:"email" binding:"required,email"`                 // 用户邮箱地址
	Password string `form:"password" uri:"password" json:"password" binding:"required,min=6,max=32"` // 密码
}

// EmailLoginPwd 邮箱登录 密码登录方式
func (*userLogic) EmailLoginPwd(ctx *gin.Context) {
	var req emailLoginPwdRequest
	if !lib.ShouldBind(ctx, &req) {
		return
	}
	userAccountInfo, err := store.UserStore.FindAccountByEmailAndPwd(req.Email, req.Password)
	if !lib.ServerFailf(ctx, 500, "登录失败", err) {
		return
	}
	tokenKey, err := store.UserTokenStore.LoginSave(userAccountInfo.ID)
	if !lib.ServerFailf(ctx, 500, "登录失败", err) {
		return
	}
	lib.ServerSuccess(ctx, "登录成功", tokenKey)
}

// emailUpdatePwdRequest 邮箱修改密码请求体
type emailUpdatePwdRequest struct {
	Email    string `form:"email" uri:"email" json:"email" binding:"required,email"`                 // 用户邮箱地址
	Code     string `form:"code" uri:"code" json:"code" binding:"required,len=6"`                    // 邮箱验证码
	Uid      string `form:"uid" uri:"uid" json:"uid" binding:"required"`                             // 邮箱验证码的uid
	Password string `form:"password" uri:"password" json:"password" binding:"required,min=6,max=32"` // 修改后的密码
}

// EmailUpdatePwd 邮箱修改密码
func (*userLogic) EmailUpdatePwd(ctx *gin.Context) {
	var req emailUpdatePwdRequest
	if !lib.ShouldBind(ctx, &req) {
		return
	}

	err := EmailLogic.CheckCode(req.Email, "update", req.Uid, req.Code)
	if !lib.ServerFailf(ctx, 500, "验证码校验失败", err) {
		return
	}

	userAccountInfo, err := store.UserStore.FindAccountByEmail(req.Email)
	if !lib.ServerFail(ctx, err) {
		return
	}

	if !lib.ServerFail(ctx, store.UserStore.UpdatePasswordById(userAccountInfo.ID, req.Password)) {
		return
	}
	lib.ServerSuccess(ctx, "修改成功", nil)
}

// updateUserDetailRequest 修改用户信息请求结构体
type updateUserDetailRequest struct {
	NickName      string `form:"nickName" json:"nickName" uri:"nickName"`                // nickName
	Sex           string `form:"sex" json:"sex" uri:"sex"`                               // sex
	BirthdayYear  int    `form:"birthdayYear" json:"birthdayYear" uri:"birthdayYear"`    // birthdayYear
	BirthdayMonth int    `form:"birthdayMonth" json:"birthdayMonth" uri:"birthdayMonth"` // birthdayMonth
	BirthdayDay   int    `form:"birthdayDay" json:"birthdayDay" uri:"birthdayDay"`       // birthdayDay
}

func (u *updateUserDetailRequest) ToModel(id, accountId string) model.UserDetailedInfo {
	return model.UserDetailedInfo{
		ID:            id,
		UserAccountID: accountId,
		NickName:      u.NickName,
		Sex:           u.Sex,
		BirthdayYear:  u.BirthdayYear,
		BirthdayMonth: u.BirthdayMonth,
		BirthdayDay:   u.BirthdayDay,
	}
}

// UpdateDetailed 修改用户信息
func (*userLogic) UpdateDetailed(ctx *gin.Context) {
	var req updateUserDetailRequest
	if !lib.ShouldBindJSON(ctx, &req) {
		return
	}

	mo, ok := store.TokenGet(ctx)
	if !ok {
		return
	}

	if err := store.UserStore.UpdateDetailed(req.ToModel(mo.DetailedInfo.ID, mo.AccountInfo.ID)); err != nil {
		lib.ServerResult(ctx, 500, "修改用户信息失败", nil, err)
		return
	}
	if !lib.ServerFail(ctx, store.TokenFlush(ctx)) {
		return
	}
	lib.ServerSuccess(ctx, "修改成功", nil)
}

// UpdateHeadPic 修改用户头像
func (*userLogic) UpdateHeadPic(ctx *gin.Context) {
	var req fileDefaultUploadRequest
	if !lib.ShouldBind(ctx, &req) {
		return
	}

	mo, ok := store.TokenGet(ctx)
	if !ok {
		return
	}

	systemFileInfoModel, err := SystemFileLogic.UploadLogic(mo.AccountInfo.ID, "USER_HEAD_PIC", req, ".jpg", ".jpeg", ".png", ".svg", ".webp")
	if !lib.ServerFail(ctx, err) {
		return
	}

	if err := store.UserStore.UpdateUserHeadPic(mo.AccountInfo.ID, systemFileInfoModel.ID); err != nil {
		lib.ServerFail(ctx, err)
		return
	}

	if !lib.ServerFail(ctx, store.TokenFlush(ctx)) {
		return
	}
	lib.ServerSuccess(ctx, "修改成功", systemFileInfoModel.ID)
}

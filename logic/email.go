package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/store"
)

type emailLogic struct{}

type IEmailLogic interface {
	SendCheckCode(ctx *gin.Context)                     // SendCheckCode 发送验证码
	CheckCode(email, emailType, uid, code string) error // CheckCode 校验验证码
}

var EmailLogic IEmailLogic = (*emailLogic)(nil)

type emailSendRequest struct {
	Email     string `form:"email" uri:"email" json:"email" binding:"required,email"` // 收件人地址
	EmailType string `form:"emailType" uri:"emailType" json:"emailType" binding:"required"`
}

// SendCheckCode 发送验证码
func (e *emailLogic) SendCheckCode(ctx *gin.Context) {
	req := emailSendRequest{}
	if !lib.ShouldBind(ctx, &req) {
		return
	}

	switch req.EmailType {
	case "register":
		ok, err := store.UserStore.EmailRepeat(req.Email)
		if !lib.ServerFail(ctx, err) {
			return
		}
		if ok {
			lib.ServerResult(ctx, 400, "当前邮箱已经被注册,请勿重复注册", nil, nil)
			return
		}
		uid, err := e.emailRegister(req.Email)
		if !lib.ServerFail(ctx, err) {
			return
		}
		lib.ServerSuccess(ctx, "发送成功", uid)
	case "login":
		ok, err := store.UserStore.EmailRepeat(req.Email)
		if !lib.ServerFail(ctx, err) {
			return
		}
		if !ok {
			lib.ServerResult(ctx, 400, "当前邮箱未注册", nil, nil)
			return
		}
		uid, err := e.emailLogin(req.Email)
		if !lib.ServerFail(ctx, err) {
			return
		}
		lib.ServerSuccess(ctx, "发送成功", uid)
	case "update":
		uid, err := e.emailRegister(req.Email)
		if !lib.ServerFail(ctx, err) {
			return
		}
		lib.ServerSuccess(ctx, "发送成功", uid)
	default:
		lib.ServerResult(ctx, 400, "参数错误", nil, errors.New("未知的邮件推送请求"))
	}
}

// CheckCode 校验验证码
func (e *emailLogic) CheckCode(email, emailType, uid, code string) error {
	rc := context.Background()
	var key string
	switch emailType {
	case "register":
		key = "EMAIL:REGISTER:" + email + ":" + uid
	case "login":
		key = "EMAIL:LOGIN:" + email + ":" + uid
	case "update":
		key = "EMAIL:UPDATE:" + email + ":" + uid
	}
	if key == "" {
		return errors.New("未知的校验方式")
	}

	rcode, err := lib.RedisClient.Get(rc, key).Result()
	if err != nil {
		return err
	}
	if rcode != code {
		return errors.New("验证码错误")
	}
	lib.RedisClient.Del(rc, key)
	return nil
}

func (*emailLogic) emailRegister(email string) (string, error) {
	var msg strings.Builder
	msg.WriteString("vishanti:\n")
	msg.WriteString("    您好!感谢您注册vishanti平台,您的验证码是: %s\n")
	msg.WriteString("    当前验证码的有效时间是: 30分钟\n")
	msg.WriteString("如果非您本人操作,请删除该邮件并不要将验证码告诉给他人!")
	m := msg.String()
	code := lib.RandCheckCode(6)
	m = fmt.Sprintf(m, code)

	uid := lib.UID()

	if err := lib.EmailCfg.Send(email, "欢迎您注册vishanti平台", m); err != nil {
		return "", err
	}

	rc := context.Background()
	if err := lib.RedisClient.SetEx(rc, "EMAIL:REGISTER:"+email+":"+uid, code, time.Minute*30).Err(); err != nil {
		return "", err
	}
	return uid, nil
}

func (*emailLogic) emailLogin(email string) (string, error) {
	var msg strings.Builder
	msg.WriteString("vishanti:\n")
	msg.WriteString("    您好!您当前正在使用邮箱验证码登录,您的验证码是: %s\n")
	msg.WriteString("    当前验证码的有效时间是: 30分钟\n")
	msg.WriteString("如果非您本人操作,请删除该邮件并不要将验证码告诉给他人!")
	m := msg.String()
	code := lib.RandCheckCode(6)
	m = fmt.Sprintf(m, code)

	uid := lib.UID()

	if err := lib.EmailCfg.Send(email, "欢迎您使用vishanti平台", m); err != nil {
		return "", err
	}

	rc := context.Background()
	if err := lib.RedisClient.SetEx(rc, "EMAIL:LOGIN:"+email+":"+uid, code, time.Minute*30).Err(); err != nil {
		return "", err
	}
	return uid, nil
}

func (*emailLogic) emailUpdate(email string) (string, error) {
	var msg strings.Builder
	msg.WriteString("vishanti:\n")
	msg.WriteString("    您好!您当前正在申请修改用户信息或密码,您的验证码是: %s\n")
	msg.WriteString("    当前验证码的有效时间是: 30分钟\n")
	msg.WriteString("如果非您本人操作,请删除该邮件并不要将验证码告诉给他人!")
	m := msg.String()
	code := lib.RandCheckCode(6)
	m = fmt.Sprintf(m, code)

	uid := lib.UID()

	if err := lib.EmailCfg.Send(email, "欢迎您使用vishanti平台", m); err != nil {
		return "", err
	}

	rc := context.Background()
	if err := lib.RedisClient.SetEx(rc, "EMAIL:UPDATE:"+email+":"+uid, code, time.Minute*30).Err(); err != nil {
		return "", err
	}
	return uid, nil
}

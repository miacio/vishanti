package logic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/lib"
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
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数错误", "err": lib.TransError(err)})
		return
	} else {
		switch req.EmailType {
		case "register":
			if code, err := e.emailRegister(req.Email); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器异常", "err": err})
				return
			} else {
				ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "发送成功", "data": code})
				return
			}
		case "update":
			if code, err := e.emailRegister(req.Email); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器异常", "err": err})
				return
			} else {
				ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "发送成功", "data": code})
				return
			}
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数错误", "err": "未知的邮件推送请求"})
		}
	}
}

// CheckCode 校验验证码
func (e *emailLogic) CheckCode(email, emailType, uid, code string) error {
	rc := context.Background()
	var key string
	switch emailType {
	case "register":
		key = "EMAIL:REGISTER:" + email + ":" + uid
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

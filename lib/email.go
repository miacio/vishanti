package lib

import (
	"fmt"
	"net/smtp"

	"github.com/miacio/varietas/email"
)

// 邮箱配置
type EmailCfgParam struct {
	Name     string `toml:"name" xml:"name" json:"name"`             // 发送者名称
	MailAddr string `toml:"mailAddr" xml:"mailAddr" json:"mailAddr"` // 发送者邮箱地址 youemail@email.com
	SmtpAddr string `toml:"smtpAddr" xml:"smtpAddr" json:"smtpAddr"` // 邮箱服务器地址(包含端口号)非SSL协议: 例如网易的 smtp.163.com:25
	HostAddr string `toml:"hostAddr" xml:"hostAddr" json:"hostAddr"` // 登录授权时服务器地址,不包含端口号: 例如网易的 smtp.163.com
	Password string `toml:"password" xml:"password" json:"password"` // 发送者邮箱授权密钥
	IsStatus bool   // 是否启用
}

// 发送邮件方法 (正文方式)
func (ep *EmailCfgParam) Send(to string, title, message string) error {
	if ep.IsStatus {
		e := email.New()
		e.From = fmt.Sprintf("%s <%s>", ep.Name, ep.MailAddr)
		e.To = []string{to}
		e.Subject = title
		e.Text = []byte(message)
		return e.Send(ep.SmtpAddr, smtp.PlainAuth("", ep.MailAddr, ep.Password, ep.HostAddr))
	}
	return fmt.Errorf("未开通邮件服务,请配置后重试")
}

// 发送邮件方法 (网页方式)
func (ep *EmailCfgParam) SendHtml(to string, title, html string) error {
	if ep.IsStatus {
		e := email.New()
		e.From = fmt.Sprintf("%s <%s>", ep.Name, ep.MailAddr)
		e.To = []string{to}
		e.Subject = title
		e.Html = []byte(html)
		return e.Send(ep.SmtpAddr, smtp.PlainAuth("", ep.MailAddr, ep.Password, ep.HostAddr))
	}
	return fmt.Errorf("未开通邮件服务,请配置后重试")
}

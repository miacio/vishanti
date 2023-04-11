package lib

import (
	"net/smtp"
	"testing"

	"github.com/miacio/varietas/email"
)

func TestSendEmail(t *testing.T) {
	e := email.New()
	e.From = "miajio <miajio@163.com>"
	e.To = []string{"1877378299@qq.com"}
	e.Subject = "test标题"
	e.Text = []byte("这是正文内容")
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "miajio@163.com", "IDLEHXQGHMBYBONM", "smtp.163.com"))
	if err != nil {
		t.Fatalf("send email: %v", err)
	}
}

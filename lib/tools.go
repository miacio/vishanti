package lib

import (
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

const (
	bs = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandString randomly generate a specified length string based on the given string
// 依据给定字符串生成给定长度的随机字符串
func RandString(base string, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = base[rand.Intn(len(base))]
	}
	return string(b)
}

// 随机验证码
func RandCheckCode(n int) string {
	return RandString(bs, n)
}

// 获取UUID (转大写,不带-)
func UID() string {
	return strings.ToUpper(strings.ReplaceAll(uuid.New().String(), "-", ""))
}

package lib

import (
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Trans ut.Translator
)

func InitValidateTrans(validate *validator.Validate) ut.Translator {
	uni := ut.New(zh.New())

	trans, _ := uni.GetTranslator("zh")
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		Log.Errorf("翻译器处理失败")
	}
	return trans
}

func TransError(err error) string {
	if ve, ok := err.(validator.ValidationErrors); !ok {
		return err.Error()
	} else {
		var res strings.Builder
		for _, e := range ve {
			res.WriteString(e.Translate(Trans) + ";")
		}
		return res.String()
	}
}

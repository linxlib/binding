package binding

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func init() {
	chinese := zh.New()
	uni = ut.New(chinese, chinese)
	trans, _ = uni.GetTranslator("zh")
	validate := Validator.Engine().(*validator.Validate)
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
}

func HandleValidationErrors(err error) ([]string, bool) {
	var fields []string
	if v, ok := err.(validator.ValidationErrors); ok {
		for _, err := range v {
			tmp := err.Translate(trans)
			fields = append(fields, tmp)
		}
		return fields, true
	} else {
		return fields, false
	}
}

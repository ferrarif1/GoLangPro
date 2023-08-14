package util

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	ValidateUtil = validateUtil{}
	uni          *ut.UniversalTranslator
	validate     *validator.Validate
	trans        ut.Translator
)

// gin验证器翻译
type validateUtil struct{}

// Init 初始化
func (vu validateUtil) Init() {
	zh := zh.New()
	uni = ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

// Translate 翻译错误信息
func (vu validateUtil) Translate(err error) (result string) {
	result = ""
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		// 按顺序提示首个
		result = err.Translate(trans)
		break
	}
	return
}

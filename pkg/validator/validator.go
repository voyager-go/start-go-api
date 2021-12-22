package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

// NewValidate 构造验证器
func NewValidate() {
	// 注册翻译器
	zhTrans := zh.New()
	uni = ut.New(zhTrans, zhTrans)
	trans, _ = uni.GetTranslator("zh")

	// 获取gin的验证器
	validate = binding.Validator.Engine().(*validator.Validate)
	// 注册翻译器
	zhtranslations.RegisterDefaultTranslations(validate, trans)
}

// Translate 翻译错误信息
func Translate(err error) []string {
	var result []string
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result = append(result, err.Translate(trans))
	}
	return result
}

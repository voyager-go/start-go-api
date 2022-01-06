package validator

import (
	"encoding/json"
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
func Translate(errs error) []string {
	var result []string
	switch errs.(type) {
	case validator.ValidationErrors:
		errorSlice := errs.(validator.ValidationErrors)
		for _, err := range errorSlice {
			result = append(result, err.Translate(trans))
		}
		break
	case *json.UnmarshalTypeError:
		result = append(result, "参数类型解析失败")
		break
	default:
		result = append(result, "参数错误")
	}

	return result
}

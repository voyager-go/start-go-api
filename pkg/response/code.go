package response

type Code int

// 定义通用错误码，因为使用了iota，为了避免常量混乱，请新增枚举值时往下追加而不是在中间插入
const (
	Success Code = 0
	Failed  Code = 10000 + iota
	UnAuthed
	InternalErr
	RequestMethodErr
	RequestParamErr
	RequestFormErr
	AccountInfoErr
)

// CodeMap 错误码说明
var CodeMap = map[Code]string{
	Success:          "请求成功",
	Failed:           "请求失败",
	UnAuthed:         "未认证",
	InternalErr:      "服务器内部错误",
	RequestMethodErr: "请求方式错误",
	RequestParamErr:  "请求参数错误",
	RequestFormErr:   "请求表单错误",
	AccountInfoErr:   "账号信息有误",
}

// Msg 返回错误码对应的说明
func (c Code) Msg() string {
	if v, ok := CodeMap[c]; ok {
		return v
	}
	return ``
}

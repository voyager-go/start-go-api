package enum

type (
	MethodType int8
	IsUseType  int8
)

const (
	SysApiMethodPost   = iota + 1 // 创建POST
	SysApiMethodGet               // 查看GET
	SysApiMethodPut               // 更新PUT
	SysApiMethodDelete            // 删除DELETE
)

const (
	IsUseForbidden  = iota // 禁用
	IsUseTrueNormal        // 启用
)

func (m MethodType) String() string {
	switch m {
	case SysApiMethodPost:
		return "POST"
	case SysApiMethodGet:
		return "GET"
	case SysApiMethodPut:
		return "PUT"
	case SysApiMethodDelete:
		return "DELETE"
	default:
		return "UnKnown"
	}
}

func (t IsUseType) String() string {
	switch t {
	case IsUseForbidden:
		return "禁用"
	case IsUseTrueNormal:
		return "启用"
	default:
		return "UnKnown"
	}
}

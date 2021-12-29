package task_util

import (
	"github.com/google/uuid"
	"strings"
)

// JoinUUIDPrefix 给UUID加上自定义前缀
func JoinUUIDPrefix(prefix string) string {
	return prefix + GetPureUUID()
}

// GetPureUUID 把UUID中的-去掉
func GetPureUUID() string {
	uid, _ := uuid.NewUUID()
	return strings.Replace(uid.String(), "-", "", -1)
}

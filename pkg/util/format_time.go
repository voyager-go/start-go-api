package util

import (
	"fmt"
	"time"
)

type FormatTime time.Time

// MarshalJSON 实现时间的json序列化方法
func (t FormatTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

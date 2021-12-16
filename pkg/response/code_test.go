package response

import "testing"

func TestCode_Msg(t *testing.T) {
	tests := []struct {
		name string
		c    Code
		want string
	}{
		{name: "发送成功请求", c: Success, want: "请求成功"},
		{name: "发送失败请求", c: Failed, want: "请求失败"},
		{name: "用户未认证", c: UnAuthed, want: "未认证"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Msg(); got != tt.want {
				t.Errorf("Msg() = %v, code值为: %v, want值为: %v", got, tt.c, tt.want)
			}
		})
	}
}

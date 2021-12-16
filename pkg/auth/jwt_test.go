package auth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type User struct {
	ID       int64
	Nickname string
	Phone    string
}

func TestGenerateJwtToken(t *testing.T) {
	var user = User{
		ID:       1001,
		Nickname: "张文杰",
		Phone:    "15105191181",
	}
	token, err := GenerateJwtToken("123456", time.Now().Unix()+60, user, "start - go - api")
	assert.Nil(t, err)
	fmt.Println(token)
}

func TestParseJwtToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7IklEIjoxMDAxLCJOaWNrbmFtZSI6IuW8oOaWh-adsCIsIlBob25lIjoiMTUxMDUxOTExODEifSwiZXhwIjoxNjM5NjQ1MjUwLCJpc3MiOiJzdGFydCAtIGdvIC0gYXBpIn0.oWxbVDWMZxRn18y0sDat1BvS-iPdwTVqmyyj2Be1VtM"
	payload, err := ParseJwtToken(token, "123456")
	assert.Nil(t, err)
	fmt.Println(payload)
}

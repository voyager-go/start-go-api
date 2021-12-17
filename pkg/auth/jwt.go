package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JwtPayload 载荷
type JwtPayload struct {
	User interface{} `json:"user"`
	jwt.StandardClaims
}

// GenerateJwtToken 生成token
func GenerateJwtToken(secret string, expire int64, user interface{}, issuer string) (string, error) {
	data := JwtPayload{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expire,
			Issuer:    issuer,
		},
	}
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	if token, err := j.SignedString([]byte(secret)); err != nil {
		return "", errors.New("jwt 生成token失败: " + err.Error())
	} else {
		return token, nil
	}
}

// ParseJwtToken 解析jwt token
func ParseJwtToken(jwtToken string, secret string) (*JwtPayload, error) {
	if jwtToken == "" {
		return nil, errors.New("jwt token 为空")
	}
	token, err := jwt.ParseWithClaims(jwtToken, &JwtPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("jwt 解析失败:" + err.Error())
	}
	if claims, ok := token.Claims.(*JwtPayload); ok && token.Valid {
		return claims, nil
	} else {
		return claims, errors.New("jwt 解析后验证失败")
	}
}

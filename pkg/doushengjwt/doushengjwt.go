package doushengjwt

import (
	"dousheng/config"
	"dousheng/pkg/errdeal"
	"github.com/dgrijalva/jwt-go"
	"sync"
	"time"
)

var (
	tokenExpireDuration time.Duration
	// 秘钥
	secret = []byte("dousheng")
	// 签发人
	issuer = "dousheng.com"
	once   sync.Once
)

type DouShengClaims struct {
	UserName string `json:"username"`
	UserID   int64  `json:"userid"`
	jwt.StandardClaims
}

// GenerateToken 生成用户鉴权
// 用户名 和 用户id
func GenerateToken(username string, userid int64) (token string, err error) {
	// token过期时间 24小时 * 天
	dsc := DouShengClaims{
		UserName: username,
		UserID:   userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiresAt(), // 过期时间
			Issuer:    issuer,           // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, dsc)
	token, err = tokenClaims.SignedString(secret)
	return
}

// ParseToken 解析JWT
func ParseToken(token string) (*DouShengClaims, error) {

	dsc := new(DouShengClaims) // 存放解析出来的数据

	tokenClaims, err := jwt.ParseWithClaims(token, dsc, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !tokenClaims.Valid {
		return nil, errdeal.CodeInvalidTokenErr
	}

	return dsc, nil
}

// 获取过期时间
func tokenExpiresAt() int64 {
	once.Do(func() {
		tokenExpireDuration = 24 * time.Hour * time.Duration(config.ConfInstance().DurationConfig.Token)
	})
	return time.Now().Add(tokenExpireDuration).Unix()
}

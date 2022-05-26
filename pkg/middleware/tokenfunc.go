package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// ContextTokenFunc 获取token的方法 可自己实现
type ContextTokenFunc interface {
	QueryToken(*gin.Context) string
}

// 从url中获取token
type urlToken struct {
	urlParam string
}

func (u *urlToken) QueryToken(c *gin.Context) string {
	return c.Query(u.urlParam)
}

// URLToken 从URL中获取token
func URLToken(urlParam string) ContextTokenFunc {
	return &urlToken{
		urlParam: urlParam,
	}
}

type formToken struct {
	formParam string // form-data的标识名
	TokenType *struct {
		Token string `json:"token"`
	}
}

func (f *formToken) QueryToken(c *gin.Context) string {
	req := c.PostForm(f.formParam)
	if err := json.Unmarshal([]byte(req), f.TokenType); err != nil {
		return ""
	}
	return f.TokenType.Token
}

// FormToken 从form-data中获取token
func FormToken(formParam string) ContextTokenFunc {
	return &formToken{
		formParam: formParam,
	}
}

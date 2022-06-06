package middlewares

import (
	"github.com/gin-gonic/gin"
)

// ContextTokenFunc 获取token的接口 可自己实现
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
	token := c.PostForm(f.formParam)
	// 反序列化到结构体中 即便出错 token默认值为 “”
	//_ = json.Unmarshal([]byte(req), f.TokenType)
	//return f.TokenType.Token
	return token
}

// FormToken 从form-data中获取token
func FormToken(formParam string) ContextTokenFunc {
	return &formToken{
		formParam: formParam,
		TokenType: &struct {
			Token string `json:"token"`
		}{Token: ""},
	}
}

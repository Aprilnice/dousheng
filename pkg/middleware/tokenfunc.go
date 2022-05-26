package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// ContextTokenFunc 获取token的方法 可自己实现
type ContextTokenFunc func(*gin.Context) string

// URLToken 从URL中获取token
func URLToken() ContextTokenFunc {
	return func(c *gin.Context) string {
		return c.Query("token")
	}
}

// FormToken 从form-data中获取token
func FormToken(formKey string) ContextTokenFunc {
	type tmpToken struct {
		Token string `json:"token"`
	}
	return func(c *gin.Context) string {
		req := c.PostForm(formKey)
		var tokenRes tmpToken
		if err := json.Unmarshal([]byte(req), &tokenRes); err != nil {
			return ""
		}
		return tokenRes.Token
	}

}

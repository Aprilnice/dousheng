package middlewares

import (
	jwt "dousheng/pkg/doushengjwt"
	"dousheng/pkg/errdeal"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ContextUserID = "ContextUserID" // 将解析出来的用户名存放在Context中 方便后续以在上下文中获取到
)

// JwtTokenMiddleware token验证的中间件
func JwtTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求中的token
		token := c.Query("token")
		//token := c.Request.Header.Get("Authorization")
		if len(token) == 0 { // 没有token
			c.JSON(http.StatusOK, errdeal.NewResponse(errdeal.CodeWithoutTokenErr))
			c.Abort()
			return
		}

		// 标准的Token格式
		// Bearer xxx.xxx.xxx
		// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IuW8oOS4iSIsImV4cCI6MTY0ODUzMTI0OSwiaXNzIjoiYmx1ZWJlbGwuY29tIn0.15XVaBOhwKqrjhQpI1RKAVL0vqdWKHHnYNtAIBPt-RM
		// 这里对客户端上传的token进行分割
		//parts := strings.SplitN(token, " ", 2)
		//if !(len(parts) == 2 || parts[0] == "Bearer") { //说明token格式不正确
		//	c.JSON(http.StatusOK, errdeal.NewResponse(errdeal.CodeInvalidTokenErr)) // 无效的token
		//	c.Abort()
		//	return
		//}

		// 解析token
		parseToken, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, errdeal.NewResponse(errdeal.CodeInvalidTokenErr)) // 无效的token
			c.Abort()
			return
		}

		// 中间件判断用户登录的话 将用户id保存在 context里面 保证后续可以获取到
		c.Set(ContextUserID, parseToken.UserID)

		c.Next()
	}
}

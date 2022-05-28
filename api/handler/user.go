package handler

import (
	"context"
	"dousheng/api/rpc"
	"dousheng/pkg/errdeal"
	user "dousheng/user/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	//var param UserRegisterParam
	//if err := c.ShouldBindJSON(&param); err != nil {
	//	HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
	//	return
	//}
	//req := user.DouyinUserRegisterRequest{
	//	Username: param.Username,
	//	Password: param.Password,
	//}
	req := user.DouyinUserRegisterRequest{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}
	if len(req.Username) <= 0 {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("用户名不能为空"))
		return
	}
	if len(req.Password) <= 5 {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("密码不能少于6位"))
		return
	}
	res, err := rpc.Register(context.Background(), &req)
	if err != nil {
		response := errdeal.NewResponse(errdeal.CodeErr(10002)).WithErr(err)
		HttpResponse(c, response)
		return
	}
	// 成功
	response := errdeal.NewResponse(errdeal.CodeErr(res.StatusCode))
	HttpResponse(c, response)
}

func LoginHandler(c *gin.Context) {
	//var param UserRegisterParam
	//if err := c.ShouldBindJSON(&param); err != nil {
	//	HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
	//	return
	//}
	req := user.DouyinUserLoginRequest{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}
	if len(req.Username) <= 0 {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("用户名不能为空"))
		return
	}
	if len(req.Password) <= 5 {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("密码不能少于6位"))
		return
	}
	res, err := rpc.Login(context.Background(), &req)
	if err != nil {
		response := errdeal.NewResponse(errdeal.CodeErr(10002)).WithErr(err)
		HttpResponse(c, response)
		return
	}
	// 成功
	fmt.Println("1")
	response := errdeal.NewResponse(errdeal.CodeErr(res.StatusCode))
	fmt.Println("2")
	HttpResponse(c, response)
	fmt.Println("3")
}

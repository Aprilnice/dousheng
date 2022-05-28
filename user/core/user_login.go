package core

import (
	"context"
	"dousheng/pkg/dao/mysqldb"
	"dousheng/pkg/errdeal"
	"dousheng/user/service"
	user "dousheng/user/service"
	"fmt"
)

type UserResp struct {
	*errdeal.Response
}

func (*UserService) Login(ctx context.Context, req *service.DouyinUserLoginRequest, res *service.DouyinUserLoginResponse) (err error) {
	username := ctx.Value("username")
	fmt.Println(username)
	if !mysqldb.UserLogin(req) {
		// 账号密码错误
		LoginResponseErr(err).BindTo(res)
		fmt.Println("user core err :", err)
		return err
	}
	fmt.Println("register success")
	// 成功
	tmp := errdeal.NewResponse(errdeal.CodeSuccess).WithData("nil")
	res.StatusCode = tmp.StatusCode
	res.StatusMsg = tmp.StatusMessage
	res.UserId = mysqldb.GetUserId(req.GetUsername())
	return nil
}
func (c *UserResp) BindTo(response *user.DouyinUserLoginResponse) {
	response.StatusCode = c.Response.StatusCode
	response.StatusMsg = c.Response.StatusMessage
}

func LoginResponseErr(err error) *UserResp {
	var resp *errdeal.Response
	// 如果是自定义的那些错误
	if codeErr, ok := err.(errdeal.CodeErr); ok {
		resp = errdeal.NewResponse(codeErr)
	}
	// 否则直接视为服务错误
	resp = errdeal.NewResponse(errdeal.CodeServiceErr).WithErr(err)
	return &UserResp{
		resp,
	}
}

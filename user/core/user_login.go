package core

import (
	"context"
	"dousheng/pkg/dao/mysqldb"
	"dousheng/pkg/errdeal"
	"dousheng/user/service"
	"fmt"
)

func (*UserService) Login(ctx context.Context, req *service.DouyinUserLoginRequest, res *service.DouyinUserLoginResponse) (err error) {
	if !mysqldb.UserLogin(req) {
		// 账号密码错误
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

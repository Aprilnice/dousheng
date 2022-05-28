package core

import (
	"context"
	"dousheng/pkg/dao/mysqldb"
	"dousheng/pkg/doushengjwt"
	"dousheng/pkg/errdeal"
	"dousheng/pkg/snowflaker"
	"dousheng/user/service"
	"fmt"
)

func (*UserService) Register(ctx context.Context, req *service.DouyinUserRegisterRequest, res *service.DouyinUserRegisterResponse) (err error) {
	userModel := new(mysqldb.User)
	res.UserId = snowflaker.NextID()
	res.Token, err = doushengjwt.GenerateToken(req.GetUsername(), res.UserId)
	if err != nil {
		return err
	}
	if err = userModel.BindWithReq(req); err != nil {
		return err
	}
	userModel.UserID = res.UserId

	if err = mysqldb.UserRegister(userModel); err != nil {
		// 出现错误  这里一般都是数据库错误
		return err
	}
	fmt.Println("register success")
	// 成功
	tmp := errdeal.NewResponse(errdeal.CodeSuccess).WithData("nil")
	res.StatusCode = tmp.StatusCode
	res.StatusMsg = tmp.StatusMessage
	return nil
}

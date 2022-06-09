package core

import (
	"context"
	"dousheng/cmd/user/dal/mysqldb"
	"dousheng/cmd/user/service"
	"dousheng/pkg/doushengjwt"
	"dousheng/pkg/errdeal"
	middlewares "dousheng/pkg/middleware"
	"dousheng/pkg/snowflaker"
)

func (*UserService) Register(ctx context.Context, req *service.DouyinUserRegisterRequest, res *service.DouyinUserRegisterResponse) (err error) {
	userModel := new(mysqldb.User)
	res.UserId = snowflaker.NextID()
	res.Token, err = doushengjwt.GenerateToken(req.GetUsername(), res.UserId)
	if err != nil {
		return err
	}
	password := middlewares.Md5Crypt("dousheng")
	req.Password = password
	if err = userModel.BindWithReq(req); err != nil {
		return err
	}
	userModel.UserID = res.UserId
	// 出现错误  这里一般都是数据库错误
	if err = mysqldb.UserRegister(userModel); err != nil {
		tmp := errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("用户名已存在，请重新输入")
		res.StatusCode = tmp.StatusCode
		res.StatusMsg = tmp.StatusMessage
		return err
	}
	userInfoModel := new(mysqldb.UserInfo)
	userInfoModel.UserId = userModel.UserID
	userInfoModel.Name = userModel.Username
	userInfoModel.FollowerCount = 0
	userInfoModel.FollowCount = 0
	userInfoModel.IsFollow = false
	if err = mysqldb.SetUserInfo(userInfoModel); err != nil {
		// 出现错误  这里一般都是数据库错误
		return err
	}
	// 成功
	tmp := errdeal.NewResponse(errdeal.CodeSuccess)
	res.StatusCode = tmp.StatusCode
	res.StatusMsg = tmp.StatusMessage
	return nil
}

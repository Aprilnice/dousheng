package handler

import (
	"context"
	"dousheng/api/rpc"
	"dousheng/pkg/errdeal"
	user "dousheng/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		c.JSON(http.StatusOK, errdeal.LoginResponse{
			StatusCode:    int32(errdeal.CodeParamErr),
			StatusMessage: "用户名不能为空",
		})
		return
	}
	if len(req.Password) <= 5 {
		c.JSON(http.StatusOK, errdeal.LoginResponse{
			StatusCode:    int32(errdeal.CodeParamErr),
			StatusMessage: "密码不能少于6位",
		})
		return
	}
	res, err := rpc.Register(context.Background(), &req)
	if err != nil {
		response := errdeal.NewResponse(errdeal.CodeParamErr).WithData("用户名已重复，请重新输入")
		HttpResponse(c, response)
		return
	}
	// 成功
	c.JSON(http.StatusOK, errdeal.LoginResponse{
		StatusCode:    res.StatusCode,
		StatusMessage: res.StatusMsg,
		UserId:        res.UserId,
		Token:         res.Token,
	})
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
		c.JSON(http.StatusOK, errdeal.LoginResponse{
			StatusCode:    int32(errdeal.CodeParamErr),
			StatusMessage: "用户名不能为空",
		})
		return
	}
	if len(req.Password) <= 5 {
		c.JSON(http.StatusOK, errdeal.LoginResponse{
			StatusCode:    int32(errdeal.CodeParamErr),
			StatusMessage: "密码不能少于6位",
		})
		return
	}
	res, err := rpc.Login(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusOK, errdeal.LoginResponse{
			StatusCode: int32(errdeal.CodeParamErr),
		})
		return
	}
	// 成功
	c.JSON(http.StatusOK, errdeal.LoginResponse{
		StatusCode:    res.StatusCode,
		StatusMessage: res.StatusMsg,
		UserId:        res.UserId,
		Token:         res.Token,
	})
}

func UserInfoHandler(c *gin.Context) {
	UserId, _ := strconv.Atoi(c.Query("user_id"))
	req := user.DouyinUserRequest{
		UserId: int64(UserId),
		Token:  c.Query("token"),
	}
	res, err := rpc.UserInfo(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusOK, errdeal.UserResp{
			StatusCode:    int32(errdeal.CodeParamErr),
			StatusMessage: res.StatusMsg,
			User:          nil,
		})
		return
	}
	// 成功
	UserData := new(errdeal.UserInfo)
	UserData.Id = res.User.Id
	UserData.Name = res.User.Name
	UserData.FollowCount = res.User.FollowCount
	UserData.FollowerCount = res.User.FollowerCount
	UserData.IsFollow = res.User.IsFollow
	c.JSON(http.StatusOK, errdeal.UserResp{
		StatusCode:    res.StatusCode,
		StatusMessage: res.StatusMsg,
		User:          UserData,
	})
}

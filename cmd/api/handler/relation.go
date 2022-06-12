package handler

import (
	"context"
	relation "dousheng/cmd/relation/service"
	"dousheng/pkg/constant"
	"dousheng/pkg/doushengjwt"
	"dousheng/pkg/errdeal"
	"fmt"
	"github.com/gin-gonic/gin"
)

// RelationActionHandler 关注或取关
func RelationActionHandler(c *gin.Context) {
	var relationParam RelationActionParam
	if err := c.ShouldBindQuery(&relationParam); err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}

	// 解析token
	token, err := doushengjwt.ParseToken(relationParam.Token)
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("无效的 token"))
		return
	}
	// 绑定参数
	relationReq := relation.RelationActionRequest{
		Token:      relationParam.Token,
		UserId:     token.UserID,
		ToUserId:   relationParam.ToUserId,
		ActionType: relationParam.ActionType,
	}
	relationRPC := c.Keys[constant.ClientRelation].(relation.RelationService)
	resp, err := relationRPC.RelationAction(context.Background(), &relationReq)
	if err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeServiceErr))
		return
	}
	HttpResponse(c, resp)
	return
}

// FollowListHandler 关注列表
func FollowListHandler(c *gin.Context) {
	var relationListParam RelationListParam
	if err := c.ShouldBindQuery(&relationListParam); err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}

	// 解析token
	token, err := doushengjwt.ParseToken(relationListParam.Token)
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("无效的 token"))
		return
	}
	// 绑定参数
	followReq := relation.FollowListRequest{
		Token:  relationListParam.Token,
		UserId: token.UserID,
	}
	relationRPC := c.Keys[constant.ClientRelation].(relation.RelationService)
	resp, err := relationRPC.FollowList(context.Background(), &followReq)
	if err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeServiceErr))
		return
	}
	HttpResponse(c, resp)
	return
}

// FollowerListHandler 粉丝列表
func FollowerListHandler(c *gin.Context) {
	var relationListParam RelationListParam
	if err := c.ShouldBindQuery(&relationListParam); err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}

	// 解析token
	token, err := doushengjwt.ParseToken(relationListParam.Token)
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("无效的 token"))
		return
	}
	// 绑定参数
	followerReq := relation.FollowerListRequest{
		Token:  relationListParam.Token,
		UserId: token.UserID,
	}
	relationRPC := c.Keys[constant.ClientRelation].(relation.RelationService)
	resp, err := relationRPC.FollowerList(context.Background(), &followerReq)
	if err != nil {
		fmt.Println(err)
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeServiceErr))
		return
	}
	HttpResponse(c, resp)
	return
}

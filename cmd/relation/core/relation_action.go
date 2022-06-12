package core

import (
	"context"
	"dousheng/cmd/relation/dal/mysqldb"
	"dousheng/cmd/relation/dal/redisdb"
	relation "dousheng/cmd/relation/service"
	"dousheng/pkg/errdeal"
)

func (*RelationService) RelationAction(ctx context.Context, req *relation.RelationActionRequest,
	resp *relation.RelationActionResponse) error {

	if req.ActionType == 1 {
		return doFollow(ctx, req, resp)
	} else if req.ActionType == 2 {
		return cancelFollow(ctx, req, resp)
	} else {
		r := errdeal.NewResponse(errdeal.CodeParamErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMessage
		return errdeal.CodeParamErr
	}

}

func doFollow(ctx context.Context, req *relation.RelationActionRequest,
	resp *relation.RelationActionResponse) error {
	if err := mysqldb.DoFollow(req.GetUserId(), req.GetToUserId()); err != nil {
		// 出现错误  这里一般都是数据库错误
		r := errdeal.NewResponse(errdeal.CodeServiceErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMessage
		return err
	}

	if err := redisdb.DoFollow(req.GetUserId(), req.GetToUserId()); err != nil {
		r := errdeal.NewResponse(errdeal.CodeServiceErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMessage
		return err
	}
	r := errdeal.NewResponse(errdeal.CodeSuccess)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMessage
	return nil
}

func cancelFollow(ctx context.Context, req *relation.RelationActionRequest,
	resp *relation.RelationActionResponse) error {
	if err := mysqldb.CancelFollow(req.GetUserId(), req.GetToUserId()); err != nil {
		// 出现错误  这里一般都是数据库错误
		r := errdeal.NewResponse(errdeal.CodeServiceErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMessage
		return err
	}

	if err := redisdb.CancelFollow(req.GetUserId(), req.GetToUserId()); err != nil {
		r := errdeal.NewResponse(errdeal.CodeServiceErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMessage
		return err
	}
	r := errdeal.NewResponse(errdeal.CodeSuccess)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMessage
	return nil
}

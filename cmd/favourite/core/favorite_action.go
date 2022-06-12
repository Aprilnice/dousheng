package core

import (
	"context"
	favoriteDB "dousheng/cmd/favorite/dal/mysqldb"
	favoriteRDB "dousheng/cmd/favorite/dal/redisdb"
	favorite "dousheng/cmd/favorite/service"
	"dousheng/pkg/errdeal"
	"fmt"
)

// FavoriteAction 点赞
func (*FavoriteService) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest, resp *favorite.FavoriteActionResponse) error {

	if req.ActionType == 1 {

		return doFavorite(req, resp)
	} else if req.ActionType == 2 {
		return cancelFavorite(req, resp)
	} else {
		ResponseError(errdeal.CodeParamErr).FavoriteActionResponse(resp)
		return errdeal.CodeParamErr
	}
}

func doFavorite(req *favorite.FavoriteActionRequest,
	resp *favorite.FavoriteActionResponse) error {
	var favoriteReq favoriteDB.Favorite
	favoriteReq.BindWithReq(req)
	if err := favoriteDB.CreateFavorite(&favoriteReq); err != nil {
		fmt.Println("err: ", err)
		ResponseError(errdeal.CodeServiceErr).FavoriteActionResponse(resp)
		return err
	}

	// redis
	if err := favoriteRDB.CreateFavorite(req); err != nil {
		fmt.Println("err: ", err)
		ResponseError(errdeal.CodeServiceErr).FavoriteActionResponse(resp)
		return err
	}

	ResponseSuccess().FavoriteActionResponse(resp)
	return nil
}

func cancelFavorite(req *favorite.FavoriteActionRequest,
	resp *favorite.FavoriteActionResponse) error {
	var favoriteReq favoriteDB.Favorite
	favoriteReq.BindWithReq(req)
	if err := favoriteDB.CancelFavorite(&favoriteReq); err != nil {
		ResponseError(errdeal.CodeServiceErr).FavoriteActionResponse(resp)
		return err
	}
	if err := favoriteRDB.CancelFavorite(req); err != nil {
		ResponseError(errdeal.CodeServiceErr).FavoriteActionResponse(resp)
		return err
	}
	ResponseSuccess().FavoriteActionResponse(resp)
	return nil
}

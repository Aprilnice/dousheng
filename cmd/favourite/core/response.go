package core

import (
	favorite "dousheng/cmd/favourite/service"
	"dousheng/pkg/errdeal"
)

type FavouriteResp struct {
	*errdeal.Response
}

// FavoriteActionResponse 这里实现将错误信息绑定到 自己的service Response 中
func (f *FavouriteResp) FavoriteActionResponse(response *favorite.FavoriteActionResponse) {
	response.StatusCode = f.Response.StatusCode
	response.StatusMsg = f.Response.StatusMessage
}

func (f *FavouriteResp) FavoriteListResponse(response *favorite.FavoriteListResponse) {
	response.StatusCode = f.Response.StatusCode
	response.StatusMsg = f.Response.StatusMessage
}

func ResponseError(err error) *FavouriteResp {
	var resp *errdeal.Response
	// 如果是自定义的那些错误
	if codeErr, ok := err.(errdeal.CodeErr); ok {
		resp = errdeal.NewResponse(codeErr)
	}
	// 否则直接视为服务错误
	resp = errdeal.NewResponse(errdeal.CodeServiceErr).WithErr(err)
	return &FavouriteResp{
		resp,
	}
}

func ResponseSuccess() *FavouriteResp {
	return &FavouriteResp{
		errdeal.NewResponse(errdeal.CodeSuccess),
	}

}

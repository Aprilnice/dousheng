package core

import (
	favorite "dousheng/cmd/favorite/service"
	"dousheng/pkg/errdeal"
)

type FavoriteResp struct {
	*errdeal.Response
}

// FavoriteActionResponse 这里实现将错误信息绑定到 自己的service Response 中
func (f *FavoriteResp) FavoriteActionResponse(response *favorite.FavoriteActionResponse) {
	response.StatusCode = f.Response.StatusCode
	response.StatusMsg = f.Response.StatusMessage
}

func (f *FavoriteResp) FavoriteListResponse(response *favorite.FavoriteListResponse) {
	response.StatusCode = f.Response.StatusCode
	response.StatusMsg = f.Response.StatusMessage
}

func ResponseError(err error) *FavoriteResp {
	var resp *errdeal.Response
	// 如果是自定义的那些错误
	if codeErr, ok := err.(errdeal.CodeErr); ok {
		resp = errdeal.NewResponse(codeErr)
	}
	// 否则直接视为服务错误
	resp = errdeal.NewResponse(errdeal.CodeServiceErr).WithErr(err)
	return &FavoriteResp{
		resp,
	}
}

func ResponseSuccess() *FavoriteResp {
	return &FavoriteResp{
		errdeal.NewResponse(errdeal.CodeSuccess),
	}

}

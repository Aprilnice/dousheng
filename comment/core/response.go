package core

import (
	comment "dousheng/comment/service"
	"dousheng/pkg/errdeal"
)

type CommentResp struct {
	*errdeal.Response
}

// BindTo 这里实现将错误信息绑定到 自己的service Response 中
func (c *CommentResp) BindTo(response *comment.CommentResponse) {
	response.StatusCode = c.Response.StatusCode
	response.StatusMsg = c.Response.StatusMessage
}

func ResponseSuccess(data interface{}) *CommentResp {
	resp := errdeal.NewResponse(errdeal.CodeSuccess).WithData(data)
	return &CommentResp{
		resp,
	}
}

func ResponseErr(err error) *CommentResp {
	var resp *errdeal.Response
	// 如果是自定义的那些错误
	if codeErr, ok := err.(errdeal.CodeErr); ok {
		resp = errdeal.NewResponse(codeErr)
	}
	// 否则直接视为服务错误
	resp = errdeal.NewResponse(errdeal.CodeServiceErr).WithErr(err)
	return &CommentResp{
		resp,
	}
}

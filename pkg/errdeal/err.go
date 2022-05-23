package errdeal

import (
	"fmt"
)

type CodeErr int32

// 常量错误码
const (
	CodeSuccess         CodeErr = 0
	CodeServiceErr      CodeErr = 10001
	CodeParamErr        CodeErr = 10002
	CodeInvalidTokenErr CodeErr = 10003
)

// 错误码对应的消息
var codeMessage = map[CodeErr]string{
	CodeSuccess:         "Success",
	CodeServiceErr:      "服务繁忙",
	CodeParamErr:        "请求参数错误",
	CodeInvalidTokenErr: "无效的Token",
}

func (c CodeErr) Message() string {
	msg, ok := codeMessage[c]
	if !ok { // 如果没有这样的错误码 直接返回服务繁忙
		msg = codeMessage[CodeServiceErr]
	}
	return msg
}

func (c CodeErr) Error() string {
	return fmt.Sprintf("code: %d, message: %s", c, c.Message())
}

// Response 定义一个返回响应的结构体
// StatusCode:    自定义的响应码
// StatusMessage: 响应码绑定的错误消息
type Response struct {
	StatusCode    int32       `json:"code"`
	StatusMessage string      `json:"message"`
	Data          interface{} `json:"data,omitempty"` // 忽略掉空值
}

// NewResponse 创建一个响应 错误码初始化
func NewResponse(code CodeErr) *Response {
	return &Response{
		StatusCode:    int32(code),
		StatusMessage: code.Message(),
	}
}

// WithMsg 自定义消息
func (r *Response) WithMsg(msg string) *Response {
	r.StatusMessage = msg
	return r
}

// WithErr 自定义错误消息
func (r *Response) WithErr(err error) *Response {
	r.StatusMessage = err.Error()
	return r
}

// WithData 数据
func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}
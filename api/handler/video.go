package handler

import (
	"context"
	"dousheng/api/rpc"
	"dousheng/pkg/doushengjwt"
	"dousheng/pkg/errdeal"
	video "dousheng/video/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func VideoPublishHandler(c *gin.Context) {
	var param VideoPublishParam
	if err := c.ShouldBindJSON(&param); err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}

	// 解析token,获取userid
	claims, err := doushengjwt.ParseToken(param.Token)

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("视频上传错误"))
		return
	}
	bFile, _ := ioutil.ReadAll(file)
	req := video.DouyinPublishActionRequest{
		UserId: claims.UserID,
		Title: param.Title,
		Data:  bFile,
	}

	// rpc 调用
	resp, err := rpc.VideoPublish(context.Background(), &req)

	var response *errdeal.Response

	if err != nil {
		response = errdeal.NewResponse(errdeal.CodeErr(resp.StatusCode)).WithErr(err)
		HttpResponse(c, response)
		return
	}
	// 成功
	response = errdeal.NewResponse(errdeal.CodeErr(resp.StatusCode))
	HttpResponse(c, response)

}
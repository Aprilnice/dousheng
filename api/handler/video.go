package handler

import (
	"context"
	"dousheng/api/rpc"
	"dousheng/pkg/doushengjwt"
	"dousheng/pkg/errdeal"
	video "dousheng/video/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
)

// VideoPublishHandler 视频发布
func VideoPublishHandler(c *gin.Context) {
	// 获取参数
	var param VideoPublishParam
	if err := c.ShouldBindJSON(&param); err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}

	// 解析token,获取userid
	claims, err := doushengjwt.ParseToken(param.Token)

	// 获取视频文件
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("视频上传错误"))
		return
	}

	// 获得文件二进制流
	bFile, _ := ioutil.ReadAll(file)
	req := video.DouyinPublishActionRequest{
		UserId: claims.UserID,
		Title: param.Title,
		Data:  bFile,
	}

	// rpc 调用
	resp, err := rpc.VideoPublish(context.Background(), &req)

	// 处理错误
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

// VideoPlayHandler 视频播放
func VideoPlayHandler(c *gin.Context) {
	// 获取参数
	id := c.Query("video_id")
	if id == "" {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}
	tmp, _ := strconv.ParseInt(id, 10, 64)
	req := video.PlayVideoReq{
		Id: tmp,
	}

	// rpc 调用
	resp, _ := rpc.VideoPlay(context.Background(), &req)

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Writer.Write(resp.Data)
}

// GetCoverHandler 获取封面
func GetCoverHandler(c *gin.Context) {
	// 获取参数
	id := c.Query("cover_id")
	if id == "" {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("参数解析错误"))
		return
	}
	tmp, _ := strconv.ParseInt(id, 10, 64)
	req := video.GetCoverReq{
		Id: tmp,
	}

	// rpc 调用
	resp, _ := rpc.CoverGet(context.Background(), &req)

	c.Header("Content-Type", "image/jpeg")
	c.Writer.Write(resp.Data)
}

// GetVideoFeedHandler 获取视频流
func GetVideoFeedHandler(c *gin.Context) {
	// 获取参数
	latest := c.Query("latest_time")
	token := c.Query("token")
	var tmp,id int64
	id = 0
	if latest == "" {
		tmp = time.Now().UnixMilli()
	}else{
		tmp, _ = strconv.ParseInt(latest, 10, 64)
	}

	// 如果登录,解析token
	if token != "" {
		claims,_ := doushengjwt.ParseToken(token)
		id = claims.UserID
	}
	req := video.DouyinFeedRequest{
		UserId: id,
		LatestTime: tmp,
	}

	// rpc调用
	resp, err := rpc.VideoFeed(context.Background(), &req)


	// 处理错误
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
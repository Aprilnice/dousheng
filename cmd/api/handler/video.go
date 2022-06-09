package handler

import (
	"context"
	video "dousheng/cmd/video/service"
	"dousheng/pkg/constant"
	"dousheng/pkg/doushengjwt"
	"dousheng/pkg/errdeal"
	middlewares "dousheng/pkg/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// VideoPublishHandler 视频发布
func VideoPublishHandler(c *gin.Context) {

	// 获取userid
	userId := c.GetInt64(middlewares.ContextUserID)

	// 获取解析后表单
	form, err := c.MultipartForm()
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("视频上传错误"))
		return
	}

	// 获取视频文件
	file := form.File["data"][0]
	fh, err := file.Open()
	if err != nil {
		HttpResponse(c, errdeal.NewResponse(errdeal.CodeParamErr).WithMsg("视频获取错误"))
		return
	}

	// 获得文件二进制流
	bFile, _ := ioutil.ReadAll(fh)

	// 关闭文件
	fh.Close()

	if len(bFile) == 0 {
		fmt.Println("文件不存在")
	}

	// 绑定参数
	req := video.DouyinPublishActionRequest{
		UserId: userId,
		Title:  (form.Value)["title"][0],
		Data:   bFile,
	}

	// rpc 调用
	videoRPC := c.Keys[constant.ClientVideo].(video.VideoModuleService)
	resp, err := videoRPC.VideoPublish(context.Background(), &req)

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

	// 参数绑定
	req := video.PlayVideoReq{
		Id: tmp,
	}

	// rpc 调用
	videoRPC := c.Keys[constant.ClientVideo].(video.VideoModuleService)
	resp, _ := videoRPC.PlayVideo(context.Background(), &req)

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
	videoRPC := c.Keys[constant.ClientVideo].(video.VideoModuleService)
	resp, _ := videoRPC.GetCover(context.Background(), &req)

	c.Header("Content-Type", "image/jpeg")
	c.Writer.Write(resp.Data)
}

// GetVideoFeedHandler 获取视频流
func GetVideoFeedHandler(c *gin.Context) {
	// 获取参数
	latest := c.Query("latest_time")
	token := c.Query("token")
	var tmp, id int64
	id = 0

	// 没传时间用当前时间
	if latest == "" {
		tmp = time.Now().UnixMilli()
	} else {
		tmp, _ = strconv.ParseInt(latest, 10, 64)
	}

	// 如果登录,解析token
	if token != "" {
		claims, _ := doushengjwt.ParseToken(token)
		id = claims.UserID
	}

	// 绑定参数
	req := video.DouyinFeedRequest{
		UserId:     id,
		LatestTime: tmp,
	}

	// rpc调用
	videoRPC := c.Keys[constant.ClientVideo].(video.VideoModuleService)
	resp, err := videoRPC.VideoFeed(context.Background(), &req)

	// 处理错误
	var response *errdeal.Response
	if err != nil {
		response = errdeal.NewResponse(errdeal.CodeErr(resp.StatusCode)).WithErr(err)
		c.JSON(http.StatusOK, errdeal.FeedResponse{
			StatusCode:    response.StatusCode,
			StatusMessage: response.StatusMessage,
			NextTime:      tmp,
		})
		return
	}

	// 成功
	response = errdeal.NewResponse(errdeal.CodeErr(resp.StatusCode))
	c.JSON(http.StatusOK, errdeal.FeedResponse{
		StatusCode:    response.StatusCode,
		StatusMessage: response.StatusMessage,
		NextTime:      resp.NextTime,
		VideoList:     resp.VideoList,
	})
}

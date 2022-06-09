package core

import (
	"context"
	"dousheng/cmd/video/dal/mysqldb"
	"dousheng/cmd/video/dal/redisdb"
	video "dousheng/cmd/video/service"
	"dousheng/config"
	"dousheng/pkg/errdeal"
	"dousheng/pkg/snowflaker"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

// VideoPublish 上传视频
func (*VideoModuleService) VideoPublish(c context.Context, req *video.DouyinPublishActionRequest, resp *video.DouyinPublishActionResponse) (err error) {

	videoId := snowflaker.NextID()

	// 视频文件存在本地
	filePath := "./file/video/" + strconv.FormatInt(videoId, 10) + ".mp4"
	err = ioutil.WriteFile(filePath, req.Data, 0777)
	if err != nil {
		ResponsePublishErr(err, resp)
		return err
	}

	// 拼接播放地址
	playURL := fmt.Sprintf("%s:%s/play/?video_id=%s",
		config.Instance().BaseConfig.Host, // "192.168.43.241"
		config.Instance().BaseConfig.Port,
		strconv.FormatInt(videoId, 10),
	)

	// 获取视频信息
	videoModule := &mysqldb.VideoInfo{
		Id:            videoId,
		Title:         req.Title,
		AuthorId:      req.UserId,
		PlayUrl:       playURL, // 播放地址
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		// 获得毫秒级时间戳
		PublishTime: time.Now().UnixMilli(),
	}

	// 视频信息存入数据库
	if err = mysqldb.PublishVideo(videoModule); err != nil {
		// 出现错误  这里一般都是数据库错误
		ResponsePublishErr(err, resp)
		return err
	}

	// 视频信息存入redis的Hash表中
	if err = redisdb.AddVideoInfo(videoModule); err != nil {
		ResponsePublishErr(err, resp)
		return err
	}

	// 视频id存入有序集合feed,用于返回视频列表
	if err = redisdb.AddVideoId(videoModule.Id, videoModule.PublishTime); err != nil {
		ResponsePublishErr(err, resp)
		return err
	}

	// 成功
	tmp := errdeal.NewResponse(errdeal.CodeSuccess)
	resp.StatusCode = tmp.StatusCode
	resp.StatusMsg = tmp.StatusMessage
	return nil
}

func ResponsePublishErr(err error, resp *video.DouyinPublishActionResponse) {
	var errResp *errdeal.Response
	// 如果是自定义的那些错误
	if codeErr, ok := err.(errdeal.CodeErr); ok {
		errResp = errdeal.NewResponse(codeErr)
	}
	// 否则直接视为服务错误
	errResp = errdeal.NewResponse(errdeal.CodeServiceErr).WithErr(err)

	resp.StatusCode = errResp.StatusCode
	resp.StatusMsg = errResp.StatusMessage
}

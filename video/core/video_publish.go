package core

import (
	"context"
	"dousheng/pkg/dao/mysqldb"
	"dousheng/pkg/snowflaker"
	video "dousheng/video/service"
	"io/ioutil"
	"os"
	"strconv"
)

// VideoPublish 上传视频
func (*VideoModuleService) VideoPublish(c context.Context, req *video.DouyinPublishActionRequest, resp *video.DouyinPublishActionResponse) (err error) {

	videoId := snowflaker.NextID()

	// 视频文件存在本地
	filePath := "/dousheng/file/video/" + strconv.FormatInt(videoId, 10) + ".mp4"
	err = ioutil.WriteFile(filePath, req.Data, 0777)

	// 获取视频信息
	videoMoudle := &mysqldb.VideoInfo{
		Id: videoId,
		Title: req.Title,
		AuthorId: c.Value("authorId"),
		PlayUrl: "",
		CoverUrl: "",
		FavoriteCount: 0,
		CommentCount: 0,
	}

	// 视频信息存入数据库
	if err = mysqldb.PublishVideo(videoMoudle); err != nil {
		// 出现错误  这里一般都是数据库错误
		ResponseErr(err).BindTo(resp)
		return err
	}

	// 成功
	ResponseSuccess(nil).BindTo(resp)
	return nil
}

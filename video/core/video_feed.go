package core

import (
	"context"
	"dousheng/pkg/dao/mysqldb"
	"dousheng/pkg/errdeal"
	video "dousheng/video/service"
	"fmt"
)

// VideoFeed 视频流
func (*VideoModuleService) VideoFeed(c context.Context, req *video.DouyinFeedRequest, resp *video.DouyinFeedResponse) (err error) {
	// 游客状态下均为false
	//like := false
	//follow := false

	// 获取视频信息
	videos,err := mysqldb.GetVideoFeed(req.LatestTime)
	if err != nil {
		ResponseFeedErr(err, resp)
		resp.NextTime = req.LatestTime
		return err
	}

	//格式化
	tmp := new(video.Video)
	for i := range videos {
		// 登录状态
		//if req.UserId != 0 {
		//	// 获取like和follow
		//	// like =
		//	// follow =
		//}
		//tmp.IsFavorite = like

		// 获取视频作者信息
		//author,err := mysqldb.GetUserInfo((*videos)[i].AuthorId)
		//if err != nil {
		//	ResponseFeedErr(err, resp)
		//	resp.NextTime = req.LatestTime
		//	return err
		//}
		//tmp.Author = &video.User{
		//	Id: author.Id,
		//	Name: author.Name,
		//	FollowCount: author.FollowCount,
		//	FollowerCount: author.FollowerCount,
		//	IsFollow: follow,
		//}
		author := &video.User{
			Id: 0,
			Name: "test",
			FollowerCount: 0,
			FollowCount: 0,
			IsFollow: false,
		}

		tmp = &video.Video{
			Id: videos[i].Id,
			PlayUrl: videos[i].PlayUrl,
			CoverUrl: "",
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount: videos[i].CommentCount,
			Author: author,
		}

		// 添加至返回的视频列表中
		resp.VideoList = append(resp.VideoList,tmp)
	}

	resp.StatusCode = 0
	resp.StatusMsg = "Success"
	resp.NextTime = videos[len(videos)-1].PublishTime
	fmt.Print("resp=")
	fmt.Println(resp)
	return nil
}

func ResponseFeedErr(err error, resp *video.DouyinFeedResponse) {
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
package core

import (
	"context"
	"dousheng/pkg/dao/mysqldb"
	"dousheng/pkg/errdeal"
	video "dousheng/video/service"
)

// VideoFeed 视频流
func (*VideoModuleService) VideoFeed(c context.Context, req *video.DouyinFeedRequest, resp *video.DouyinFeedResponse) (err error) {
	// 游客状态下均为false
	like := false
	follow := false

	// 获取视频信息
	videos,err := mysqldb.GetVideoFeed(req.LatestTime)
	if err != nil {
		ResponseFeedErr(err, resp)
		resp.NextTime = req.LatestTime
		return err
	}

	// 格式化
	tmp := new(video.Video)
	for i := range *videos {
		tmp = &video.Video{
			Id: (*videos)[i].Id,
			PlayUrl: (*videos)[i].PlayUrl,
			CoverUrl: (*videos)[i].CoverUrl,
			FavoriteCount: (*videos)[i].FavoriteCount,
			CommentCount: (*videos)[i].CommentCount,
		}

		// 登录状态
		if req.UserId != 0 {
			// 获取like和follow
			// like =
			// follow =
		}
		tmp.IsFavorite = like

		// 获取视频作者信息
		auther,err := mysqldb.GetUserInfo((*videos)[i].AuthorId)
		if err != nil {
			ResponseFeedErr(err, resp)
			resp.NextTime = req.LatestTime
			return err
		}
		tmp.Author = &video.User{
			Id: auther.Id,
			Name: auther.Name,
			FollowCount: auther.FollowCount,
			FollowerCount: auther.FollowerCount,
			IsFollow: follow,
		}

		// 添加至返回的视频列表中
		resp.VideoList = append(resp.VideoList,tmp)
	}
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
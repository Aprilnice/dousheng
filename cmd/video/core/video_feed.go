package core

import (
	"context"
	"dousheng/cmd/video/dal/mysqldb"
	"dousheng/cmd/video/dal/redisdb"
	video "dousheng/cmd/video/service"
	"dousheng/pkg/errdeal"
	"errors"
)

// VideoFeed 视频流
func (*VideoModuleService) VideoFeed(c context.Context, req *video.DouyinFeedRequest, resp *video.DouyinFeedResponse) (err error) {
	// 游客状态下均为false
	like := false
	follow := false

	// 从redis获取视频信息
	videos, err := redisdb.GetFeed(req.LatestTime)
	if err != nil {
		if errors.Is(err, errors.New("redis: nil")) {
			// 从MySQL获取视频信息
			videos, err = mysqldb.GetVideoFeed(req.LatestTime)
			if err != nil {
				ResponseFeedErr(err, resp)
				resp.NextTime = req.LatestTime
				return err
			}
		} else {
			return err
		}
	}

	//格式化
	tmp := new(video.Video)
	for i := range videos {
		// 登录状态
		if req.UserId != 0 {
			// 获取like和follow
			like = redisdb.GetLike(req.UserId, videos[i].Id)
			follow = redisdb.GetFollow(req.UserId, videos[i].AuthorId)
		}

		// 从redis获取作者信息
		author, err := redisdb.GetAuthorInfo(videos[i].AuthorId)
		if err != nil || len(videos) == 0 {
			// 从MySQL获取视频作者信息
			author, err = mysqldb.GetUserInfo(videos[i].AuthorId)
			if err != nil {
				ResponseFeedErr(err, resp)
				resp.NextTime = req.LatestTime
				return err
			}
		}

		tmp = &video.Video{
			Id:            videos[i].Id,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    like,
		}

		tmp.Author = &video.User{
			Id:            author.UserId,
			Name:          author.Name,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      follow,
		}

		// 添加至返回的视频列表中
		resp.VideoList = append(resp.VideoList, tmp)
	}

	resp.StatusCode = 0
	resp.StatusMsg = "Success"

	if len(videos) == 0 {
		resp.NextTime = req.LatestTime
	} else {
		resp.NextTime = videos[len(videos)-1].PublishTime
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

package core

import (
	"context"
	"dousheng/cmd/video/dal/mysqldb"
	"dousheng/cmd/video/dal/redisdb"
	video "dousheng/cmd/video/service"
	"errors"
)

// GetVideoList 获取用户视频列表
func (*VideoModuleService) GetVideoList(c context.Context, req *video.GetVideoListReq, resp *video.GetVideoListResp) (err error) {
	// 从redis获取视频信息
	videos, err := redisdb.GetVideoList(req.UserId)
	if err != nil {
		if errors.Is(err, errors.New("redis: nil")) {
			//从MySQL获取视频信息
			videos, err = mysqldb.GetVideoList(req.UserId)
			if err != nil {
				resp.StatusCode = 10001
				resp.StatusMsg = "服务繁忙"
				return err
			}
		}else {
			resp.StatusCode = 10001
			resp.StatusMsg = "服务繁忙"
			return err
		}
	}

	// 从redis获取用户信息
	author, err := redisdb.GetAuthorInfo(req.UserId)
	if err != nil{
		if errors.Is(err, errors.New("redis: nil")) {
			// 从MySQL获取视频作者信息
			author, err = mysqldb.GetUserInfo(req.UserId)
			if err != nil {
				resp.StatusCode = 10001
				resp.StatusMsg = "服务繁忙"
				return err
			}
		}else {
			resp.StatusCode = 10001
			resp.StatusMsg = "服务繁忙"
			return err
		}
	}

	//格式化
	tmp := new(video.Video)
	for i := range videos {
		tmp = &video.Video{
			Id:            videos[i].Id,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    true,
		}

		tmp.Author = &video.User{
			Id:            author.UserId,
			Name:          author.Name,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      true,
		}

		// 添加至返回的视频列表中
		resp.VideoList = append(resp.VideoList, tmp)
	}

	resp.StatusCode = 0
	resp.StatusMsg = "Success"
	return nil
}
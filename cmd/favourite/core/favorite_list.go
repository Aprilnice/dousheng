package core

import (
	"context"
	favoriteDB "dousheng/cmd/favorite/dal/mysqldb"
	favoriteRDB "dousheng/cmd/favorite/dal/redisdb"
	favorite "dousheng/cmd/favorite/service"
	userDB "dousheng/cmd/user/dal/mysqldb"
	videoDB "dousheng/cmd/video/dal/mysqldb"
	"fmt"
	"sync"
)

// FavoriteList 点赞列表
func (*FavoriteService) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest,
	resp *favorite.FavoriteListResponse) (err error) {
	videoIds := favoriteRDB.FavoriteVideosID(req.UserId)
	if len(videoIds) == 0 { // 无点赞
		videoIds, err = favoriteDB.FavoriteVideosID(req.UserId)

		if err != nil {
			ResponseError(err).FavoriteListResponse(resp)
		}
		// 真的没点赞
		if len(videoIds) == 0 {
			ResponseSuccess().FavoriteListResponse(resp)
			return nil
		}
	}

	fmt.Println("videoIds: ", videoIds)
	var err1, err2 error
	var videosInfo []*videoDB.VideoInfo
	var authorsInfo []*userDB.UserInfo
	var wg sync.WaitGroup
	wg.Add(2)
	// 查询视频信息
	go func() {
		defer wg.Done()
		videosInfo, err1 = favoriteDB.QueryVideosInfo(videoIds)
	}()
	// 查询视频作者信息
	go func() {
		defer wg.Done()
		authorsInfo, err2 = favoriteDB.QueryAuthorsInfo(videoIds)
	}()

	wg.Wait()
	if err1 != nil {

		ResponseError(err1).FavoriteListResponse(resp)
	}
	if err2 != nil {

		ResponseError(err2).FavoriteListResponse(resp)
	}

	var mapAuthor = make(map[int64]*userDB.UserInfo, len(authorsInfo))
	var videosList = make([]*favorite.Video, len(videoIds))
	for _, info := range authorsInfo {
		mapAuthor[info.UserId] = info
	}
	fmt.Printf("%#v \n", videosInfo)
	for i, video := range videosInfo {
		// 查询到的视频作者信息
		var author *userDB.UserInfo
		ok := false
		if author, ok = mapAuthor[video.AuthorId]; !ok {
			continue
		}
		//
		u := favorite.User{
			Id:            author.UserId,
			Name:          author.Name,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      author.IsFollow,
		}
		// 完整的视频信息
		v := favorite.Video{
			Id:            video.Id,
			Author:        &u,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: 100,
			CommentCount:  30,
			IsFavorite:    true,
			Title:         video.Title,
		}
		videosList[i] = &v
	}
	ResponseSuccess().FavoriteListResponse(resp)
	resp.VideoList = videosList

	fmt.Printf("%#v", resp.VideoList)

	return nil
}

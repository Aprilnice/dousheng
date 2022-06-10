package rediskey

import "strings"

const (
	KeySep          = ":" // 分隔符
	KeyPrefix       = "dousheng"
	KeyFavoriteZSet = "favorite:user"
	KeyFeedZSet     = "video:feed"
	KeyVideoHash    = "video"
	KeyUserHash     = "user"
	KeyUserFollow   = "follow:user"
	KeyUserFollower = "follower:user"
)

func NewRedisKey(keys ...string) (redisKey string) {
	subKeys := make([]string, 0, len(keys)+1)
	subKeys = append(subKeys, KeyPrefix)
	subKeys = append(subKeys, keys...)
	redisKey = strings.Join(subKeys, KeySep)
	return
}

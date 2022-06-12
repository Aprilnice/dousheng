package middlewares

import (
	"dousheng/cmd/api/rpc"
	"dousheng/pkg/constant"
	"github.com/gin-gonic/gin"
)

func SetupServiceMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{}, 5)
		c.Keys[constant.ClientUser] = rpc.UserRPC
		c.Keys[constant.ClientVideo] = rpc.VideoRPC
		c.Keys[constant.ClientComment] = rpc.CommentRPC
		c.Keys[constant.ClientFavorite] = rpc.FavoriteRPC
		c.Keys[constant.ClientRelation] = rpc.RelationRPC
		c.Next()
	}
}

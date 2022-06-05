package handler

// CommentParam 评论相关参数
type CommentParam struct {
	UserId      int64  `form:"user_id" json:"user_id,omitempty,string"`                      // 用户id 加上string 防止 js 精度溢出
	Token       string `form:"token" json:"token,omitempty" binding:"required"`              // 用户鉴权
	VideoId     int64  `form:"video_id" json:"video_id,omitempty,string" binding:"required"` // 视频id
	ActionType  int32  `form:"action_type" json:"action_type,omitempty" binding:"required"`  // 1 发布评论 2删除评论
	CommentText string `form:"comment_text" json:"comment_text,omitempty"`                   // 评论内容 在action_type=1时使用
	CommentId   int64  `form:"comment_id" json:"comment_id,omitempty,string"`                // 评论id 删除评论时使用
}

type CommentListParam struct {
	VideoId int64  `form:"video_id" json:"video_id,omitempty,string" binding:"required"` // 视频id
	Token   string `form:"token" json:"token,omitempty" binding:"required"`              // 用户鉴权
}

// VideoPublishParam 视频发布相关参数
type VideoPublishParam struct {
	// Token 用户鉴权
	Token string `json:"token,omitempty" form:"token" binding:"required"`
	// Title 视频标题
	Title string `json:"title" form:"title"`
}

//UserRegisterParam 用户注册相关参数
type UserRegisterParam struct {
	//用户名
	Username string `json:"username" binding:"required"`
	//密码
	Password string `json:"password" binding:"required"`
}

type FavoriteActionParam struct {
	Token      string `form:"token" json:"token,omitempty" binding:"required"`                       // 用户鉴权
	VideoId    int64  `form:"video_id" json:"video_id,omitempty,string" binding:"required"`          // 视频id
	ActionType int32  `form:"action_type" json:"action_type,omitempty" binding:"required,oneof=1 2"` // 1 发布评论 2删除评论
}

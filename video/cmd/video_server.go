package cmd

import (
	"context"
	"dousheng/config"
	"dousheng/pkg/logx"
	"dousheng/pkg/mysql"
	"dousheng/video/pb"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// VideoServerContext 视频服务context结构体
type VideoServerContext struct {
	// Conf 配置结构体
	Conf config.Config
	// VideoModel 视频持久化结构体
	VideoModel VideoModel
}

// NewVideoServerContext 创建视频服务context实例
func NewVideoServerContext(conf config.Config, db *gorm.DB) *VideoServerContext {
	return &VideoServerContext{
		Conf: 		conf,
		VideoModel: NewVideoModel(db),
	}
}

// VideoServer  视频服务逻辑结构体
type VideoServer struct {
	context	   context.Context
	svcContext *VideoServerContext
}

// newVideoServer 创建视频服务实例
func newVideoServer(c context.Context, serviceContext *VideoServerContext) *VideoServer {
	return &VideoServer{
		context: 	c,
		svcContext: serviceContext,
	}
}

// StartVideoServer 开启视频服务
func StartVideoServer() {
	// 读取配置文件
	conf :=config.NewConfig().WithVideoConfig().WithLogConfig().WithMySQLConfig()
	// 配置日志
	logx.InitLogger(*conf)
	// 声明grpc服务
	server := grpc.NewServer()
	// 初始化连接MySQL
	db, err := mysql.Init(&conf.MySQLConfig,"video")
	if err != nil {
		log.Fatal(err)
	}

	// 注册服务
	pb.RegisterVideoModuleServer(
		server,
		newVideoServer(context.Background(),NewVideoServerContext(*conf,db)),
	)

	// 查询gRPC服务或调用gRPC方法
	reflection.Register(server)

	// 指定监听视频服务请求的端口
	lis, err := net.Listen("tcp",conf.VideoServerConfig.Port)
	if err != nil {
		logx.Log.Error("TCP Video 监听失败:",
			zap.Error(err))
		return
	}

	logx.Log.Info("视频服务正在监听",
		zap.String("端口", conf.VideoServerConfig.Port))

	//启用gRPC服务
	err = server.Serve(lis)
	if err != nil {
		logx.Log.Error("视频服务启动失败:",
			zap.Error(err))
	}
}

// VideoPublish 实现上传视频方法
func (v *VideoServer) VideoPublish(c context.Context, req *pb.DouyinPublishActionRequest) (resp *pb.DouyinPublishActionResponse, err error) {
	resp, err = v.svcContext.VideoModel.VideoPublish(req.Data)
	return resp, err
}

// VideoFeed 实现视频流方法
func (v *VideoServer) VideoFeed(c context.Context, req *pb.DouyinFeedRequest) (resp *pb.DouyinFeedResponse, err error) {
	resp, err = v.svcContext.VideoModel.VideoFeed(req.LatestTime)
	return resp, err
}

// PlayVideo 实现视频播放方法
func (v *VideoServer) PlayVideo(c context.Context, req *pb.PlayVideoReq) (resp *pb.PlayVideoResp, err error) {
	resp, err = v.svcContext.VideoModel.PlayVideo(req.Id)
	return resp, err
}

// GetCover 实现封面下载方法
func (v *VideoServer) GetCover(c context.Context, req *pb.GetCoverReq) (resp *pb.GetCoverResp, err error) {
	resp, err = v.svcContext.VideoModel.GetCover(req.Id)
	return resp, err
}
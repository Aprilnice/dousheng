package core

import (
	"context"
	video "dousheng/video/service"
	"io/ioutil"
	"os"
	"strconv"
)

// PlayVideo 视频播放/下载
func (*VideoModuleService) PlayVideo(c context.Context, req *video.PlayVideoReq, resp *video.PlayVideoResp) (err error) {
	//打开文件
	filePath := "./file/video/" + strconv.FormatInt(req.Id, 10) + ".mp4"
	fileTmp, err := os.Open(filePath)
	defer fileTmp.Close()
	if err != nil {
		return err
	}

	// 转为二进制文件流
	file, err := ioutil.ReadAll(fileTmp)
	if err != nil {
		return err
	}

	resp.Data = file
	return nil
}
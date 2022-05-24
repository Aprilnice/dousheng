package core

import (
	"context"
	video "dousheng/video/service"
	"io/ioutil"
	"os"
	"strconv"
)

// GetCover 封面下载
func (*VideoModuleService) GetCover(c context.Context, req *video.GetCoverReq, resp *video.GetCoverResp) error {
	//打开图片
	imagePath := "./file/image/" + strconv.FormatInt(req.Id, 10) + ".jpg"
	imageTmp, err := os.Open(imagePath)
	defer imageTmp.Close()
	if err != nil {
		return err
	}

	// 转为二进制文件流
	file, err := ioutil.ReadAll(imageTmp)
	if err != nil {
		return err
	}

	resp.Data = file
	return nil
}

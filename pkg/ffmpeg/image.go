package ffmpeg

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

var (
	_       = "ffmpeg -i ./file/video/58063111768248320.mp4 -vf select='eq(pict_type\\,I)' -vsync 2 -frames:v 1 -f image2 ./file/image/58063111768248320.jpg"
	cmdArgs = []string{"-i", "", "-vf", "select='eq(pict_type\\,I)'", "-vsync", "2", "-frames:v", "1", "-f", "image2", ""}
)

func CaptureVideoWin(videoID int64) {
	videoFile := fmt.Sprintf("./file/video/%d.mp4", videoID)
	imageFile := fmt.Sprintf("./file/image/%d.jpg", videoID)
	cmdArgs[1] = videoFile
	cmdArgs[len(cmdArgs)-1] = imageFile
	cmd := exec.Command("ffmpeg", cmdArgs...)
	//fmt.Println(cmd.Path)
	//fmt.Println(cmd.Args)
	//cmd := exec.Command("/bin/sh",  "-c"," go run /Users/java0904/goProject/demo-os-exec/demo07/main.go")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

}

func main() {
	CaptureVideoWin(58063111768248320)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func connect() (*net.TCPConn, error) {
	// 创建地址结构
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8700")
	if err != nil {
		// 返回错误
		return nil, err
	}
	// 建立TCP连接
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		// 返回错误
		return nil, err
	}
	return conn, nil
}

func stop(conn *net.TCPConn) error {
	// 停止全部
	data := make(map[string]interface{})
	data["type"] = "stop"
	data["params"] = map[string]bool{"all": true}
	// 创建JSON结构
	js, _ := json.Marshal(data)
	// 发送给xplay
	_, err := conn.Write(append(js, []byte("\n#End\n")...))
	// 返回结果
	return err
}

func play(
	conn *net.TCPConn,
	libName string,
	zIndex int,
	path string,
	top, left, width, height int,
	screenMode string, screenRotate int) error {
	// 播放素材
	data := make(map[string]interface{})
	params := make(map[string]interface{})
	params["zIndex"] = zIndex
	params["path"] = path
	params["top"] = top
	params["left"] = left
	params["width"] = width
	params["height"] = height
	params["screen_mode"] = screenMode
	params["screen_rotate"] = screenRotate
	data["id"] = fmt.Sprintf("PLAY_VIDEO_%d", time.Now().Unix())
	data["type"] = "play"
	data["libName"] = libName
	data["params"] = params
	// 创建JSON结构
	js, _ := json.Marshal(data)
	// 发送给xplay
	_, err := conn.Write(append(js, []byte("\n#End\n")...))
	// 返回结果
	return err
}

func getFiles(ph string) []string {
	files := make([]string, 0)
	// 遍历目录
	filepath.Walk(ph, func(ph string, f os.FileInfo, err error) error {
		// 文件不存在
		if f == nil {
			return nil
		}
		// 跳过文件夹
		if f.IsDir() {
			return nil
		}
		files = append(files, ph)
		// 返回空
		return nil
	})
	return files
}

func main() {

	// 设置一下日志的结构
	log.SetFlags(log.Lshortfile | log.Ltime | log.Lmicroseconds)

	// 连接 xplay ...
	conn, err := connect()
	if err != nil {
		// 连接失败 !!!
		log.Panic(err)
	}
	// 结束时断开
	defer conn.Close()

	// 停止全部
	err = stop(conn)
	if err != nil {
		log.Println(err.Error())
	}

	for {
		// 获取目录素材文件
		files := getFiles("E:/test")
		if len(files) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}
		// 遍历素材
		for _, f := range files {
			// 是否为视频文件
			if strings.HasSuffix(f, ".mp4") {
				// 播放视频
				err = play(conn, "video", 10, f, 0, 0, 1920, 1080, "landscape", 0)
				if err != nil {
					log.Println(err.Error())
				}
			}
			// 是否为图片文件
			if strings.HasSuffix(f, ".jpg") || strings.HasSuffix(f, ".png") {
				// 播放视频
				err = play(conn, "pic", 10, f, 0, 0, 1920, 1080, "landscape", 0)
				if err != nil {
					log.Println(err.Error())
				}
			}
			// 间隔5s切换
			time.Sleep(5 * time.Second)
		}
	}
}

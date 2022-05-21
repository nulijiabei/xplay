package main

// 20220520

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Sync Wait Group
var wg sync.WaitGroup

// Connect Address
var addr = flag.String("addr", "127.0.0.1:8700", "xplay tcp address")

// Query
var query = flag.Bool("query", false, "query playing status")

// Snap
var snap = flag.Bool("snap", false, "screen snapshot")

// Move
var move = flag.Bool("move", false, "move playing")

// Change
var change = flag.Bool("change", false, "change playing")

// General
// var id = flag.String("id", "", "debug id") // 自动生成
var start = flag.Int64("start", -1, "start time(ms)")

// Stop
var stop = flag.Bool("stop", false, "stop or play")
var all = flag.Bool("all", false, "stop all")
var ids = flag.String("ids", "", "stop ids")

// Play
var play = flag.Bool("play", false, "play or stop")
var libName = flag.String("libName", "", "video、pic、sequence、camera、gif、qrcode、text、scroll、background")

// Play Params
var zIndex = flag.Int("zIndex", 10, "1-999")
var rect = flag.String("rect", "0,0,1920,1080", "")
var screen_mode = flag.String("screen_mode", "landscape", "landscape、portrait")
var screen_rotate = flag.Int("screen_rotate", 0, "landscape: 0、180 or portrait: 90、270")
var path = flag.String("path", "", "file path")
var content = flag.String("content", "", "data content")
var offset = flag.Int64("offset", -1, "video offset(ms)")
var timeout = flag.Int64("timeout", -1, "video stream timeout(ms)")

// Change Params
var newIndex = flag.Int("newIndex", 0, "1-999")

// Text
var font_ttf = flag.String("font_ttf", "", "TrueTypeFont")
var font_size = flag.Int("font_size", 18, "")
var color = flag.String("color", "rgba(0,128,0,100%)", "")
var bgcolor = flag.String("bgcolor", "rgba(0,0,0,20%)", "")
var align = flag.String("align", "center", "center、right、left")
var style = flag.String("style", "normal", "normal、bold、italic、underline、strikethrough")

// Camera
var device = flag.String("device", "/dev/video0", "camera device")
var camera_width = flag.Int("camera_width", 1280, "camera video width")
var camera_height = flag.Int("camera_height", 720, "camera video height")

// Scroll
var speed = flag.Int("speed", 1, "move pixel / frame")
var orientation = flag.String("orientation", "horizontal", "horizontal | vertical")

// Toast
var toast_type = flag.String("toast_type", "notice", "notice、success、warning、error")
var duration = flag.Int("duration", 0, "toast duration(s)")

type XPlay struct {
	conn *net.TCPConn
}

func NewXPlay() *XPlay {
	xplay := new(XPlay)
	xplay.conn = nil
	return xplay
}

// 连接 XPlay
func (this *XPlay) connect(_addr string) error {
	addr, err := net.ResolveTCPAddr("tcp", _addr)
	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return err
	}
	this.conn = conn
	return nil
}

// 从 XPlay 接收返回
func (this *XPlay) result() {
	fmt.Println(">>> RESULT >>>")
	br := bufio.NewReader(this.conn)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			log.Println(err.Error())
			break
		}
		data := string(line)
		fmt.Println(data)
		if strings.Trim(data, " ") == "#End" {
			break
		}
	}
}

// 发送指令到 XPlay
func (this *XPlay) send(_data map[string]interface{}) error {
	js, _ := json.MarshalIndent(_data, "", "  ")
	fmt.Println(">>> SEND >>>")
	fmt.Println(string(js))
	_, err := this.conn.Write(append(js, []byte("\n#End\n")...))
	return err
}

// 停止全部
func (this *XPlay) stop_all() error {
	data := make(map[string]interface{})
	data["id"] = fmt.Sprintf("STOP_%d", time.Now().Unix())
	data["type"] = "stop"
	data["start"] = *start
	data["params"] = map[string]bool{"all": true}
	return this.send(data)
}

// 停止指定层
func (this *XPlay) stop(_ids string) error {
	vs := strings.Split(_ids, ",")
	data := make(map[string]interface{})
	params := make(map[string]interface{})
	params["ids"] = vs
	data["id"] = fmt.Sprintf("STOP_%d", time.Now().Unix())
	data["type"] = "stop"
	data["start"] = *start
	data["params"] = params
	return this.send(data)
}

func (this *XPlay) play() error {
	params := make(map[string]interface{})
	rs := strings.Split(*rect, ",")
	x, _ := strconv.Atoi(rs[0])
	y, _ := strconv.Atoi(rs[1])
	width, _ := strconv.Atoi(rs[2])
	height, _ := strconv.Atoi(rs[3])
	if (*libName) == "sequence" {
		params["zIndex"] = *zIndex
		params["left"] = x
		params["top"] = y
		params["width"] = width
		params["height"] = height
		params["screen_mode"] = *screen_mode
		params["screen_rotate"] = *screen_rotate
		params["path"] = this.path()
		params["duration"] = *duration
	} else if (*libName) == "video" {
		params["zIndex"] = *zIndex
		params["left"] = x
		params["top"] = y
		params["width"] = width
		params["height"] = height
		params["screen_mode"] = *screen_mode
		params["screen_rotate"] = *screen_rotate
		params["path"] = this.path()
		params["offset"] = *offset
		params["timeout"] = *timeout
	} else if (*libName) == "pic" {
		params["zIndex"] = *zIndex
		params["left"] = x
		params["top"] = y
		params["width"] = width
		params["height"] = height
		params["screen_mode"] = *screen_mode
		params["screen_rotate"] = *screen_rotate
		params["path"] = this.path()
	} else if (*libName) == "gif" {
		params["zIndex"] = *zIndex
		params["left"] = x
		params["top"] = y
		params["width"] = width
		params["height"] = height
		params["screen_mode"] = *screen_mode
		params["screen_rotate"] = *screen_rotate
		params["path"] = this.path()
	} else if (*libName) == "qrcode" {
		params["zIndex"] = *zIndex
		params["left"] = x
		params["top"] = y
		params["width"] = width
		params["height"] = height
		params["screen_mode"] = *screen_mode
		params["screen_rotate"] = *screen_rotate
		params["content"] = *content
	} else if (*libName) == "camera" {
		params["zIndex"] = *zIndex
		params["left"] = x
		params["top"] = y
		params["width"] = width
		params["height"] = height
		params["screen_mode"] = *screen_mode
		params["screen_rotate"] = *screen_rotate
		params["device"] = *device
		params["camera_width"] = *camera_width
		params["camera_height"] = *camera_height
	} else if (*libName) == "text" {
		params["zIndex"] = *zIndex
		params["left"] = x
		params["top"] = y
		params["width"] = width
		params["height"] = height
		params["screen_mode"] = *screen_mode
		params["screen_rotate"] = *screen_rotate
		params["content"] = *content
		params["font_ttf"] = *font_ttf
		params["font_size"] = *font_size
		params["color"] = *color
		params["bgcolor"] = *bgcolor
		params["align"] = *align
		params["style"] = *style
	} else if (*libName) == "scroll" {
		params["zIndex"] = *zIndex
		params["left"] = x
		params["top"] = y
		params["width"] = width
		params["height"] = height
		params["screen_mode"] = *screen_mode
		params["screen_rotate"] = *screen_rotate
		params["content"] = *content
		params["font_ttf"] = *font_ttf
		params["font_size"] = *font_size
		params["color"] = *color
		params["bgcolor"] = *bgcolor
		params["style"] = *style
		params["orientation"] = *orientation
		params["speed"] = *speed
	} else if (*libName) == "datetime" {
		params["zIndex"] = *zIndex
		params["left"] = x
		params["top"] = y
		params["width"] = width
		params["height"] = height
		params["screen_mode"] = *screen_mode
		params["screen_rotate"] = *screen_rotate
		params["font_ttf"] = *font_ttf
		params["font_size"] = *font_size
		params["color"] = *color
		params["bgcolor"] = *bgcolor
		params["align"] = *align
		params["style"] = *style
	} else if (*libName) == "toast" {
		params["zIndex"] = *zIndex
		params["screen_mode"] = *screen_mode
		params["screen_rotate"] = *screen_rotate
		params["content"] = *content
		params["toast_type"] = *toast_type
		params["duration"] = *duration
	}
	data := make(map[string]interface{})
	data["id"] = fmt.Sprintf("PLAY_Z%d_%s_%d", *zIndex, strings.ToUpper(*libName), time.Now().Unix())
	data["type"] = "play"
	data["start"] = *start
	data["libName"] = *libName
	data["params"] = params
	return this.send(data)
}

func (this *XPlay) move() error {
	params := make(map[string]interface{})
	rs := strings.Split(*rect, ",")
	x, _ := strconv.Atoi(rs[0])
	y, _ := strconv.Atoi(rs[1])
	width, _ := strconv.Atoi(rs[2])
	height, _ := strconv.Atoi(rs[3])
	params["zIndex"] = *zIndex
	params["left"] = x
	params["top"] = y
	params["width"] = width
	params["height"] = height
	data := make(map[string]interface{})
	data["type"] = "move"
	data["params"] = params
	return this.send(data)
}

func (this *XPlay) change() error {
	params := make(map[string]interface{})
	params["zIndex"] = *zIndex
	params["newIndex"] = *newIndex
	data := make(map[string]interface{})
	data["type"] = "change"
	data["params"] = params
	return this.send(data)
}

func (this *XPlay) snap() error {
	params := make(map[string]interface{})
	params["path"] = *path
	data := make(map[string]interface{})
	data["type"] = "snap"
	data["params"] = params
	return this.send(data)
}

func (this *XPlay) query() error {
	data := make(map[string]interface{})
	data["type"] = "query"
	return this.send(data)
}

// TODO 暂时只支持视频 ...
func (this *XPlay) path() string {
	// 路径不能为空 ...
	if *path == "" {
		log.Fatal("not found path file ...")
	}
	// 本机播放 ...
	if *addr == "127.0.0.1:8700" {
		// 当前路径文件 ...
		if strings.HasPrefix(*path, "./") {
			dir, _ := os.Getwd()
			ph := dir + string(os.PathSeparator) + (*path)[2:]
			return ph
		}
	} else { // 远端播放 ...
		// 当前路径文件 ...
		if strings.HasPrefix(*path, "./") {
			go func() {
				wg.Add(1) // 添加计数
				// 启动文件服务 ...
				dir, _ := os.Getwd()
				http.Handle("/xplay/", http.StripPrefix("/xplay/", http.FileServer(http.Dir(dir))))
				log.Fatal(http.ListenAndServe(":8711", nil))
				wg.Done() // 移除计数
			}()
			return "http://" + this.ipaddr() + ":8711/xplay/" + (*path)[2:]
		}
	}
	return *path
}

func (this *XPlay) ipaddr() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, address := range addrs {
		// 检查IP地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip := ipnet.IP.String()
				a1 := strings.Split(*addr, ":")
				if len(a1) == 2 {
					a2 := strings.Split(a1[0], ".")
					if len(a2) == 4 {
						a3 := a2[0] + "." + a2[1] + "." + a2[2]
						if strings.HasPrefix(ip, a3) {
							return ip
						}
					}
				}
			}
		}
	}
	return ""
}

func main() {

	// 设置一下日志的结构
	log.SetFlags(log.Lshortfile | log.Ltime | log.Lmicroseconds)

	// 解析程序参数
	flag.Parse()

	// New XPlay
	xplay := NewXPlay()

	// 连接到 XPlay
	if err := xplay.connect(*addr); err != nil {
		// 连接异常退出 ...
		log.Fatal(err.Error())
	} else {
		// 结束时断开连接
		defer xplay.conn.Close()
	}

	// Error
	var err error

	// 模式
	if *play { // Play
		err = xplay.play()
	} else if *move { // Move
		err = xplay.move()
	} else if *change { // Change
		err = xplay.change()
	} else if *snap { // Snap
		err = xplay.snap()
	} else if *stop { // Stop
		if *all { // Stop ALL
			err = xplay.stop_all()
		} else if *ids != "" { // Stop IDS
			err = xplay.stop(*ids)
		}
	} else if *query { // Query
		xplay.query()
	}

	// Print Error
	if err == nil {
		// 接收返回
		xplay.result()
	} else {
		// 输出异常 ...
		log.Fatal(err.Error())
	}

	// 等待计数结束
	log.Println("正在检查推流状态 ...")
	log.Println("等待推流结束 ...")
	log.Println("退出(Ctrl+C)")
	wg.Wait()
	log.Println("完成")

}

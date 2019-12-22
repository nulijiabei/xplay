package main

// 2019/11/18 v1.3

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

// Connect Address
var addr = flag.String("addr", "127.0.0.1:8700", "xplay tcp address")

// Default
var id = flag.String("id", "", "debug id")
var start = flag.Int64("start", -1, "start time millisecond")
var libName = flag.String("libName", "", "video、pic、camera、gif、qrcode、text、scroll、background")

// Stop
var stop = flag.Bool("stop", false, "stop or play")
var all = flag.Bool("all", false, "stop all")
var ids = flag.String("ids", "", "stop ids")

// Play
var play = flag.Bool("play", false, "play or stop")

// Play Params
var zIndex = flag.Int("zIndex", 10, "1-9999")
var rect = flag.String("rect", "0,0,1920,1080", "")
var screen_mode = flag.String("screen_mode", "landscape", "landscape、portrait")
var screen_rotate = flag.Int("screen_rotate", 0, "landscape: 0、180 or portrait: 90、270")
var path = flag.String("path", "", "file path")
var content = flag.String("content", "", "data content")

// Play Camera
var device = flag.String("device", "/dev/video0", "camera device")
var camera_width = flag.Int("camera_width", 1280, "camera video width")
var camera_height = flag.Int("camera_height", 720, "camera video height")

// Play Text
var font_ttf = flag.String("font_ttf", "", "TrueTypeFont")
var font_size = flag.Int("font_size", 18, "")
var color = flag.String("color", "rgba(0,128,0,100%)", "")
var bgcolor = flag.String("bgcolor", "rgba(0,0,0,20%)", "")
var align = flag.String("align", "center", "center、right、left")
var style = flag.String("style", "normal", "normal、bold、italic、underline、strikethrough")

// Play Scroll
var speed = flag.Int("speed", 1, "move pixel / frame")
var orientation = flag.String("orientation", "horizontal", "horizontal | vertical")

// Play Toast
var toast_type = flag.String("toast_type", "notice", "notice、success、warning、error")
var duration = flag.Int("duration", 0, "timeout stop toast")

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

// 发送 to XPlay
func (this *XPlay) send(_data map[string]interface{}) error {
	js, _ := json.MarshalIndent(_data, "", "  ")
	fmt.Println(string(js))
	_, err := this.conn.Write(append(js, []byte("\n#End\n")...))
	return err
}

// 停止全部
func (this *XPlay) stop_all() error {
	data := make(map[string]interface{})
	data["type"] = "stop"
	data["params"] = map[string]bool{"all": true}
	return this.send(data)
}

// 停止指定层
func (this *XPlay) stop(_ids string) error {
	vs := strings.Split(_ids, ",")
	data := make(map[string]interface{})
	params := make(map[string]interface{})
	params["ids"] = vs
	data["type"] = "stop"
	data["params"] = params
	return this.send(data)
}

func (this *XPlay) params(_zIndex int, _x, _y, _width, _height int, _screen_mode string, _screen_rotate int) map[string]interface{} {
	params := make(map[string]interface{})
	params["zIndex"] = _zIndex
	params["top"] = _y
	params["left"] = _x
	params["width"] = _width
	params["height"] = _height
	params["screen_mode"] = _screen_mode
	params["screen_rotate"] = _screen_rotate
	return params
}

func (this *XPlay) play() error {
	var params map[string]interface{}
	rs := strings.Split(*rect, ",")
	x, _ := strconv.Atoi(rs[0])
	y, _ := strconv.Atoi(rs[1])
	width, _ := strconv.Atoi(rs[2])
	height, _ := strconv.Atoi(rs[3])
	params = this.params(*zIndex, x, y, width, height, *screen_mode, *screen_rotate)
	if (*libName) == "video" || (*libName) == "pic" || (*libName) == "gif" {
		params["path"] = *path
	} else if (*libName) == "camera" {
		params["device"] = *device
		params["camera_width"] = *camera_width
		params["camera_height"] = *camera_height
	} else if (*libName) == "qrcode" {
		params["content"] = *content
	} else if (*libName) == "text" {
		params["content"] = *content
		params["font_ttf"] = *font_ttf
		params["font_size"] = *font_size
		params["color"] = *color
		params["bgcolor"] = *bgcolor
		params["align"] = *align
		params["style"] = *style
	} else if (*libName) == "scroll" {
		params["content"] = *content
		params["font_ttf"] = *font_ttf
		params["font_size"] = *font_size
		params["color"] = *color
		params["style"] = *style
		params["orientation"] = *orientation
		params["speed"] = *speed
	} else if (*libName) == "datetime" {
		params["font_ttf"] = *font_ttf
		params["font_size"] = *font_size
		params["color"] = *color
		params["bgcolor"] = *bgcolor
		params["align"] = *align
		params["style"] = *style
	} else if (*libName) == "toast" {
		params["content"] = *content
		params["toast_type"] = *toast_type
		params["duration"] = *duration
	} else if (*libName) == "background" {
		params["bgcolor"] = *bgcolor
	}
	data := make(map[string]interface{})
	if *id == "" {
		data["id"] = fmt.Sprintf("PLAY_Z%d_%s_%d", *zIndex, strings.ToUpper(*libName), time.Now().Unix())
	} else {
		data["id"] = *id
	}
	data["type"] = "play"
	data["start"] = *start
	data["libName"] = *libName
	data["params"] = params
	return this.send(data)
}

func main() {

	// 设置一下日志的结构
	log.SetFlags(log.Lshortfile | log.Ltime | log.Lmicroseconds)

	// 解析程序参数
	flag.Parse()

	// New XPlay
	xplay := NewXPlay()
	if err := xplay.connect(*addr); err != nil {
		log.Fatal(err.Error())
	} else {
		defer xplay.conn.Close()
	}

	// 模式
	var err error
	if *play { // Play
		err = xplay.play()
	} else if *stop { // Stop
		if *all { // Stop ALL
			err = xplay.stop_all()
		} else if *ids != "" { // Stop IDS
			err = xplay.stop(*ids)
		}
	}
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("Success")
	}
}

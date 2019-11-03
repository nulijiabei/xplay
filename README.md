# xplay

专为树莓派(Raspberry Pi)设计的多媒体播放器

支持(视频、音频、流媒体、图片、摄像头、动画、文本、滚动字幕、二维码)

---

![image](images/%E6%A8%AA%E7%AB%96%E5%B1%8F.jpg)
![image](images/%E5%A4%9A%E5%88%86%E5%B1%8F.gif)

---

*注意：这个开源项目免费提供 xplay 的控制端源码(暂时支持：python、golang)，且免费提供 xplay 程序 ...*

---
### 支持硬件

| 硬件 | 分辨率 | FPS | 测试 |
| --- | --- | --- | --- |
| Raspberry Pi Zero  | 1080p/v720p  | 30 | 已测试 |
| Raspberry Pi 3A+   | 1080p/v720p  | 30 | 已测试 |
| Raspberry Pi 3B+   | 1080p/v720p  | 30 | 已测试 |
| Raspberry Pi 2B    | 1080p/v720p  | 30 | 未测试 |
| Raspberry Pi 3B    | 1080p/v720p  | 30 | 未测试 |
| Raspberry Pi 4B    | GL1 Mode    | 无 | 待适配 |

***理论上支持所有树莓派，但是经过测试的树莓派只有上面几款***

### 支持系统

| 系统 | 日期 | 安装 | 测试 |
| --- | --- | --- | --- |
| Raspbian Buster with desktop | 2019-09-26 | buster/install | 已测试 |
| Raspbian Buster Lite | 2019-09-26 | buster-lite/install | 已测试 |

---
### 支持功能

1. 支持自定义播放器分辨率、帧率(FPS)、音频采样率(Sample Rate)
2. 支持使用(TCP)连接播放器发送指令控制(播放、覆盖、停止、等 ...)
3. 支持(视频、音频、流媒体、图片、摄像头、动画、文本、滚动字幕、二维码)素材播放
4. 支持(视频)多种格式(例如：MP4、AVI、MOV、等 ...)基于H264的视频编码，音频(AAC)
5. 支持(图片)JPG与PNG格式
6. 支持(动画)GIF格式
7. 支持(视频)硬解播放
8. 支持(视频)预加载
9. 支持(摄像头)设备(Raspberry Pi Camera V2)
10. 支持(视频、图片)无黑场切换播放
11. 支持(视频)音频同步(视频帧时间戳与音轨帧时间戳)播放
12. 支持(视频、流媒体、图片、摄像头、动画、文本、滚动字幕、二维码)多层(Overlay)播放
13. 支持(文本)自定义(字体大小、字体颜色、背景颜色、透明度、对齐方式、风格样式、多行)
14. 支持(滚动字幕)自定义(字体大小、字体颜色、透明度、风格样式、移动速度)方向(从右向左)
15. 支持自定义布局(通过多层功能可以实现多种自定义布局)
16. 支持自定义(视频)是否循环播放(视频在播放到结尾时是否停留在最后一帧)
17. 支持自定义素材尺寸(width，height)，任意拉伸缩放素材尺寸播放
18. 支持自定义素材位置(x，y)播放，任意定义素材播放位置
19. 支持自定义素材横竖屏播放
20. 支持自定义素材开始播放时间(多个播放器间可以实现同步播放)
21. 支持静音播放

---
### 安装方法

 ***1. 显存设置***
 
 | 文件 | 设置 |
 | --- | --- |
 | /boot/config.txt | gpu_mem=192 (显存内存分配不能少于192M)(重启生效) |
 
 ***2. 安装程序***
 ```
 git clone https://github.com/nulijiabei/xplay.git
 cd xplay/
 cd buster/ # Raspbian Buster with desktop 或者 cd buster-lite/ # Raspbian Buster Lite
 sh -x install
 ```
 
 ***3. 运行程序***
 
 | 参数 | 默认 | 说明 |
 | --- | --- | --- |
 | -R x,y,width,height | 0,0,1920,1080 | 分辨率设置 |
 | -fps n | 30 | 帧率设置 |
 | -noloop | 无 | 视频在播放到结尾时是否停留在最后一帧 |
 | -mute | 无 | 静音设置 |
 | -sample_rate n | 44100 | 音频采样频率设置 |
 | -cpuid | 无 | 显示CPUID |
 
 ```
 // 分辨率 1920x1080 【默认】
 // 帧率 30fps 【默认】
 // 音频采样率 44100 【默认】
 // 视频循环播放【默认】
 // 播放声音【默认】
 /usr/bin/xplay
 
 // 分辨率设置 1280x720 
 // 音频采样率设置 48000 
 // 视频循环设置 视频在播放到结尾时停留在最后一帧
 // 【未设置则使用默认设置】
 /usr/bin/xplay -R 0,0,1280,720 -sample_rate 48000 -noloop
 ```
 
 ***4. 播放测试***
 
 ```
 // 测试工具
 apt-get install telnet
 ```
 
 ```
 // 测试指令
 root@raspberrypi:~# telnet 127.0.0.1 8700
 Trying 127.0.0.1...
 Connected to 127.0.0.1.
 Escape character is '^]'.
 
 // 指令内容
 {
    "libName": "video",
    "params": {
        "path": "/root/sample.mp4",
        "width": 1920,
        "height": 1080,
        "zIndex": 10
    },
    "type": "play"
 }
 #End
 ```
 
 ***5. 查看日志***
 ```
 tail -n 100 /dev/shm/xplay.log
 ```

---
### 播放控制

 * 通过Socket接口与播放器建立TCP连接，向播放器发送指令，从播放器接收返回
 * 可以通过多种支持Socket连接语言开发控制程序(java、python、C++、golang、等 ...)
 * 注：运行样例之前请先启动xplay主程序(/usr/bin/xplay &)
 
 [Golang 控制样例](example/example.go)  
 [Python 控制样例](example/example.py)
 
---
### 控制指令
  
 | 指令类型(type) | 说明 |
 | --- | --- | 
 | play | 播放 |
 | stop | 停止 |
 
 | 素材类型(libName) | 说明 |
 | --- | --- |
 | video | 视频/流媒体 |
 | camera | 摄像头 |
 | pic | 图片 |
 | gif | 动画 |
 | qrcode | 二维码 |
 | text | 文本 |
 | scroll | 滚动字幕 |
 | background | 背景 |
 
 ```
 // 指令说明
 {
     "id": "Z10_Play_1557737960000", // 唯一标记(自定义唯一标识)【非必填】
     "type": "play",                 // 指令类型【必填】
     "start": -1,                    // 开始时间(默认：-1，立即播放，本地毫秒时间戳)【非必填】
     "libName": "video",             // 素材类型(video、pic、camera、gif、qrcode、text、scroll）【必填】
     "params": {                     // 参数集合【必填】
         "zIndex": 10,               // 层(支持多层播放，层数越小画面越靠前)【必填】
         "path": "/root/sample.mp4", // 素材路径【必填】
         "left": 0,                  // 距左像素(X轴)【非必填】
         "top": 0,                   // 距顶像素(Y轴)【非必填】
         "width": 1920,              // 宽(素材显示的宽，非素材原始尺寸，支持缩放拉伸)【必填】
         "height": 1080,             // 高(素材显示的高，非素材原始尺寸，支持缩放拉伸)【必填】
         "screen_mode": "landscape", // 屏幕模式(横屏：landscape，竖屏：portrait，默认横屏)【非必填】
         "screen_rotate": 0          // 旋转角度(横屏角度：0、180，竖屏角度：90、270，默认横屏)【非必填】
     }
 }
 #End                                // 指令结束【必填】
 ```
 
 ```
  // 显示视频
  {
     "id": "Z10_Play_1557737960000",
     "type": "play",
     "start": -1,
     "libName": "video",
     "params": {
         "zIndex": 10,
         "path": "/root/sample.mp4",
         "left": 0,
         "top": 0,
         "width": 1920,
         "height": 1080,
         "screen_mode": "landscape",
         "screen_rotate": 0
     }
 }
 #End
 ```
 
 ```
 // 显示流媒体(暂时通过测试的只有http流)
 {
    "id": "Z10_Play_1557737960000",
    "libName": "video",
    "start": -1,
    "type": "play",
    "params": {
        "path": "http://ivi.bupt.edu.cn/hls/cctv1hd.m3u8",
        "height": 1080,
        "width": 1920,
        "left": 0,
        "top": 0,
        "screen_mode": "landscape",
        "screen_rotate": 0,
        "zIndex": 10
     }
 }
 ```
 
 ```
 // 显示图片
 {
    "id": "Z10_Play_1556637960000",
    "type": "play",
    "start": -1,
    "libName": "pic",
    "params": {
        "zIndex": 10,
        "path": "/root/cat.jpg",
        "top": 0,
        "left": 0,
        "height": 1080,
        "width": 1920,
        "screen_mode": "landscape",
        "screen_rotate": 0
    }
 }
 #End
 ```
 
 | 摄像头(camera) | 值 | 说明 |
 | --- | --- | --- |
 | device | /dev/video0 | 设备地址 |
 | camera_width | 1280 | 摄像头 video_size 宽 |
 | camera_height | 720 | 摄像头 video_size 高 |
  
 >启用摄像头 raspi-config > Interfacing Options > P1 Camera 

 ```
 // 显示摄像头
 {
    "id": "Z9_Play_1558837960000",
    "type": "play",
    "start": -1,
    "libName": "camera",
    "params": {
        "zIndex": 9,
        "device": "/dev/video0",
        "camera_width": 1280,
        "camera_height": 720,
        "top": 0,
        "left": 0,
        "width": 1920,
        "height": 1080,
        "screen_mode": "landscape",
        "screen_rotate": 0
    }
 }
 #End
 ```
 
 | 文本(text) | 值 | 说明 |
 | --- | --- | --- |
 | color     | rgba(0, 128, 0, 100%) | 文本颜色及透明度 |
 | bgcolor   | rgba(0, 0, 0, 0%)     | 背景颜色及透明度 |
 | font_size | 14                    | 字体大小 |
 | align     | center、right、left   | 对齐方式 |
 | style     | normal、bold、italic、underline、strikethrough | 文本样式 |
 | content   | \t\t123,456,789.\n987,654,321.\n | 文本内容 |
 
 ```
 // 显示文本
 {
    "type": "play",
    "id": "Z9_Play_1572344489512",
    "libName": "text",
    "start": -1,
    "params": {
       "top": 0,
       "left": 0,
       "width": 1280,
       "height": 50,
       "zIndex": 9,
       "screen_rotate": 0,
       "screen_mode": "landscape",
       "content": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
       "color": "rgba(255, 0, 0, 100%)",
       "bgcolor": "rgba(0, 0, 0, 0%)",
       "font_size": 30,
       "align": "center",
       "style": "bold"
    }
 }
 #End
 ```
 
 | 滚动字幕(scroll) | 值 | 说明 |
 | --- | --- | --- |
 | color     | rgba(0, 128, 0, 100%) | 文本颜色及透明度 |
 | font_size | 14 | 字体大小 |
 | style     | normal、bold、italic、underline、strikethrough | 文本样式 |
 | speed     | 1 | 移动速度(每帧移动N像素) |
 | content   | ABCDEFGHIJKLMNOPQRSTUVWXYZ | 文本内容 |      
 
 ```
 // 显示滚动字幕
 {
    "type": "play",
    "id": "Z9_Play_1572344489512",
    "libName": "scroll",
    "start": -1,
    "params": {
       "top": 0,
       "height": 50,
       "zIndex": 5,
       "screen_rotate": 0,
       "screen_mode": "landscape",
       "content": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
       "color": "rgba(255, 0, 0, 100%)",
       "font_size": 30,
       "style": "bold",
       "speed": 1
    }
 }
 #End
 ```
 
 | 背景(background) | 值 | 说明 |
 | --- | --- | --- |
 | bgcolor | rgba(0, 0, 0, 20%) | 背景颜色及透明度 |
 
 >例如：可以通过背景层为跑马字幕提供半透明的背景 ...
 
 ```
 // 显示背景
 {
    "type": "play",
    "id": "Z6_Play_1572344489512",
    "libName": "background",
    "start": -1,
    "params": {
       "top": 0,
       "left": 0,
       "width": 1920,
       "height": 50,
       "zIndex": 6,
       "screen_rotate": 0,
       "screen_mode": "landscape",
       "bgcolor": "rgba(0, 0, 0, 10%)"
    }
 }
 #End
 ```
 
 >如果动画显示不正常, 您可以使用Photoshop重新存储一下
 
 ```
 // 显示动画
 {
    "id": "Z9_Play_1554437960000",
    "type": "play",
    "start": -1,
    "libName": "gif",
    "params": {
        "zIndex": 9,
        "path": "/root/aaa.gif",
        "top": 0,
        "left": 0,
        "height": 100,
        "width": 100,
        "screen_mode": "landscape",
        "screen_rotate": 0
    }
 }
 #End
 ```
 
 ```
 // 显示二维码
 {
    "id": "Z9_Play_1558837960000",
    "type": "play",
    "start": -1,
    "libName": "qrcode",
    "params": {
        "zIndex": 9,
        "content": "http://www.danoonetwork.com",
        "top": 0,
        "left": 0,
        "height": 100,
        "width": 100,
        "screen_mode": "landscape",
        "screen_rotate": 0
    }
 }
 #End
 ```
 
 ```
 // 停止全部层
 {
   "type": "stop",
   "params": {
       "all": true
   }
 }
 #End
 ```
 
 ```
 // 停止指定层
 {
   "type": "stop",
   "params": {
       "ids": ["10", "11", "12"]
   }
 }
 #End
 ```

---
### 更多功能

 * 基于云服务的信息发布系统(www.danoonetworks.com)
 * 提供一站式解决方案(软硬件、一体机、4G/WIFI网络、等 ...)

![image](images/cherry.jpg)

 ---


# xplay

专为树莓派(Raspberry Pi)设计的多媒体播放器，支持无黑场切换(视频、音频、图片、文字、动画、二维码)

---

![image](https://github.com/nulijiabei/xplay/blob/master/images/%E6%A8%AA%E7%AB%96%E5%B1%8F.jpg)
![image](https://github.com/nulijiabei/xplay/blob/master/images/%E5%A4%9A%E5%88%86%E5%B1%8F.gif)

---

*注意：这并不是一个开源项目，只是免费提供已经编译好的多媒体播放程序 ...*

---
### 支持硬件

| 硬件 | 分辨率 | FPS | 测试 |
| --- | --- | --- | --- |
| Raspberry Pi Zero  | 720p          | 25 | 已测试 |
| Raspberry Pi 3A+   | 1080p/720p    | 25 | 已测试 |
| Raspberry Pi 3B+   | 1080p/720p    | 25 | 已测试 |
| Raspberry Pi 4B    | 1080p         | 30 | 待测试 |
| Rockchip           | 1080p         | 30 | 可定制 |
| Linux Intel/NVIDIA | 1080p/4K      | 30/60 | 可定制 |
| Windows 7/10       | 1080p/4K      | 30/60 | 可定制 |
| 其它               | 1080p/4K      | 30/60 | 可定制 |

### 支持系统

| 系统 | 日期 | 安装 | 测试 |
| --- | --- | --- | --- |
| Raspbian Buster with desktop | 2019-09-26 | buster/install | 已测试 |

---
### 支持功能

1. 支持自定义播放器分辨率、帧率(FPS)、音频采样率(Sample Rate)
2. 支持使用(TCP)连接播放器发送指令控制(播放、覆盖、停止、等 ...)
3. 支持(视频、音频、图片、文字、动画、二维码)素材播放
4. 支持(视频)多种格式(例如：MP4、AVI、MOV、等 ...)基于H264的视频编码，音频(AAC)
5. 支持(图片)JPG与PNG格式
6. 支持(动画)GIF格式
7. 支持(视频)硬解播放
8. 支持(视频)预加载
9. 支持(视频、图片)无黑场切换播放
10. 支持(视频)音频同步(视频帧时间戳与音轨帧时间戳)播放
11. 支持(视频、图片、文字、动画、二维码)多层(Overlay)播放
12. 支持自定义布局(通过多层功能可以实现多种自定义布局)
13. 支持自定义(视频)是否循环播放(视频在播放到结尾时是否停留在最后一帧)
14. 支持自定义素材尺寸(Width，Height)，任意拉伸缩放素材尺寸播放
15. 支持自定义素材位置(X，Y)播放，任意定义素材播放位置
16. 支持自定义素材横竖屏播放
17. 支持自定义素材开始播放时间(多个播放器间可以实现同步播放)
18. 支持静音播放

---
### 安装方法

 ***1. 设置源***
 
 | 文件 | 地址 |
 | --- | --- |
 | /etc/apt/sources.list | https://mirrors4.tuna.tsinghua.edu.cn/raspbian/raspbian/
 | /etc/apt/sources.list.d/raspi.list | https://mirrors4.tuna.tsinghua.edu.cn/raspberrypi/

 ***2. 更新源***
  ```
 apt-get update
 ```
 
 ***3. 显存设置***
 
 | 文件 | 设置 |
 | --- | --- |
 | /boot/config.txt | gpu_mem=192 (显存内存分配不能少于192M) |
 
 ***4. 安装程序***
 ```
 git clone https://github.com/nulijiabei/xplay.git
 cd xplay/
 cd buster/ # Raspbian Buster with desktop
 sh -x install
 ```
 
 ***5. 运行程序***
 
 | 参数 | 默认 |说明 |
 | --- | --- | --- |
 | -R | 0,0,1920,1080 | 分辨率设置 |
 | -fps | 25 | 帧率设置 |
 | -noloop | 循环播放 | 视频在播放到结尾时是否停留在最后一帧 |
 | -mute | 播放声音 | 静音设置 |
 | -sample_rate | 44100 | 音频采样频率设置 |
 
 ```
 // 分辨率 1920x1080 【默认】
 // 帧率 25fps 【默认】
 // 音频采样率 44100 【默认】
 // 视频循环播放【默认】
 // 播放声音【默认】
 xplay
 // 分辨率设置 1280x720 
 // 音频采样率设置 48000 
 // 视频循环设置 视频在播放到结尾时停留在最后一帧
 // 【未设置则使用默认设置】
 xplay -R 0,0,1280,720 -sample_rate 48000 -noloop
 ```
 
 ***6. 播放测试***
 
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
 
 ```
 // 指令说明
 {
     "libName": "video",             // 类型
     "params": {                     // 参数
         "path": "/root/sample.mp4", // 路径
         "width": 1920,              // 宽
         "height": 1080,             // 高
         "zIndex": 10                // 层
     },
     "type": "play"                  // 播放
 }
 #End                                // 结束
 ```
 
 ***7. 查看日志***
 ```
 tail -n 100 /dev/shm/xplay.log
 ```

---
### 连接控制

 * 通过Socket接口与播放器建立TCP连接，向播放器发送指令，从播放器接收返回
 * 可以通过多种支持Socket连接语言开发控制程序(java、python、C++、golang、等 ...)
 
 ***连接样例***
 ```
 例：XXX
 ```
 
---
### 控制指令

 * XXX
 
---
### 系统优化

 * Raspbian 优化官方系统，使多媒体播放器更加稳定的长期运行 ...
 
---
### 更多功能

 * 基于云服务的信息发布系统(www.danoonetworks.com)
 * 提供一站式解决方案(软硬件、一体机、4G/WIFI网络、等 ...)

![image](https://github.com/nulijiabei/xplay/blob/master/images/cherry.jpg)

 ---


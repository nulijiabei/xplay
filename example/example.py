#!/usr/bin/python
# -*- coding: UTF-8 -*-

import socket
import json
import time
import os

def connect():
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect(('127.0.0.1', 8700))
    return client

def stop(client):
    data = {}
    params = {}
    params['all'] = True
    data['type'] = "stop"
    data['params'] = params
    js = json.dumps(data)
    client.send((js + '\n#End\n').encode('utf-8'))
    
def play(client, libName, zIndex, path, top, left, width, height, screenMode, screenRotate):
    data = {}
    params = {}
    params['zIndex'] = zIndex
    params['path'] = path
    params['top'] = top
    params['left'] = left
    params['width'] = width
    params['height'] = height
    params['screen_mode'] = screenMode
    params['screen_rotate'] = screenRotate
    data['id'] = ("PLAY_VIDEO_%d") % (int(round(time.time() * 1000)))
    data['type'] = "play"
    data['libName'] = libName
    data['params'] = params
    js = json.dumps(data)
    client.send((js + '\n#End\n').encode('utf-8'))

def getFiles(str):
    fs = []
    for root,dirs,files in os.walk(str):
        for file in files:
            fs.append(str + '/' + file)
    return fs

if __name__ == '__main__':

    # 建立连接
    client = connect()

    # 停止全部
    stop(client)

    # 循环播放
    while True:

        # 获取目录素材文件
        fs = getFiles(r'E:\test')
        if len(fs) == 0 :
            time.sleep(5)
            continue

        # 遍历素材
        for f in fs:
            # 是否为视频
            if f.endswith(".mp4") :
                # 播放视频
                play(client, "video", 10, f, 0, 0, 1920, 1080, "landscape", 0)
            # 是否为图片
            if f.endswith(".jpg") or f.endswith(".png") :
                # 播放图片
                play(client, "pic", 10, f, 0, 0, 1920, 1080, "landscape", 0)
            # 间隔5s切换
            time.sleep(5)

    # 断开连接
    client.close()


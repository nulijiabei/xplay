#!/bin/bash

CONTENT=/root/content # 素材全部放到这个目录中 
ZINDEX=10 # 使用层
RECT="0,0,1920,1080" # 显示区域与位置

while(true);do
    for v in `ls ${CONTENT}`;do
        if [[ $v == *.mp4 ]];then
            # 视频预加载才可以无黑场, 详情 https://gitee.com/nljb/xplay/wikis
            echo Play VIDEO ${ZINDEX} ${RECT} "${CONTENT}/${v}"
	    /usr/bin/xplayctl -play -libName video -zIndex ${ZINDEX} -rect ${RECT} -path "${CONTENT}/${v}"
	    sleep 5 # 可以按照视频的实际时间进行SLEEP ...
        elif [[ $v == *.jpg ]] || [[ $v == *.png ]];then
	    echo Play PIC ${ZINDEX} ${RECT} "${CONTENT}/${v}"
	    /usr/bin/xplayctl -play -libName pic -zIndex ${ZINDEX} -rect ${RECT} -path "${CONTENT}/${v}"
	    sleep 5 # 可以按照图片的时间间隔进行SLEEP ...
	fi
    done
done

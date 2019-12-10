#!/bin/bash

CONTENT=/root/content # 最少放4个素材
ZINDEX_1=10 # 位置 1 使用层
ZINDEX_2=11 # 位置 2 使用层
ZINDEX_3=12 # 位置 3 使用层
ZINDEX_4=13 # 位置 4 使用层
RECT_1="0,0,960,540"     # 位置 1 显示区域与位置
RECT_2="960,0,960,540"   # 位置 2 显示区域与位置
RECT_3="0,540,960,540"   # 位置 3 显示区域与位置
RECT_4="960,540,960,540" # 位置 4 显示区域与位置

index=0
while(true);do
    for v in `ls ${CONTENT}`;do
        if [[ $v == *.jpg ]] || [[ $v == *.png ]];then
            if [ ${index} -eq 0 ];then
	        echo Play PIC ${ZINDEX_1} ${RECT_1} "${CONTENT}/${v}"
	        /usr/bin/xplayctl -play -libName pic -zIndex ${ZINDEX_1} -rect ${RECT_1} -path "${CONTENT}/${v}"
                index=$(($index+1))
		continue
	    fi
            if [ ${index} -eq 1 ];then
	        echo Play PIC ${ZINDEX_2} ${RECT_2} "${CONTENT}/${v}"
	        /usr/bin/xplayctl -play -libName pic -zIndex ${ZINDEX_2} -rect ${RECT_2} -path "${CONTENT}/${v}"
                index=$(($index+1))
		continue
	    fi
            if [ ${index} -eq 2 ];then
	        echo Play PIC ${ZINDEX_3} ${RECT_3} "${CONTENT}/${v}"
	        /usr/bin/xplayctl -play -libName pic -zIndex ${ZINDEX_3} -rect ${RECT_3} -path "${CONTENT}/${v}"
                index=$(($index+1))
		continue
	    fi
            if [ ${index} -eq 3 ];then
	        echo Play PIC ${ZINDEX_4} ${RECT_4} "${CONTENT}/${v}"
	        /usr/bin/xplayctl -play -libName pic -zIndex ${ZINDEX_4} -rect ${RECT_4} -path "${CONTENT}/${v}"
                index=0
   	        sleep 5
		continue
	    fi
	fi 
    done
done

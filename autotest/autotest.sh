#!/bin/bash

/usr/bin/xplayctl -stop -all

sleep 1

/usr/bin/xplayctl -play -libName pic -zIndex 10 -screen_mode landscape -screen_rotate 0 -rect 0,0,1920,1080 -path `pwd`/a1_landscape.jpg
sleep 3
/usr/bin/xplayctl -play -libName pic -zIndex 10 -screen_mode landscape -screen_rotate 180 -rect 0,0,1920,1080 -path `pwd`/a1_landscape.jpg
sleep 3
/usr/bin/xplayctl -play -libName pic -zIndex 10 -screen_mode portrait -screen_rotate 90 -rect 0,0,1080,1920 -path `pwd`/a1_portrait.jpg
sleep 3
/usr/bin/xplayctl -play -libName pic -zIndex 10 -screen_mode portrait -screen_rotate 270 -rect 0,0,1080,1920 -path `pwd`/a1_portrait.jpg

sleep 3

/usr/bin/xplayctl -play -libName video -zIndex 10 -screen_mode landscape -screen_rotate 0 -rect 0,0,1920,1080 -path `pwd`/yiyezi.mp4
sleep 3
/usr/bin/xplayctl -play -libName video -zIndex 10 -screen_mode landscape -screen_rotate 180 -rect 0,0,1920,1080 -path `pwd`/yiyezi.mp4

sleep 3

/usr/bin/xplayctl -play -libName pic -zIndex 9 -screen_mode landscape -screen_rotate 0 -rect 100,100,800,600 -path `pwd`/a2.jpg
sleep 3
/usr/bin/xplayctl -play -libName pic -zIndex 9 -screen_mode landscape -screen_rotate 180 -rect 100,100,800,600 -path `pwd`/a2.jpg
sleep 3
/usr/bin/xplayctl -play -libName pic -zIndex 9 -screen_mode portrait -screen_rotate 90 -rect 100,100,800,600 -path `pwd`/a2.jpg
sleep 3
/usr/bin/xplayctl -play -libName pic -zIndex 9 -screen_mode portrait -screen_rotate 270 -rect 100,100,800,600 -path `pwd`/a2.jpg

sleep 3

/usr/bin/xplayctl -play -libName text -zIndex 8 -screen_mode landscape -screen_rotate 0 -rect "0,0,1920,50" -font_size 20 -color "rgba(255,0,0,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统设计的多媒体播放器" -align center -style bold
sleep 3
/usr/bin/xplayctl -play -libName text -zIndex 8 -screen_mode landscape -screen_rotate 180 -rect "0,0,1920,50" -font_size 20 -color "rgba(0,255,0,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统设计的多媒体播放器" -align center -style italic 
sleep 3
/usr/bin/xplayctl -play -libName text -zIndex 8 -screen_mode portrait -screen_rotate 90 -rect "0,0,1080,50" -font_size 20 -color "rgba(0,0,255,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统设计的多媒体播放器" -align center -style underline
sleep 3
/usr/bin/xplayctl -play -libName text -zIndex 8 -screen_mode portrait -screen_rotate 270 -rect "0,0,1080,50" -font_size 20 -color "rgba(255,255,255,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统设计的多媒体播放器" -align center -style strikethrough

sleep 3

/usr/bin/xplayctl -play -libName scroll -zIndex 8 -screen_mode landscape -screen_rotate 0 -rect "0,0,-1,50" -font_size 20 -color "rgba(255,0,0,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统设计的多媒体播放器" -style bold -orientation horizontal -speed 2
sleep 3
/usr/bin/xplayctl -play -libName scroll -zIndex 8 -screen_mode landscape -screen_rotate 180 -rect "0,0,-1,50" -font_size 20 -color "rgba(255,0,0,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统设计的多媒体播放器" -style bold -orientation horizontal -speed 2
sleep 3
/usr/bin/xplayctl -play -libName scroll -zIndex 8 -screen_mode portrait -screen_rotate 90 -rect "0,0,-1,50" -font_size 20 -color "rgba(255,0,0,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统设计的多媒体播放器" -style bold -orientation horizontal -speed 2
sleep 3
/usr/bin/xplayctl -play -libName scroll -zIndex 8 -screen_mode portrait -screen_rotate 270 -rect "0,0,-1,50" -font_size 20 -color "rgba(255,0,0,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统设计的多媒体播放器" -style bold -orientation horizontal -speed 2

sleep 3

/usr/bin/xplayctl -play -libName scroll -zIndex 8 -screen_mode landscape -screen_rotate 0 -rect "710,0,500,-1" -font_size 20 -color "rgba(255,0,0,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统 设计的多媒体播放器" -style bold -orientation vertical -speed 2
sleep 3
/usr/bin/xplayctl -play -libName scroll -zIndex 8 -screen_mode landscape -screen_rotate 180 -rect "710,0,500,-1" -font_size 20 -color "rgba(255,0,0,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统 设计的多媒体播放器" -style bold -orientation vertical -speed 2
sleep 3
/usr/bin/xplayctl -play -libName scroll -zIndex 8 -screen_mode portrait -screen_rotate 90 -rect "290,0,500,-1" -font_size 20 -color "rgba(255,0,0,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统 设计的多媒体播放器" -style bold -orientation vertical -speed 2
sleep 3
/usr/bin/xplayctl -play -libName scroll -zIndex 8 -screen_mode portrait -screen_rotate 270 -rect "290,0,500,-1" -font_size 20 -color "rgba(255,0,0,100%)" -bgcolor "rgba(0,0,0,30%)" -content "专为树莓派(Raspberry Pi)与Windows系统 设计的多媒体播放器" -style bold -orientation vertical -speed 2

sleep 3

/usr/bin/xplayctl -play -libName datetime -zIndex 7 -screen_mode landscape -screen_rotate 0 -font_size 30 -color "rgba(255,255,255,100%)" -bgcolor "rgba(0,0,0,20%)" -align center -style bold -rect "0,0,500,50"
sleep 3
/usr/bin/xplayctl -play -libName datetime -zIndex 7 -screen_mode landscape -screen_rotate 180 -font_size 30 -color "rgba(255,255,255,100%)" -bgcolor "rgba(0,0,0,20%)" -align center -style bold -rect "0,0,500,50"
sleep 3
/usr/bin/xplayctl -play -libName datetime -zIndex 7 -screen_mode portrait -screen_rotate 90 -font_size 30 -color "rgba(255,255,255,100%)" -bgcolor "rgba(0,0,0,20%)" -align center -style bold -rect "0,0,500,50"
sleep 3
/usr/bin/xplayctl -play -libName datetime -zIndex 7 -screen_mode portrait -screen_rotate 270 -font_size 30 -color "rgba(255,255,255,100%)" -bgcolor "rgba(0,0,0,20%)" -align center -style bold -rect "0,0,500,50"

sleep 3

/usr/bin/xplayctl -play -libName toast -zIndex 6 -screen_mode landscape -screen_rotate 0 -content "通知" -toast_type notice -duration 3
sleep 3
/usr/bin/xplayctl -play -libName toast -zIndex 6 -screen_mode landscape -screen_rotate 180 -content "成功" -toast_type success -duration 3
sleep 3
/usr/bin/xplayctl -play -libName toast -zIndex 6 -screen_mode portrait -screen_rotate 90 -content "警告" -toast_type warning -duration 3
sleep 3
/usr/bin/xplayctl -play -libName toast -zIndex 6 -screen_mode portrait -screen_rotate 270 -content "错误" -toast_type error -duration 3

sleep 3

/usr/bin/xplayctl -play -libName gif -zIndex 5 -screen_mode landscape -screen_rotate 0 -rect 0,0,495,374 -path `pwd`/balloon.gif
sleep 3
/usr/bin/xplayctl -play -libName gif -zIndex 5 -screen_mode landscape -screen_rotate 180 -rect 0,0,495,374 -path `pwd`/balloon.gif
sleep 3
/usr/bin/xplayctl -play -libName gif -zIndex 5 -screen_mode portrait -screen_rotate 90 -rect 0,0,495,374 -path `pwd`/balloon.gif
sleep 3
/usr/bin/xplayctl -play -libName gif -zIndex 5 -screen_mode portrait -screen_rotate 270 -rect 0,0,495,374 -path `pwd`/balloon.gif

sleep 3


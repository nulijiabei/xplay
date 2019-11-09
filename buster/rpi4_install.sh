#!/bin/bash

dir=$(/usr/bin/dirname $0)

cp -aRrf ${dir}/bin/xplay /usr/bin/
cp -aRrf ${dir}/libs/rpi4/* /usr/local/lib/

mkdir /etc/xplay
cp -aRrf ${dir}/etc/log4qt.properties /etc/xplay/
cp -aRrf ${dir}/etc/simsun.ttc /etc/xplay/
cp -aRrf ${dir}/etc/00-xplay.conf /etc/ld.so.conf.d/
ldconfig

systemctl disable display-manager

echo "Successful installation -> reboot -> /usr/bin/xplay"

#!/bin/bash

dir=$(/usr/bin/dirname $0)

if [ ! -d /etc/xplay ];then
    mkdir /etc/xplay
fi

cp -aRrf ${dir}/bin/xplay /usr/bin/
cp -aRrf ${dir}/bin/xplayctl /usr/bin/
cp -aRrf ${dir}/bin/xplayrun /usr/bin/
cp -aRrf ${dir}/libs/rpi/* /usr/local/lib/
cp -aRrf ${dir}/libs/rpi4_x11/* /usr/local/lib/

cp -aRrf ${dir}/etc/log4qt.properties /etc/xplay/
cp -aRrf ${dir}/etc/simsun.ttc /etc/xplay/
cp -aRrf ${dir}/etc/00-xplay.conf /etc/ld.so.conf.d/
ldconfig

echo "/usr/bin/xplay" > /etc/xplay/library.file
echo "/usr/bin/xplayctl" >> /etc/xplay/library.file
echo "/usr/bin/xplayrun" >> /etc/xplay/library.file
for lib in `ls ${dir}/libs/rpi/`;do
    echo "/usr/local/lib/$lib" >> /etc/xplay/library.file
done
for lib in `ls ${dir}/libs/rpi4_x11/`;do
    echo "/usr/local/lib/$lib" >> /etc/xplay/library.file
done

echo "/etc/xplay/log4qt.properties" >> /etc/xplay/library.file
echo "/etc/xplay/simsun.ttc" >> /etc/xplay/library.file
echo "/etc/ld.so.conf.d/00-xplay.conf" >> /etc/xplay/library.file

raspi-config nonint do_memory_split 256

echo "Successful installation -> reboot -> /usr/bin/xplay"

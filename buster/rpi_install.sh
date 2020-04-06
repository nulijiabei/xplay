#!/bin/bash

dir=$(/usr/bin/dirname $0)

cp -aRrf ${dir}/bin/xplay /usr/bin/
cp -aRrf ${dir}/bin/xplayctl /usr/bin/
cp -aRrf ${dir}/libs/rpi/* /usr/local/lib/

mkdir /etc/xplay
cp -aRrf ${dir}/etc/log4qt.properties /etc/xplay/
cp -aRrf ${dir}/etc/simsun.ttc /etc/xplay/
cp -aRrf ${dir}/etc/00-xplay.conf /etc/ld.so.conf.d/
ldconfig

raspi-config nonint do_memory_split 256

echo "Successful installation -> /usr/bin/xplay"



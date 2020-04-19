#!/bin/bash

dir=$(/usr/bin/dirname $0)

sh -x ${dir}/../buster/rpi_install.sh

cp -aRrf ${dir}/libs/rpi/* /usr/local/lib/
for lib in `ls ${dir}/libs/rpi/`;do
    echo "/usr/local/lib/$lib" >> /etc/xplay/library.file
done

ldconfig

echo "Successful installation -> /usr/bin/xplay"

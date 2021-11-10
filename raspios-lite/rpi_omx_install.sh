#!/bin/bash

dir=$(/usr/bin/dirname $0)

sh -x ${dir}/../buster/rpi_omx_install.sh

cp -aRrf ${dir}/libs/rpi_omx/* /usr/local/lib/
for lib in `ls ${dir}/libs/rpi_omx/`;do
    echo "/usr/local/lib/$lib" >> /etc/xplay/library.file
done

ldconfig

echo "Successful installation -> /usr/bin/xplay"

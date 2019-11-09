#!/bin/bash

dir=$(/usr/bin/dirname $0)

sh -x ${dir}/../buster/rpi_install.sh

cp -aRrf ${dir}/libs/rpi/* /usr/local/lib/

ldconfig

echo "Successful installation -> /usr/bin/xplay"

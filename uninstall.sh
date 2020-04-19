#!/bin/bash

if [ -e /etc/xplay/library.file ];then
    for file in `cat /etc/xplay/library.file`;do
        rm -f $file
    done
fi

echo "Successful uninstall"

# @Author: jiangzhuolin
# @Date:   2017-12-16 18:58:35
# @Last Modified by:   jiangzhuolin
# @Last Modified time: 2017-12-16 20:07:30

#! /bin/sh

kill -term $(pgrep webserver}
cd /opt/software/golang/first-go
git pull https://github.com/jiangzhuolin/first-go.git
cd webserver/
./webserver &

# @Author: jiangzhuolin
# @Date:   2017-12-16 18:58:35
# @Last Modified by:   jiangzhuolin
# @Last Modified time: 2017-12-17 01:31:05

#! /bin/sh

DATE=`date +"%Y-%m-%d %H:%M:%S"`

echo "${DATE} | start to restart webserver..." >> ./deploy.log
kill -9 $(pgrep webserver}
cd /opt/software/golang/first-go
#git pull https://github.com/jiangzhuolin/first-go.git
git pull
cd webserver/
./webserver &

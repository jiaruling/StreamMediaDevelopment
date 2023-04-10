#! /bin/bash

# 拷贝资源
echo "copy resource"
cp -R ../template ../bin/
cp ../conf.json ../bin/

if [ ! -d ../bin/videos ]; then
    mkdir ../bin/videos
fi

cd ../bin

# 删除之前的进程
echo "kill api..."
pid=$(netstat -ntlp | grep ./api | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }')
if [ $pid -gt 0 ]; then
    kill $pid
fi
if [ ! -d ./api.log ]; then
    rm -f ./api.log
fi
echo "restart api..."
nohup ./api > api.log 2>&1 &

echo "kill scheduler..."
pid=$(netstat -ntlp | grep ./scheduler | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }')
if [ $pid -gt 0 ]; then
    kill $pid
fi
if [ ! -d ./scheduler.log ]; then
    rm -f ./scheduler.log
fi
echo "restart scheduler..."
nohup ./scheduler > scheduler.log 2>&1 &

echo "kill streamserver..."
pid=$(netstat -ntlp | grep ./streamserver | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }')
if [ $pid -gt 0 ]; then
    kill $pid
fi
if [ ! -d ./streamserver.log ]; then
    rm -f ./streamserver.log
fi
echo "restart streamserver..."
nohup ./streamserver > streamserver.log 2>&1 &

echo "kill web..."
pid=$(netstat -ntlp | grep ./web | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }')
if [ $pid -gt 0 ]; then
    kill $pid
fi
if [ ! -d ./web.log ]; then
    rm -f ./web.log
fi
echo "restart web..."
nohup ./web > web.log 2>&1 &

echo "deploy finished"
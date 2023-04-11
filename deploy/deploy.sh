#! /bin/bash

# 命令行参数
mod=$1 # stop start restart
if [ "$mod" != "stop" ] && [ "$mod" != "start" ] && [ "$mod" != "restart" ]; then
    echo "arguments error: $mod"
    exit 1
fi

# 拷贝资源
if [ "$mod" == "start" ] || [ "$mod" == "restart" ]; then
    echo "copy resource"
    cp -R ../template ../bin/
    cp ../conf.json ../bin/

    if [ ! -d ../bin/videos ]; then
        mkdir ../bin/videos
    fi

    cd ../bin
fi

# 删除之前的进程
if [ "$mod" == "stop" ] || [ "$mod" == "restart" ]; then
    echo "kill api..."
    pid=$(netstat -ntlp | grep ./api | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }')
    if [ -n "$pid" ] && [ $pid -gt 0 ]; then
        kill $pid
    fi
    echo "kill scheduler..."
    pid=$(netstat -ntlp | grep ./scheduler | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }')
    if [ -n "$pid" ] && [ $pid -gt 0 ]; then
        kill $pid
    fi
    echo "kill stream..."
    pid=$(netstat -ntlp | grep ./stream | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }')
    if [ -n "$pid" ] && [ $pid -gt 0 ]; then
        kill $pid
    fi
    echo "kill web..."
    pid=$(netstat -ntlp | grep ./web | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }')
    if [ -n "$pid" ] && [ $pid -gt 0 ]; then
        kill $pid
    fi
    echo "kill finished"
fi

if [ "$mod" == "start" ] || [ "$mod" == "restart" ]; then
    if [ ! -d ./api.log ]; then
        rm -f ./api.log
    fi
    echo "start api..."
    nohup ./api > api.log 2>&1 &

    if [ ! -d ./scheduler.log ]; then
        rm -f ./scheduler.log
    fi
    echo "start scheduler..."
    nohup ./scheduler > scheduler.log 2>&1 &

    if [ ! -d ./stream.log ]; then
        rm -f ./stream.log
    fi
    echo "start stream..."
    nohup ./stream > stream.log 2>&1 &

    if [ ! -d ./web.log ]; then
        rm -f ./web.log
    fi
    echo "start web..."
    nohup ./web > web.log 2>&1 &
    echo "deploy finished"
    sleep 1
    echo "*************************************************************************************"
    netstat -ntlp | grep -E "api|stream|scheduler|web"
    echo "*************************************************************************************"
fi

echo "execute finished"



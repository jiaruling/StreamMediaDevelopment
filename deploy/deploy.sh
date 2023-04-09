#! /bin/bash

cp -R ../templates ../bin/

mkdir ../bin/videos

cd ../bin

nohup ./api > api.log 2>&1 &
nohup ./scheduler > scheduler.log 2>&1 &
nohup ./streamserver > streamserver.log 2>&1 &
nohup ./web > web.log 2>&1 &

echo "deploy finished"
#! /bin/bash

pkill -f "template_server"
cd ..
./template_server -alsologtostderr -v=2 > ./script/stdout 2>&1 &

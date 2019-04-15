#! /bin/bash

curl -i -d '{"header":{"request_id":"0001","client_id":"c0001","client_type":0,"version":"v1"}, "data":{"dummy":1}}' "http://127.0.0.1:9003/gotp/TemplateServer/TemplateService/Test"

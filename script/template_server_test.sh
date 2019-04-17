#! /bin/bash

curl -i -d '{"requestId":"0001","clientId":"c0001","clientType":0,"version":"v1","data":{"dummy":1}}' "http://127.0.0.1:9003/gotp/TemplateServer/TemplateService/Test"

#!/bin/bash

if which docker 
then 
    echo "docker installed!!!"
else 
    echo "please install docker"
    exit 1
fi
if which docker-compose 
then
    echo "docker-compose installed!!!"
else
    echo "please install docker-compose"
    exit 2
fi
go get gopkg.in/yaml.v2
go test -v
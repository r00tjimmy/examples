#! /bin/bash

#-----------server------------
consul agent -server -bootstrap-expect 1  -data-dir /tmp/consul -node=agent-one -bind=172.20.73.239 &


#----------client------------
#consul agent -data-dir /tmp/consul -node=agent-two -bind=172.20.73.229
#consul join 172.20.73.239

cd /opt/gitcode/examples/puffin/srv/
go run main.go &

cd /opt/gitcode/examples/puffin/api/
go run api.go &

micro api





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


#------------------------- test ----------------------------------------------------------

http://172.20.73.239:8080/puffin/say/hello?name=Asim+Aslam


curl -d 'service=go.micro.srv.puffin' \
		-d 'method=Say.Hello' \
			-d 'request={"name": "Asim Aslam"}' \
				http://localhost:8080/rpc




curl -H 'Content-Type: application/json' \
		-d '{"service": "go.micro.srv.puffin", "method": "Say.Hello", "request": {"name": "Asim Aslam"}}' \
			http://localhost:8080/rpc


#-------------------------------------------------------------------------------------------



用这个东西来写API的流程:

1.  定义好接口服务，也就是 srv/proto 下面的proto 文件，生成接口服务的go代码文件
2.  编写srv下面的main.go， 为接口定义具体的逻辑
3.  跟srv同一级目录建立api目录，其中api目录可以分别建立rest和rpc两个目录， rest是指提供支持HTTP和通讯型RPC调用的功能，  rpc目录是提供源码级（例如python 客户端代码）的服务。 该目录的api.go 是声明该api的路由的，该路由应用于rest和rpc两个文件夹都生效。










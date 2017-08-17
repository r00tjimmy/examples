### micro service sample code with go-micro.

- make with devops, base on jenkins
- deploy engine with k8s & docker



##### greeter sample code for k8s
in folder greeter, steps like this:
1.  构建步骤:   CI in jenkins, make docker images, push to docker-hub, every code version for every image tag.
2.  构建后步骤:  SSH to the nodes of k8s, create deploy & service in k8s.

script for jenkins

.Dockerfile
```shell
FROM alpine:3.2
ADD greeter-srv /greeter-srv
ENTRYPOINT [ "/greeter-srv" ]
```

script for jenkins
```shell

docker built -t r00t-micro/xxx:v1 .
docker push xx.xx.xx/xx

```


script for k8s
```shell

kubectl create -f xx.yml

```



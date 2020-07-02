# go-runlike
根据`docker inspect`指令生成的Json串反向解析`docker run`指令的工具。

借鉴了 https://github.com/lavie/runlike 项目的代码编写的Golang版本，并对一些冗余的参数做了优化修剪。

目前仅对`docker run`的部分参数进行了解析，设计参数如下：
* -d
* -ti
* --rm
* --restart
* --name
* --pid
* --ipc
* --network
* --privileged
* -p, --port
* --hostname
* --user
* --mac-address
* --link
* --add-host
* -e, --env
* -v, --volume
* --volumes-from
* --cap-add
* --cap-drop
* --device
* --dns
* --log-driver
* --log-opt
* --label
* --entrypoint

# 声明
项目仅供个人学习使用。
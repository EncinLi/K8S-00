FROM ubuntu
ENV MY_SERVICE_PORT=80
# ENV MY_SERVICE_PORT1=80
# ENV MY_SERVICE_PORT2=80
# ENV MY_SERVICE_PORT3=80
# LABEL multi.label1="value1" multi.label2="value2" other="value3"
# 把本地编译好的文件添加到镜像，win系统应指定GOOS=linux系统打包 main.go
# https://github.com/EncinLi/K8S-00/tree/main/HomeWork-01/main.go 不确定怎么编译GitHub上的，或者先添加源文件再go build
ADD D:/Study/K8S-00/HomeWork-01/main /httpserver
EXPOSE 80
# Dockerfile中只允许有一个ENTRYPOINT命令，多指定时会覆盖前面的设置，而只执行最后的ENTRYPOINT指令
ENTRYPOINT /httpserver
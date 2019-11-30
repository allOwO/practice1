# image
FROM centos:latest

MAINTAINER lzx

ENV GOPATH=/root/go
ENV GOPROXY=https://goproxy.io
CMD ["/usr/sbin/init"] 
## 开启服务
ADD ./initserver.sh /root/messenger/
ADD ./db.sql 		/root/messenger/
ADD ./messenger 	/root/messenger/
ADD ./config.yaml   /root/messenger/
ADD ./vendor        /root/messenger/
WORKDIR /root/messenger/
RUN yum install -y golang mysql socat epel-release git
RUN yum install -y http://www.rabbitmq.com/releases/erlang/erlang-19.0.4-1.el7.centos.x86_64.rpm http://www.rabbitmq.com/releases/rabbitmq-server/v3.6.15/rabbitmq-server-3.6.15-1.el7.noarch.rpm
#	&& mkdir -p /root/go/src/google.golang.org \
#	&& cd /root/go/src/google.golang.org/ \
#	&& git clone https://github.com/grpc/grpc-go.git \
#    && mv grpc-go grpc \
RUN go get github.com/grpc/grpc-go \
	&& cp /root/go/src/github.com/grpc/grpc-go /root/go/src/google.golang.org/grpc \
	&& ./initserver.sh


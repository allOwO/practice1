# image
FROM centos:latest

MAINTAINER lzx

ENV GOPATH=/root 
ENV GOPROXY=https://goproxy.io
CMD ["/usr/sbin/init"] 
## 开启服务
ADD ./initserver.sh /root/messenger/
ADD ./db.sql 		/root/messenger/
ADD ./messenger 	/root/messenger/
ADD ./config.yaml   /root/messenger
# 方便后面go get使用goproxy代理
ADD ./go.sum 		/root/messenger/
ADD ./go.mod    	/root/messenger/
WORKDIR /root/messenger/
RUN  yum install -y golang wget mysql socat epel-release \
	&& wget http://packages.erlang-solutions.com/erlang-solutions-1.0-1.noarch.rpm \
	&& wget http://www.rabbitmq.com/releases/rabbitmq-server/v3.6.6/rabbitmq-server-3.6.6-1.el6.noarch.rpm \
	&& rpm -Uvh erlang-solutions-1.0-1.noarch.rpm \
	&& yum install -y erlang rabbitmq-server-3.6.6-1.el6.noarch.rpm \
	&& go get -u google.golang.org/grpc \
	&& ./initserver.sh


# image
FROM centos:latest

MAINTAINER lzx

ENV GOPATH=/root/go
## 开启服务
ADD ./initserver.sh /root/messenger/
ADD ./db.sql 		/root/messenger/
ADD ./messenger 	/root/messenger/
ADD ./config.yaml   /root/messenger/
ADD ./vendor        /root/go/vendor
WORKDIR /root/messenger/
RUN yum install -y golang mysql-server socat epel-release git
RUN yum install -y http://www.rabbitmq.com/releases/erlang/erlang-19.0.4-1.el7.centos.x86_64.rpm http://www.rabbitmq.com/releases/rabbitmq-server/v3.6.15/rabbitmq-server-3.6.15-1.el7.noarch.rpm
RUN cp -r /root/go/vendor /root/go/src \
    && go install  google.golang.org/grpc \
CMD ["/bin/sh","/root/messenger/initserver.sh"]


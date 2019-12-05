# image
FROM golang:1.13

MAINTAINER lzx

ENV GOPATH=/root/go
ENV GOPROXY=direct,https://goproxy.io
## 开启服务
RUN mkdir -p /root/messenger
WORKDIR /root/messenger
#CMD ["/root/messenger/messenger","jobs","sender","--type","mail","-n","2"]
ENTRYPOINT  ["/root/messenger/init_messenger.sh"]
# 练习项目
- 项目中的Dockerfile文件
    - Dockerfile
    - initserver.sh
    - db.sql
    
- 配置文件
    - config.yaml
- 三种分组类型(需要输入这三种才能发送)
    - `service_staff` 客服
    - `system_user` 微信公众号管理员
    - `worker`      运营
## 测试
- 运行`make`
- 运行`Dockerfile`
- `docker run -tid --privileged <image id> /usr/sbin/init`
- 配置文件默认在messenger同文件夹
- 消息服务
    - `./messenger server`
    - `worker#hello world`
- web  
    - `./messenger dashboard`
- 消费通知队列
    - `./messenger jobs notification -n 2`
- 启动消息发送程序发送消息
    - `./messenger jobs sender --type mail -n 2`

## 遇到一个问题
- `Exception (404) Reason: "NOT_FOUND - no queue 'message' in vhost 'guest`
-  jobManenger不能创建队列，只能使用现有的队列，使用amqp创建队列后就不会报错
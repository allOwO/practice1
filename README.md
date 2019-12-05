# 练习项目
- 配置文件
    - config.yaml
- 三种分组类型(需要输入这三种才能发送)
    - `service_staff` 客服
    - `system_user` 微信公众号管理员
    - `worker`      运营
- 项目中的Dockerfile文件
    - `Dockerfile`
    - `init_messenger.sh` 程序自启动脚本
    - `db.sql`
    - `docker-compose.yaml`
## 测试
- 运行`make` (window 需要`set GOOS=linux`)
- 运行`docker-compose up`
- 配置文件默认值在messenger同文件夹
- 消息服务
    - `./messenger server`
    - `worker#hello world`
- web  
    - `./messenger dashboard`
- 消费通知队列
    - `./messenger jobs notification -n 2`
- 启动消息发送程序发送消息
    - `./messenger jobs sender --type mail -n 2`

## 遇到的问题
- `Exception (404) Reason: "NOT_FOUND - no queue 'message' in vhost 'guest`
-  ~~jobManenger不能创建队列，只能使用现有的队列，使用amqp创建队列后就不会报错~~
- 对mq理解不到位，先启动消费者再启动生产者就不会出现这个问题。

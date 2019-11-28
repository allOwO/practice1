# 练习项目
- 项目中的Dockerfile文件
    - Dockerfile
    - initserver.sh
    - db.sql
    
- 配置文件
    - config.yaml

## 测试
- 配置文件默认在messenger同文件夹
- 消息服务
    - `./messenger server`
- web  
    - `./messenger dashboard`
- 消费通知队列
    - `./messenger jobs notification -n 2`
- 启动消息发送程序发送消息
    - `./messenger jobs sender --type mail -n 2`
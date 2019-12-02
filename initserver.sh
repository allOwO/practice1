#!/bin/bash
rabbitmq-server -detached
rabbitmq-plugins enable rabbitmq_management
systemctl start mysqld.service
mysql < /root/messenger/db.sql
echo `service mysql status`
echo `rabbitmqctl status`
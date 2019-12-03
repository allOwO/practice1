CREATE database `practice`;
use practice;
DROP TABLE IF EXISTS `all_users`;
CREATE TABLE `all_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL,
  `user_phone` varchar(255) NOT NULL,
  `user_mail` varchar(255) NOT NULL,
  `system_user` tinyint(1) DEFAULT NULL,
  `worker` tinyint(1) DEFAULT NULL,
  `service_staff` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_mail` (`user_mail`),
  KEY `idx_all_users_user_mail` (`user_mail`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
INSERT INTO `all_users` VALUES (1,'lzx','17645027795','575361715@qq.com',1,1,1,'2019-12-02 10:59:11','2019-12-02 10:59:11',NULL);
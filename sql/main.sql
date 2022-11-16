DROP TABLE IF EXISTS `cb_user`;

CREATE TABLE `cb_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user` varchar(255) DEFAULT NULL,
  `passwd` varchar(255) DEFAULT NULL,
  `balance_usdt` decimal(30,20) DEFAULT 0.00000000000000000000,
  `user_status` tinyint(2) DEFAULT NULL COMMENT '狀態(0)正常(1)封鎖',
  PRIMARY KEY (`id`),
  KEY `user` (`user`),
  KEY `status` (`user_status`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='會員資料';

LOCK TABLES `cb_user` WRITE;
/*!40000 ALTER TABLE `cb_user` DISABLE KEYS */;

INSERT INTO `cb_user` (`id`, `user`, `passwd`, `balance_usdt`, `user_status`)
VALUES
	(1,'Ody','fwh32rj',0.00000000000000000000,0),
	(2,'Jone','dfghwt',0.00000000000000000000,0),
	(3,'Jeff','fgjwerigj',0.00000000000000000000,0),
	(4,'Nara','ifiujkrmc',0.00000000000000000000,0),
	(5,'Noman','8unafjiawerf',0.00000000000000000000,0),
	(6,'Ketty','mcieakdf',0.00000000000000000000,0),
	(7,'Smith','aviwejfim78',0.00000000000000000000,0),
	(8,'Mickey','ewrfjo34mf',0.00000000000000000000,0),
	(9,'Jenny','werifko',0.00000000000000000000,0),
	(10,'Ocess','cmawerifj',0.00000000000000000000,0),
	(11,'Lion','erifjmrwmf',0.00000000000000000000,0),
	(12,'Hans','fkmormerk',0.00000000000000000000,0),
	(13,'Wula','jficerfmmo',0.00000000000000000000,0),
	(14,'JoJo','fie3dmg',0.00000000000000000000,0),
	(15,'Grace','12ed8icfj',0.00000000000000000000,0);

/*!40000 ALTER TABLE `cb_user` ENABLE KEYS */;
UNLOCK TABLES;

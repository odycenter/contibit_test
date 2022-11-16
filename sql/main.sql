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


DROP TABLE IF EXISTS `cb_account_book`;

CREATE TABLE `cb_account_book` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '會員ID',
  `target_id` bigint(20) NOT NULL COMMENT '轉給誰',
  `transfer_type` tinyint(2) NOT NULL COMMENT '(in)轉入(out)轉出',
  `coin` varchar(20) NOT NULL DEFAULT '' COMMENT '幣種',
  `amount` decimal(30,10) NOT NULL COMMENT '轉帳金額',
  `status` tinyint(2) NOT NULL COMMENT '(0)完成(1)失敗',
  `api_timestamp` varchar(255) NOT NULL DEFAULT '',
  `api_datetime` datetime NOT NULL,
  `api_key` varchar(255) NOT NULL DEFAULT '',
  `create_datetime` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `api_key` (`api_key`),
  KEY `user_id` (`user_id`),
  KEY `target_id` (`target_id`),
  KEY `api_key_2` (`api_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
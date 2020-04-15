-- Host: localhost  (Version 8.0.15)
-- Date: 2020-04-15 18:13:17
-- Generator: MySQL-Front 6.0  (Build 2.20)


--
-- Structure for table "tbl_casbin_rule"
--

CREATE TABLE `tbl_casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Data for table "tbl_casbin_rule"
--

INSERT INTO `tbl_casbin_rule` VALUES
    ('p','admin','/v1/auth','GET',NULL,NULL,NULL),
    ('p','admin','/v1/user','GET',NULL,NULL,NULL),
    ('p','admin','/v1/user/:uid','GET',NULL,NULL,NULL),
    ('p','admin','/v1/user/:uid','PUT',NULL,NULL,NULL),
    ('p','admin','/v1/user/:uid','DELETE',NULL,NULL,NULL),
    ('p','normal','/v1/auth','GET',NULL,NULL,NULL),
    ('p','normal','/v1/user/:uid','GET',NULL,NULL,NULL),
    ('p','admin','/v1/auth/logout','POST',NULL,NULL,NULL),
    ('p','admin','/v1/auth/password','PUT',NULL,NULL,NULL),
    ('p','normal','/v1/auth/logout','POST',NULL,NULL,NULL),
    ('p','normal','/v1/auth/password','PUT',NULL,NULL,NULL);

--
-- Structure for table "tbl_user"
--

CREATE TABLE `tbl_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT 'normal',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT '2000-01-01 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

--
-- Data for table "tbl_user"
--

INSERT INTO `tbl_user` VALUES
    (1,'user1','$2a$10$OsijT.xFdPKiZFZCnvZfBOapelOY58a55Vl7pzyYNS/ENBA5oVksa','admin','2020-04-15 18:11:40','2020-04-15 18:11:40','2000-01-01 00:00:00'),
    (2,'user2','$2a$10$G8byB0ZzywrvgAMb1EFjaee3qh0J4O.eSRdZbc9AJxB7SqgDt89fC','normal','2020-04-15 18:11:46','2020-04-15 18:11:46','2020-04-15 18:12:33'),
    (3,'user3','$2a$10$dlGuUk3DvGs9kvpMTrH65uy7XPHyiF5/vgWK3tWqj9xgcyDOq8YY2','normal','2020-04-15 18:11:49','2020-04-15 18:11:49','2000-01-01 00:00:00'),
    (4,'user4','$2a$10$vHDtQfLJ5mMB6eBUfkohX.pueMy3LrSKLJelc8zI0iXBadDF566Wa','normal','2020-04-15 18:11:51','2020-04-15 18:11:51','2000-01-01 00:00:00'),
    (5,'user5','$2a$10$tWidzpk4B/vhAa8QnqWq5OpKVe6II460skLpGTcP8uTJgV7hvAmPO','normal','2020-04-15 18:11:53','2020-04-15 18:11:53','2000-01-01 00:00:00');

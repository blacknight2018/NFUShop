/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50640
Source Host           : localhost:3306
Source Database       : nfushop

Target Server Type    : MYSQL
Target Server Version : 50640
File Encoding         : 65001

Date: 2020-12-21 13:43:40
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(255) NOT NULL,
  `sex` char(1) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `detail` varchar(255) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of address
-- ----------------------------
INSERT INTO `address` VALUES ('15', '吴浩东', 'F', '110', '中大南方', '2020-12-18 15:38:34', '2020-12-18 15:38:34', '1');
INSERT INTO `address` VALUES ('16', '陈华泽', 'M', '13078255125', '天堂', '2020-12-18 16:08:41', '2020-12-18 16:08:41', '1');
INSERT INTO `address` VALUES ('17', '老王', 'F', '15580398648', '地狱', '2020-12-18 16:09:04', '2020-12-18 16:09:04', '1');
INSERT INTO `address` VALUES ('18', '吴浩东', 'F', '110', '中大南方东35的518D床', '2020-12-18 16:20:13', '2020-12-18 16:20:13', '1');
INSERT INTO `address` VALUES ('19', '吴浩东', 'M', '110', '中大南方东35的518D床', '2020-12-18 18:45:44', '2020-12-18 18:45:44', '1');
INSERT INTO `address` VALUES ('20', '吴浩东', 'M', '110', '中大南方东35的518D床', '2020-12-18 18:45:51', '2020-12-18 18:45:51', '1');
INSERT INTO `address` VALUES ('21', '1', 'M', '2', '3', '2020-12-20 22:49:16', '2020-12-20 22:49:16', '1');
INSERT INTO `address` VALUES ('22', '4', 'M', '5', '6', '2020-12-20 22:49:30', '2020-12-20 22:49:30', '1');
INSERT INTO `address` VALUES ('23', '7', 'M', '8', '9', '2020-12-20 22:49:43', '2020-12-20 22:49:43', '1');

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `sub_goods_id` int(11) NOT NULL,
  `amount` int(255) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of cart
-- ----------------------------

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `banner` text NOT NULL,
  `template` varchar(255) DEFAULT '',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  `desc` varchar(255) NOT NULL,
  `detail_img` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES ('1', '三只松鼠坚果零食大礼包', '[\"https://gd1.alicdn.com/imgextra/i1/2208039733002/O1CN012fHK5J1Y2xbHX0oRy_!!2208039733002.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-10 00:32:18', '2020-12-20 15:25:58', '非常好吃', '[\"https://img.alicdn.com/imgextra/i1/2208039733002/O1CN01ypfIPr1Y2xdLvXcNv_!!2208039733002.jpg\"]');
INSERT INTO `goods` VALUES ('2', '好吃的零食裤子', '[\"https://img.alicdn.com/imgextra/i2/725677994/O1CN01dOUiO528vIi9UcOvh_!!725677994.jpg_430x430q90.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-10 00:32:40', '2020-12-20 15:01:57', '三只老鼠', '[\"https://img.alicdn.com/imgextra/i4/2200628041767/O1CN01NeOVji1OvKPOeV76Q_!!2200628041767-0-scmitem6000.jpg\"]');
INSERT INTO `goods` VALUES ('3', '裤子零食', '[\"https://img.alicdn.com/imgextra/i1/2209151972136/O1CN01EMZjGS1ReKdG2HCw8_!!2209151972136.jpg\",\"https://img.alicdn.com/imgextra/i4/2209151972136/O1CN0159vHdp1ReKdQ6UTdB_!!2209151972136.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-11 18:20:02', '2020-12-20 15:28:18', '保暖好穿', '[\"https://img.alicdn.com/imgextra/i4/2209151972136/O1CN01T9kie31ReKdPNXBz1_!!2209151972136.jpg\",\"https://img.alicdn.com/imgextra/i4/2209151972136/O1CN01ZArePy1ReKdWtmQPy_!!2209151972136.jpg\",\"https://img.alicdn.com/imgextra/i2/2209151972136/O1CN01SJkm6x1ReKcUTly9g_!!2209151972136.jpg\",\"https://img.alicdn.com/imgextra/i2/2209151972136/O1CN010vXiAB1ReKca9307U_!!2209151972136.jpg\",\"https://img.alicdn.com/imgextra/i2/2209151972136/O1CN01Hbjtv91ReKcOHK23m_!!2209151972136.jpg\"]');
INSERT INTO `goods` VALUES ('4', '三只松鼠零食大礼包', '[\"https://gd1.alicdn.com/imgextra/i1/3198456634/O1CN01a0oEV91ysQ7aDzGy5_!!3198456634.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-20 15:26:57', '2020-12-20 15:27:06', '超级好吃', '[\"https://img.alicdn.com/imgextra/i3/3198456634/O1CN01K7CgZc1ysQ8xSGbfY_!!3198456634.jpg\"]');
INSERT INTO `goods` VALUES ('5', '三只松鼠圣诞节零食大礼包', '[\"https://gd3.alicdn.com/imgextra/i1/1736713931/O1CN014bZ3E41euRZm7ZVt0_!!1736713931.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-20 15:30:15', '2020-12-20 15:30:17', '圣诞礼包', '[\"https://img.alicdn.com/imgextra/i3/1736713931/O1CN01NHGuKQ1euRZl4PDn3_!!1736713931.jpg\"]');
INSERT INTO `goods` VALUES ('6', '三只松鼠零食大礼包组合一', '[\"https://gd1.alicdn.com/imgextra/i1/2869920720/O1CN01eMBScD1HBnoB0iiLS_!!2869920720.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-20 15:31:04', '2020-12-20 15:31:07', '大礼包组合一整箱巨型女生充饥夜宵肉食', '[\"https://img.alicdn.com/imgextra/i4/2869920720/O1CN01xO9kr81HBno88AEmB_!!2869920720.jpg\"]');
INSERT INTO `goods` VALUES ('7', '网红猪饲料零食', '[\"https://gd2.alicdn.com/imgextra/i2/1587589345/O1CN01fEzSHU2Iu3pdaxFNb_!!1587589345.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-20 15:34:11', '2020-12-20 15:44:34', '猪饲料', '[\"https://img.alicdn.com/imgextra/i4/1587589345/O1CN01EMFEOb2Iu3pfESyHR_!!1587589345.jpg\"]');
INSERT INTO `goods` VALUES ('8', '三只松鼠碧根果零食', '[\"https://gd3.alicdn.com/imgextra/i3/2336870737/O1CN01y0391G1HJaVN9c9pb_!!2336870737.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-20 16:11:54', '2020-12-20 16:12:44', '碧根果', '[\"https://img.alicdn.com/imgextra/i3/351277986/O1CN01KDoCzi28rdcS7xnAs_!!351277986.jpg\"]');

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `nick_name` varchar(255) NOT NULL,
  `sex` char(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `detail` varchar(255) NOT NULL,
  `sub_goods` varchar(255) NOT NULL,
  `total_price` decimal(10,0) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` int(255) DEFAULT NULL,
  `delivery_code` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES ('2', '0', '吴浩东', 'F', '110', '中大南方', '[3]', '44', '2020-12-18 19:38:40', '2020-12-21 01:20:37', '1', null);
INSERT INTO `order` VALUES ('3', '0', '吴浩东', 'F', '110', '中大南方', '[2]', '1600', '2020-12-18 19:59:26', '2020-12-21 01:20:38', '1', null);
INSERT INTO `order` VALUES ('4', '0', '吴浩东', 'F', '110', '中大南方', '[5]', '8888', '2020-12-18 20:00:13', '2020-12-21 01:20:39', '1', null);
INSERT INTO `order` VALUES ('5', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[5]', '8888', '2020-12-18 20:18:32', '2020-12-21 01:20:39', '1', null);
INSERT INTO `order` VALUES ('6', '1', '陈华泽', 'M', '13078255125', '天堂', '[7]', '6666', '2020-12-18 20:18:54', '2020-12-21 01:20:40', '1', null);
INSERT INTO `order` VALUES ('7', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', '11', '2020-12-18 20:36:24', '2020-12-21 01:20:40', '1', null);
INSERT INTO `order` VALUES ('8', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', '11', '2020-12-18 20:37:16', '2020-12-21 01:20:40', '1', null);
INSERT INTO `order` VALUES ('9', '1', '吴浩东', 'F', '110', '中大南方东35的518D床', '[3]', '11', '2020-12-18 20:38:59', '2020-12-21 01:20:40', '1', null);
INSERT INTO `order` VALUES ('10', '1', '吴浩东', 'F', '110', '中大南方东35的518D床', '[3]', '11', '2020-12-18 20:44:12', '2020-12-21 01:20:42', '1', null);
INSERT INTO `order` VALUES ('11', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', '11', '2020-12-18 20:45:10', '2020-12-21 01:20:41', '1', null);
INSERT INTO `order` VALUES ('12', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', '11', '2020-12-18 20:45:33', '2020-12-21 01:20:44', '1', null);
INSERT INTO `order` VALUES ('13', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', '11', '2020-12-18 20:50:18', '2020-12-21 01:20:42', '1', null);
INSERT INTO `order` VALUES ('14', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', '22', '2020-12-18 20:54:34', '2020-12-21 01:20:43', '1', null);
INSERT INTO `order` VALUES ('15', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', '11', '2020-12-18 20:55:28', '2020-12-21 01:20:43', '1', null);
INSERT INTO `order` VALUES ('16', '1', '陈华泽', 'M', '13078255125', '天堂', '[1]', '123', '2020-12-18 22:01:35', '2020-12-21 01:20:43', '1', null);
INSERT INTO `order` VALUES ('17', '1', '老王', 'F', '15580398648', '地狱', '[5]', '8888', '2020-12-18 22:08:13', '2020-12-21 01:14:57', '1', null);
INSERT INTO `order` VALUES ('18', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[5]', '8888', '2020-12-19 21:03:49', '2020-12-21 01:20:44', '1', null);
INSERT INTO `order` VALUES ('19', '1', '吴浩东', 'F', '110', '中大南方', '[5]', '8888', '2020-12-20 14:58:06', '2020-12-21 01:20:44', '1', null);
INSERT INTO `order` VALUES ('20', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[9]', '22', '2020-12-20 14:59:15', '2020-12-21 01:20:45', '1', null);
INSERT INTO `order` VALUES ('21', '1', '老王', 'F', '15580398648', '地狱', '[5]', '8888', '2020-12-20 15:55:50', '2020-12-21 01:20:45', '1', null);
INSERT INTO `order` VALUES ('22', '1', '老王', 'F', '15580398648', '地狱', '[1,9]', '145', '2020-12-20 20:32:50', '2020-12-21 01:20:46', '1', null);
INSERT INTO `order` VALUES ('23', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[2,9]', '422', '2020-12-21 01:19:00', '2020-12-21 01:19:00', '1', null);
INSERT INTO `order` VALUES ('24', '1', '陈华泽', 'M', '13078255125', '天堂', '[9,5]', '8910', '2020-12-21 01:21:30', '2020-12-21 01:21:30', '1', null);
INSERT INTO `order` VALUES ('25', '1', '吴浩东', 'M', '110', '中大南方东35的518D床', '[5]', '8888', '2020-12-21 13:19:35', '2020-12-21 13:19:35', '1', null);
INSERT INTO `order` VALUES ('26', '1', '1', 'M', '2', '3', '[9]', '22', '2020-12-21 13:32:51', '2020-12-21 13:32:51', '1', null);

-- ----------------------------
-- Table structure for sub_goods
-- ----------------------------
DROP TABLE IF EXISTS `sub_goods`;
CREATE TABLE `sub_goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `price` decimal(10,0) NOT NULL,
  `stoke` int(255) unsigned NOT NULL,
  `sell` decimal(10,0) NOT NULL,
  `img` varchar(255) NOT NULL,
  `goods_id` int(11) NOT NULL,
  `template` varchar(20) NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `f` (`goods_id`,`template`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sub_goods
-- ----------------------------
INSERT INTO `sub_goods` VALUES ('1', '123', '75', '3', 'https://gd2.alicdn.com/imgextra/i2/2329922645/O1CN012Jdvcx1VPS5fsuWku_!!2329922645.jpg', '3', '[0,0,2]', '2020-12-10 00:34:18', '2020-12-20 20:32:50');
INSERT INTO `sub_goods` VALUES ('2', '400', '95', '3', 'https://gd4.alicdn.com/imgextra/i4/2329922645/O1CN01P9qivz1VPS5gWY6wx_!!2329922645.jpg', '3', '[0,1,0]', '2020-12-10 14:17:28', '2020-12-21 01:19:00');
INSERT INTO `sub_goods` VALUES ('3', '11', '0', '10', 'https://img.alicdn.com/imgextra/i3/6000000003991/O1CN01WcHfEz1fLvLI34qPz_!!6000000003991-0-zao.jpg', '3', '[0,1,2]', '2020-12-11 18:20:19', '2020-12-18 20:55:28');
INSERT INTO `sub_goods` VALUES ('4', '9999', '9999', '1', 'https://gd1.alicdn.com/imgextra/i3/68682830/O1CN010cb1V41WmBVYSfzse_!!68682830.jpg', '3', '[1,1,1]', '2020-12-15 12:41:36', '2020-12-15 12:41:36');
INSERT INTO `sub_goods` VALUES ('5', '8888', '8880', '7', 'https://gd2.alicdn.com/imgextra/i2/1967345390/O1CN01t82aEi1pgfOAv0jGS_!!1967345390.jpg', '3', '[1,1,2]', '2020-12-15 12:44:37', '2020-12-21 13:19:35');
INSERT INTO `sub_goods` VALUES ('6', '7777', '7777', '1', 'https://gd3.alicdn.com/imgextra/i3/93763117/O1CN01IDu5KE1Ytd7hbwIIz_!!93763117.jpg', '3', '[2,2,2]', '2020-12-15 12:45:01', '2020-12-15 12:45:01');
INSERT INTO `sub_goods` VALUES ('7', '6666', '6665', '1', 'https://gd1.alicdn.com/imgextra/i1/93763117/O1CN01onRcOb1Ytd7iLtcG7_!!93763117.jpg', '3', '[1,2,2]', '2020-12-15 12:45:28', '2020-12-18 20:18:54');
INSERT INTO `sub_goods` VALUES ('8', '1269', '777', '1', 'https://img.alicdn.com/imgextra/i1/22588720/O1CN01cwzNDH2EHoK8Hc7mM_!!22588720.jpg', '3', '[0,2,1]', '2020-12-16 11:25:08', '2020-12-16 11:25:30');
INSERT INTO `sub_goods` VALUES ('9', '22', '28', '6', 'https://g-search2.alicdn.com/img/bao/uploaded/i4/i3/880734502/O1CN01yNkytf1j7xjSTQNzu_!!0-item_pic.jpg_580x580Q90.jpg_.webp', '2', '[0,0,0]', '2020-12-20 14:39:50', '2020-12-21 13:32:51');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pass_word` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `avatar` varchar(255) DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '1', '1', '2020-12-09 23:03:55', '2020-12-17 11:03:42', 'https://img.alicdn.com/imgextra/i3/727807374/TB2WRIdvrZnBKNjSZFhXXc.oXXa_!!727807374-0-beehive-https://img.alicdn.com/imgextra/https://img.alicdn.com/imgextra/i4/2207651980344/O1CN01DQZzo91EPatK9E2Rc_!!2207651980344.jpg_430x430q90.jpg', 'Apple');

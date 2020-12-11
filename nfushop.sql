/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50640
Source Host           : localhost:3306
Source Database       : nfushop

Target Server Type    : MYSQL
Target Server Version : 50640
File Encoding         : 65001

Date: 2020-12-11 11:36:01
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of address
-- ----------------------------

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
INSERT INTO `cart` VALUES ('1', '3', '2', '1', '2020-12-10 00:29:54', '2020-12-10 00:29:54');
INSERT INTO `cart` VALUES ('2', '3', '3', '1', '2020-12-10 21:16:56', '2020-12-10 21:16:56');

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `banner` varchar(255) NOT NULL,
  `template` varchar(255) DEFAULT '',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  `desc` varchar(255) NOT NULL,
  `detail_img` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES ('1', '33', '11', '44', '2020-12-10 00:32:18', '0000-00-00 00:00:00', '55', '66');
INSERT INTO `goods` VALUES ('2', '33', '11', '44', '2020-12-10 00:32:40', '0000-00-00 00:00:00', '55', '66');

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
  `goods` varchar(255) NOT NULL,
  `total_price` decimal(10,0) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES ('1', '1', '33', 'M', '123213213', '666', '888', '3366', '2020-12-10 00:38:36', '2020-12-10 00:38:36');

-- ----------------------------
-- Table structure for sub_goods
-- ----------------------------
DROP TABLE IF EXISTS `sub_goods`;
CREATE TABLE `sub_goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `price` decimal(10,0) NOT NULL,
  `stoke` int(255) NOT NULL,
  `sell` decimal(10,0) NOT NULL,
  `img` varchar(255) NOT NULL,
  `goods_id` int(11) NOT NULL,
  `template` varchar(255) NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sub_goods
-- ----------------------------
INSERT INTO `sub_goods` VALUES ('1', '123', '77', '999', '66', '33', '1', '2020-12-10 00:34:18', '2020-12-10 14:30:30');
INSERT INTO `sub_goods` VALUES ('2', '400', '100', '100', '88', '33', '1', '2020-12-10 14:17:28', '2020-12-10 14:30:17');

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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'Fuck', '1333333', '2020-12-09 23:03:55', '2020-12-09 23:03:58', null);
INSERT INTO `user` VALUES ('2', '0123', '456', '2020-12-09 23:25:54', '2020-12-09 23:52:23', null);
INSERT INTO `user` VALUES ('3', '11', '33', '2020-12-09 23:29:52', '2020-12-09 23:29:52', null);
INSERT INTO `user` VALUES ('4', '11', '33', '2020-12-09 23:30:05', '2020-12-09 23:30:05', null);
INSERT INTO `user` VALUES ('5', '1', '1', '2020-12-09 23:52:47', '2020-12-11 00:10:00', null);
INSERT INTO `user` VALUES ('6', '1100', '13078255125', '2020-12-10 20:31:10', '2020-12-10 20:31:10', null);
INSERT INTO `user` VALUES ('7', '999999', '1307825512', '2020-12-10 20:34:13', '2020-12-10 20:34:13', null);

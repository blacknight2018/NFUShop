/*
 Navicat MySQL Data Transfer

 Source Server         : main.bybyte.cn
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : main.bybyte.cn:3306
 Source Schema         : nfushop

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 02/01/2021 01:01:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `sex` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `detail` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of address
-- ----------------------------
INSERT INTO `address` VALUES (15, '吴浩东', 'F', '110', '中大南方', '2020-12-18 15:38:34', '2020-12-18 15:38:34', 1);
INSERT INTO `address` VALUES (16, '陈华泽', 'M', '13078255125', '天堂', '2020-12-18 16:08:41', '2020-12-18 16:08:41', 1);
INSERT INTO `address` VALUES (17, '老王', 'F', '15580398648', '地狱', '2020-12-18 16:09:04', '2020-12-18 16:09:04', 1);
INSERT INTO `address` VALUES (18, '吴浩东', 'F', '110', '中大南方东35的518D床', '2020-12-18 16:20:13', '2020-12-18 16:20:13', 1);
INSERT INTO `address` VALUES (19, '吴浩东', 'M', '110', '中大南方东35的518D床', '2020-12-18 18:45:44', '2020-12-18 18:45:44', 1);

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `img` varchar(255) CHARACTER SET utf32 COLLATE utf32_bin NOT NULL,
  `sub_goods_id` int(11) NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf32 COLLATE = utf32_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of banner
-- ----------------------------
INSERT INTO `banner` VALUES (1, 'https://img.alicdn.com/imgextra/i4/2114733660/TB2_7EebMsSMeJjSspcXXXjFXXa_!!2114733660.jpg', 2, '2020-12-30 10:25:21', '2020-12-30 10:25:35');
INSERT INTO `banner` VALUES (8, 'https://nfu-shop.oss-cn-beijing.aliyuncs.com/87aa4e36-4ea6-4553-9053-e6349b948764.jpg', 2, '2020-12-30 15:46:25', '2020-12-30 15:46:25');

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `sub_goods_id` int(11) NOT NULL,
  `amount` int(255) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cart
-- ----------------------------
INSERT INTO `cart` VALUES (11, 1, 14, 1, '2021-01-02 00:51:11', '2021-01-02 00:51:11');
INSERT INTO `cart` VALUES (12, 1, 5, 1, '2021-01-02 00:55:58', '2021-01-02 00:55:58');
INSERT INTO `cart` VALUES (13, 1, 8, 1, '2021-01-02 01:00:41', '2021-01-02 01:00:41');

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `banner` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `template` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `detail_img` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES (1, '1', '[\"https://gd1.alicdn.com/imgextra/i1/2208039733002/O1CN012fHK5J1Y2xbHX0oRy_!!2208039733002.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-10 00:32:18', '2020-12-30 10:06:18', '非常好吃', '[\"https://img.alicdn.com/imgextra/i1/2208039733002/O1CN01ypfIPr1Y2xdLvXcNv_!!2208039733002.jpg\"]');
INSERT INTO `goods` VALUES (2, '1', '[\"https://img.alicdn.com/imgextra/i2/725677994/O1CN01dOUiO528vIi9UcOvh_!!725677994.jpg_430x430q90.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-10 00:32:40', '2020-12-27 13:42:44', '三只老鼠', '[\"https://img.alicdn.com/imgextra/i4/2200628041767/O1CN01NeOVji1OvKPOeV76Q_!!2200628041767-0-scmitem6000.jpg\"]');
INSERT INTO `goods` VALUES (3, '裤子零食', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/5ef76d42-40ea-419a-bf60-5691136e03a9.jpg\"]', '[{  \"title\":\"大小\",  \"values\":[10,11,12]}]', '2020-12-11 18:20:02', '2020-12-30 11:55:44', '保暖好穿', '[\"https://img.alicdn.com/imgextra/i4/2209151972136/O1CN01T9kie31ReKdPNXBz1_!!2209151972136.jpg\",\"https://img.alicdn.com/imgextra/i4/2209151972136/O1CN01ZArePy1ReKdWtmQPy_!!2209151972136.jpg\",\"https://img.alicdn.com/imgextra/i2/2209151972136/O1CN01SJkm6x1ReKcUTly9g_!!2209151972136.jp\"]');
INSERT INTO `goods` VALUES (4, '1', '[\"https://gd1.alicdn.com/imgextra/i1/3198456634/O1CN01a0oEV91ysQ7aDzGy5_!!3198456634.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-20 15:26:57', '2020-12-27 13:43:11', '超级好吃', '[\"https://img.alicdn.com/imgextra/i3/3198456634/O1CN01K7CgZc1ysQ8xSGbfY_!!3198456634.jpg\"]');
INSERT INTO `goods` VALUES (5, '1', '[\"https://gd3.alicdn.com/imgextra/i1/1736713931/O1CN014bZ3E41euRZm7ZVt0_!!1736713931.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-20 15:30:15', '2020-12-27 13:43:12', '圣诞礼包', '[\"https://img.alicdn.com/imgextra/i3/1736713931/O1CN01NHGuKQ1euRZl4PDn3_!!1736713931.jpg\"]');
INSERT INTO `goods` VALUES (6, '1', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/080300f2-f8e1-42b2-95c0-1cf9985d03b5.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-20 15:31:04', '2020-12-30 15:49:57', '大礼包组合22', '[\"https://img.alicdn.com/imgextra/i4/2869920720/O1CN01xO9kr81HBno88AEmB_!!2869920720.jpg\"]');
INSERT INTO `goods` VALUES (7, '1', '[\"https://gd2.alicdn.com/imgextra/i2/1587589345/O1CN01fEzSHU2Iu3pdaxFNb_!!1587589345.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-20 15:34:11', '2020-12-27 13:43:13', '猪饲料', '[\"https://img.alicdn.com/imgextra/i4/1587589345/O1CN01EMFEOb2Iu3pfESyHR_!!1587589345.jpg\"]');
INSERT INTO `goods` VALUES (8, '1', '[\"https://gd3.alicdn.com/imgextra/i3/2336870737/O1CN01y0391G1HJaVN9c9pb_!!2336870737.jpg\"]', '[\r\n{\r\n  \"title\":\"大小\",\r\n  \"values\":[10,11,12]\r\n}\r\n,\r\n{\r\n  \"title\":\"颜色\",\r\n  \"values\":[\"red\",\"blue\",\"green\"]\r\n}\r\n,\r\n{\r\n  \"title\":\"重量\",\r\n  \"values\":[\"1kg\",\"2kg\",\"3kg\"]\r\n}\r\n\r\n]', '2020-12-20 16:11:54', '2020-12-30 10:06:08', '碧根果', '[\"https://img.alicdn.com/imgextra/i3/351277986/O1CN01KDoCzi28rdcS7xnAs_!!351277986.jpg\"]');
INSERT INTO `goods` VALUES (9, '好吃的蛋糕', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/19b7ffda-4d5c-4216-b91e-2e2340484b42.jpg\",\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/5e865874-ec4d-4616-811b-74a125fc7d61.jpg\"]', '[\n{\"title\":\"size\",\"values\":[1,2,3]}\n]', '2020-12-27 14:08:18', '2020-12-27 14:09:40', '很好吃', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/3763de45-81f7-4992-b897-d3f0d162f274.jpg\"]');

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `sex` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `detail` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `sub_goods` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `total_price` decimal(10, 0) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` int(255) NULL DEFAULT NULL,
  `delivery_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (2, 0, '吴浩东', 'F', '110', '中大南方', '[3]', 44, '2020-12-18 19:38:40', '2020-12-21 01:20:37', 1, NULL);
INSERT INTO `order` VALUES (3, 0, '吴浩东', 'F', '110', '中大南方', '[2]', 1600, '2020-12-18 19:59:26', '2020-12-21 01:20:38', 1, NULL);
INSERT INTO `order` VALUES (4, 0, '吴浩东', 'F', '110', '中大南方', '[5]', 8888, '2020-12-18 20:00:13', '2020-12-21 01:20:39', 1, NULL);
INSERT INTO `order` VALUES (5, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[5]', 8888, '2020-12-18 20:18:32', '2020-12-21 01:20:39', 1, NULL);
INSERT INTO `order` VALUES (6, 1, '陈华泽', 'M', '13078255125', '天堂', '[7]', 6666, '2020-12-18 20:18:54', '2020-12-21 01:20:40', 1, NULL);
INSERT INTO `order` VALUES (7, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', 11, '2020-12-18 20:36:24', '2020-12-21 01:20:40', 1, NULL);
INSERT INTO `order` VALUES (8, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', 11, '2020-12-18 20:37:16', '2020-12-21 01:20:40', 1, NULL);
INSERT INTO `order` VALUES (9, 1, '吴浩东', 'F', '110', '中大南方东35的518D床', '[3]', 11, '2020-12-18 20:38:59', '2020-12-21 01:20:40', 1, NULL);
INSERT INTO `order` VALUES (10, 1, '吴浩东', 'F', '110', '中大南方东35的518D床', '[3]', 11, '2020-12-18 20:44:12', '2020-12-21 01:20:42', 1, NULL);
INSERT INTO `order` VALUES (11, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', 11, '2020-12-18 20:45:10', '2020-12-21 01:20:41', 1, NULL);
INSERT INTO `order` VALUES (12, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', 11, '2020-12-18 20:45:33', '2020-12-21 01:20:44', 1, NULL);
INSERT INTO `order` VALUES (13, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', 11, '2020-12-18 20:50:18', '2020-12-21 01:20:42', 1, NULL);
INSERT INTO `order` VALUES (14, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', 22, '2020-12-18 20:54:34', '2020-12-21 01:20:43', 1, NULL);
INSERT INTO `order` VALUES (15, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[3]', 11, '2020-12-18 20:55:28', '2020-12-21 01:20:43', 1, NULL);
INSERT INTO `order` VALUES (16, 1, '陈华泽', 'M', '13078255125', '天堂', '[1]', 123, '2020-12-18 22:01:35', '2020-12-21 01:20:43', 1, NULL);
INSERT INTO `order` VALUES (17, 1, '老王', 'F', '15580398648', '地狱', '[5]', 8888, '2020-12-18 22:08:13', '2020-12-21 01:14:57', 1, NULL);
INSERT INTO `order` VALUES (18, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[5]', 8888, '2020-12-19 21:03:49', '2020-12-21 01:20:44', 1, NULL);
INSERT INTO `order` VALUES (19, 1, '吴浩东', 'F', '110', '中大南方', '[5]', 8888, '2020-12-20 14:58:06', '2020-12-21 01:20:44', 1, NULL);
INSERT INTO `order` VALUES (20, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[9]', 22, '2020-12-20 14:59:15', '2020-12-21 01:20:45', 1, NULL);
INSERT INTO `order` VALUES (21, 1, '老王', 'F', '15580398648', '地狱', '[5]', 8888, '2020-12-20 15:55:50', '2020-12-21 01:20:45', 1, NULL);
INSERT INTO `order` VALUES (22, 1, '老王', 'F', '15580398648', '地狱', '[1,9]', 145, '2020-12-20 20:32:50', '2020-12-21 01:20:46', 1, NULL);
INSERT INTO `order` VALUES (23, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[2,9]', 422, '2020-12-21 01:19:00', '2020-12-21 01:19:00', 1, NULL);
INSERT INTO `order` VALUES (24, 1, '陈华泽', 'M', '13078255125', '天堂', '[9,5]', 8910, '2020-12-21 01:21:30', '2020-12-21 01:21:30', 1, NULL);
INSERT INTO `order` VALUES (25, 1, '吴浩东', 'M', '110', '中大南方东35的518D床', '[5]', 8888, '2020-12-21 13:19:35', '2020-12-21 13:19:35', 1, NULL);
INSERT INTO `order` VALUES (26, 1, '1', 'M', '2', '3', '[9]', 22, '2020-12-21 13:32:51', '2020-12-21 13:32:51', 1, NULL);
INSERT INTO `order` VALUES (27, 1, '吴浩东', 'F', '110', '中大南方东35的518D床', '[10]', 1, '2020-12-27 13:46:52', '2020-12-27 13:46:52', 0, '');
INSERT INTO `order` VALUES (30, 1, '吴浩东', 'F', '110', '中大南方东35的518D床', '[10]', 1, '2020-12-30 00:22:03', '2020-12-30 00:22:03', 0, '');
INSERT INTO `order` VALUES (31, 1, '吴浩东', 'F', '110', '中大南方', '[14]', 3, '2020-12-30 15:48:22', '2020-12-30 15:48:22', 0, '');

-- ----------------------------
-- Table structure for sub_goods
-- ----------------------------
DROP TABLE IF EXISTS `sub_goods`;
CREATE TABLE `sub_goods`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `price` decimal(10, 0) NOT NULL,
  `stoke` int(255) UNSIGNED NOT NULL,
  `sell` decimal(10, 0) NOT NULL,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `goods_id` int(11) NOT NULL,
  `template` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `f`(`goods_id`, `template`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sub_goods
-- ----------------------------
INSERT INTO `sub_goods` VALUES (4, 8787, 9999, 1, 'https://gd1.alicdn.com/imgextra/i3/68682830/O1CN010cb1V41WmBVYSfzse_!!68682830.jpg', 3, '[1,1,1]', '2020-12-15 12:41:36', '2020-12-30 15:16:03');
INSERT INTO `sub_goods` VALUES (5, 8888, 8880, 7, 'https://gd2.alicdn.com/imgextra/i2/1967345390/O1CN01t82aEi1pgfOAv0jGS_!!1967345390.jpg', 4, '[1,1,2]', '2020-12-15 12:44:37', '2020-12-27 11:39:55');
INSERT INTO `sub_goods` VALUES (6, 7777, 7777, 1, 'https://gd3.alicdn.com/imgextra/i3/93763117/O1CN01IDu5KE1Ytd7hbwIIz_!!93763117.jpg', 3, '[2,2,2]', '2020-12-15 12:45:01', '2020-12-15 12:45:01');
INSERT INTO `sub_goods` VALUES (7, 6666, 6665, 1, 'https://gd1.alicdn.com/imgextra/i1/93763117/O1CN01onRcOb1Ytd7iLtcG7_!!93763117.jpg', 5, '[1,2,2]', '2020-12-15 12:45:28', '2020-12-27 11:39:56');
INSERT INTO `sub_goods` VALUES (8, 1269, 777, 1, 'https://img.alicdn.com/imgextra/i1/22588720/O1CN01cwzNDH2EHoK8Hc7mM_!!22588720.jpg', 6, '[0,2,1]', '2020-12-16 11:25:08', '2020-12-27 11:39:58');
INSERT INTO `sub_goods` VALUES (9, 22, 28, 6, 'https://g-search2.alicdn.com/img/bao/uploaded/i4/i3/880734502/O1CN01yNkytf1j7xjSTQNzu_!!0-item_pic.jpg_580x580Q90.jpg_.webp', 2, '[0,0,0]', '2020-12-20 14:39:50', '2020-12-21 13:32:51');
INSERT INTO `sub_goods` VALUES (10, 1, 0, 2, 'https://nfu-shop.oss-cn-beijing.aliyuncs.com/974cbdb7-43d6-4c70-8b0d-d0e052b609aa.jpg', 7, '[0,1,1]', '2020-12-27 11:31:44', '2020-12-30 00:22:03');
INSERT INTO `sub_goods` VALUES (11, 11, 77, 44, 'https://nfu-shop.oss-cn-beijing.aliyuncs.com/974cbdb7-43d6-4c70-8b0d-d0e052b609aa.jpg', 7, '[0,2,2]', '2020-12-29 23:46:16', '2020-12-29 23:46:16');
INSERT INTO `sub_goods` VALUES (13, 1000, 3, 3, 'https://nfu-shop.oss-cn-beijing.aliyuncs.com/8d4ebe40-5a67-43b1-b4dc-90658aa8c1d7.jpg', 5, '[1,1,2]', '2020-12-30 15:39:30', '2020-12-30 15:39:30');
INSERT INTO `sub_goods` VALUES (14, 3, 1, 45, 'https://nfu-shop.oss-cn-beijing.aliyuncs.com/d55e3948-1afa-40ee-89fd-81f028b2b2ea.jpg', 9, '[0]', '2020-12-30 15:41:25', '2020-12-30 15:48:22');
INSERT INTO `sub_goods` VALUES (16, 66, 77, 88, 'https://nfu-shop.oss-cn-beijing.aliyuncs.com/fd3fda74-8793-430d-8acc-7f40a0cd0049.jpg', 9, '[1]', '2020-12-30 15:41:41', '2020-12-30 15:41:41');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pass_word` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `phone`(`phone`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '1', '1', '2020-12-09 23:03:55', '2020-12-17 11:03:42', 'https://img.alicdn.com/imgextra/i3/727807374/TB2WRIdvrZnBKNjSZFhXXc.oXXa_!!727807374-0-beehive-https://img.alicdn.com/imgextra/https://img.alicdn.com/imgextra/i4/2207651980344/O1CN01DQZzo91EPatK9E2Rc_!!2207651980344.jpg_430x430q90.jpg', 'Apple');
INSERT INTO `user` VALUES (3, '123456', '123', '2020-12-28 02:06:25', '2020-12-28 02:06:30', 'https://img.alicdn.com/imgextra/i3/727807374/TB2WRIdvrZnBKNjSZFhXXc.oXXa_!!727807374-0-beehive-https://img.alicdn.com/imgextra/https://img.alicdn.com/imgextra/i4/2207651980344/O1CN01DQZzo91EPatK9E2Rc_!!2207651980344.jpg_430x430q90.jpg', 'Apple');
INSERT INTO `user` VALUES (4, '11', '11', '2020-12-30 10:06:43', '2020-12-30 10:06:47', 'https://img.alicdn.com/imgextra/i3/727807374/TB2WRIdvrZnBKNjSZFhXXc.oXXa_!!727807374-0-beehive-https://img.alicdn.com/imgextra/https://img.alicdn.com/imgextra/i4/2207651980344/O1CN01DQZzo91EPatK9E2Rc_!!2207651980344.jpg_430x430q90.jpg', 'Apple');
INSERT INTO `user` VALUES (5, '22', '33', '2020-12-30 10:06:51', '2020-12-30 10:06:55', 'https://img.alicdn.com/imgextra/i3/727807374/TB2WRIdvrZnBKNjSZFhXXc.oXXa_!!727807374-0-beehive-https://img.alicdn.com/imgextra/https://img.alicdn.com/imgextra/i4/2207651980344/O1CN01DQZzo91EPatK9E2Rc_!!2207651980344.jpg_430x430q90.jpg', 'Apple');
INSERT INTO `user` VALUES (7, '33', '44', '2020-12-30 10:07:17', '2020-12-30 10:07:21', 'https://img.alicdn.com/imgextra/i3/727807374/TB2WRIdvrZnBKNjSZFhXXc.oXXa_!!727807374-0-beehive-https://img.alicdn.com/imgextra/https://img.alicdn.com/imgextra/i4/2207651980344/O1CN01DQZzo91EPatK9E2Rc_!!2207651980344.jpg_430x430q90.jpg', 'Apple');
INSERT INTO `user` VALUES (8, '55', '66', '2020-12-30 10:07:28', '2020-12-30 10:07:31', 'https://img.alicdn.com/imgextra/i3/727807374/TB2WRIdvrZnBKNjSZFhXXc.oXXa_!!727807374-0-beehive-https://img.alicdn.com/imgextra/https://img.alicdn.com/imgextra/i4/2207651980344/O1CN01DQZzo91EPatK9E2Rc_!!2207651980344.jpg_430x430q90.jpg', 'Apple');
INSERT INTO `user` VALUES (9, '66', '88', '2020-12-30 10:07:46', '2020-12-30 10:07:49', 'https://img.alicdn.com/imgextra/i3/727807374/TB2WRIdvrZnBKNjSZFhXXc.oXXa_!!727807374-0-beehive-https://img.alicdn.com/imgextra/https://img.alicdn.com/imgextra/i4/2207651980344/O1CN01DQZzo91EPatK9E2Rc_!!2207651980344.jpg_430x430q90.jpg', 'Apple');

-- ----------------------------
-- Triggers structure for table goods
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_subGoods`;
delimiter ;;
CREATE TRIGGER `delete_subGoods` BEFORE DELETE ON `goods` FOR EACH ROW delete from sub_goods where goods_id = id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table user
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_address`;
delimiter ;;
CREATE TRIGGER `delete_address` AFTER DELETE ON `user` FOR EACH ROW delete from address where user_id = id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table user
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_cart`;
delimiter ;;
CREATE TRIGGER `delete_cart` AFTER DELETE ON `user` FOR EACH ROW delete from cart where user_id = id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table user
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_order`;
delimiter ;;
CREATE TRIGGER `delete_order` AFTER DELETE ON `user` FOR EACH ROW delete from `order` where user_id = id
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;

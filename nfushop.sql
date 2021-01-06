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

 Date: 06/01/2021 22:39:49
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
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `for_2`(`user_id`) USING BTREE,
  CONSTRAINT `for_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 37 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of address
-- ----------------------------
INSERT INTO `address` VALUES (34, '陈华泽', 'M', '13078255125', '东35', '2021-01-03 16:31:46', '2021-01-03 16:31:46', 1);

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
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `for_7`(`sub_goods_id`) USING BTREE,
  CONSTRAINT `for_7` FOREIGN KEY (`sub_goods_id`) REFERENCES `sub_goods` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf32 COLLATE = utf32_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of banner
-- ----------------------------
INSERT INTO `banner` VALUES (1, 'https://img.alicdn.com/imgextra/i4/268451883/O1CN019QSe6S1PmSRMOdLhS_!!268451883.jpg', 27, '2020-12-30 10:25:21', '2021-01-06 22:35:15');

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
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `for_5`(`user_id`) USING BTREE,
  INDEX `for_6`(`sub_goods_id`) USING BTREE,
  CONSTRAINT `for_5` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `for_6` FOREIGN KEY (`sub_goods_id`) REFERENCES `sub_goods` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 79 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cart
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES (12, 'realme 真我Q21', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/534cadd3-61b2-49a5-b301-b07e5628c936.jpg\"]', '[\n  {\"title\":\"版本\",\"values\":[\"4GB+128GB\",\"6GB+128GB\"]}\n]', '2021-01-03 15:55:58', '2021-01-05 12:49:02', '4800万像素 120Hz畅速屏', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/10b777c0-44ac-43a1-95fa-b1a157c0dbdb.jpg\"]');
INSERT INTO `goods` VALUES (14, '小米11', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/e89d6b18-3f01-4d07-bfe5-f28a45a49d4c.jpg\"]', '[\n  {\"title\":\"版本\",\"values\":[\"8GB+128GB\",\"8GB+256GB\"]}\n]', '2021-01-03 16:53:25', '2021-01-03 16:55:56', '5G 骁龙888 2K AMOLED四曲面柔性屏 1亿像素 8GB+128GB', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/af5d4b02-1223-4603-856a-eef06275fc77.jpg\"]');
INSERT INTO `goods` VALUES (15, '1', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/fc7ea6fb-b2e1-464c-b40d-c2690226cf5e.jpg\"]', '[\n  {\"title\":\"版本\",\"values\":[\"8GB+128GB\"]}\n]', '2021-01-05 17:32:34', '2021-01-05 17:32:34', '2', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/299ec37b-e552-4761-ac5e-a04f46f7433d.jpg\"]');
INSERT INTO `goods` VALUES (16, '1', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/30852b14-1fc9-4069-acd7-4d80f6d62aca.jpg\"]', '[\n  {\"title\":\"版本\",\"values\":[\"8GB+128GB\",\"8GB+256GB\"]}\n]', '2021-01-05 17:33:14', '2021-01-05 17:33:35', '2', '[\"https://nfu-shop.oss-cn-beijing.aliyuncs.com/5852a374-619c-41ea-98d8-1313aa9075b1.jpg\"]');

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
  `delivery_company` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `amount` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `for_3`(`user_id`) USING BTREE,
  CONSTRAINT `for_3` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 84 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (73, 1, '陈华泽', 'M', '13078255125', '东35', '[27]', 297, '2021-01-06 15:36:48', '2021-01-06 17:25:42', 2, '123', '顺丰', '[9]');
INSERT INTO `order` VALUES (74, 1, '陈华泽', 'M', '13078255125', '东35', '[27]', 33, '2021-01-06 15:38:48', '2021-01-06 17:26:21', 2, '25', '东风快递', '[1]');
INSERT INTO `order` VALUES (79, 1, '陈华泽', 'M', '13078255125', '东35', '[25]', 5994, '2021-01-06 17:33:57', '2021-01-06 17:34:52', 2, '东风-25', '东风快递', '[6]');
INSERT INTO `order` VALUES (81, 1, '陈华泽', 'M', '13078255125', '东35', '[25]', 25974, '2021-01-06 17:40:52', '2021-01-06 18:13:24', 2, '112', '顺', '[26]');

-- ----------------------------
-- Table structure for sub_goods
-- ----------------------------
DROP TABLE IF EXISTS `sub_goods`;
CREATE TABLE `sub_goods`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `price` decimal(10, 0) NOT NULL,
  `stoke` decimal(10, 0) UNSIGNED NOT NULL,
  `sell` decimal(10, 0) UNSIGNED NOT NULL,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `goods_id` int(11) NOT NULL,
  `template` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `f`(`goods_id`, `template`) USING BTREE,
  CONSTRAINT `for` FOREIGN KEY (`goods_id`) REFERENCES `goods` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sub_goods
-- ----------------------------
INSERT INTO `sub_goods` VALUES (25, 999, 68, 32, 'https://nfu-shop.oss-cn-beijing.aliyuncs.com/4c7724af-ae06-4b6b-81d9-54c85bf13cdd.jpg', 14, '[0]', '2021-01-06 11:27:04', '2021-01-06 17:55:37');
INSERT INTO `sub_goods` VALUES (26, 888, 10, 0, 'https://nfu-shop.oss-cn-beijing.aliyuncs.com/6566030a-7277-42c7-8c99-ca1fe64e0bd9.jpg', 14, '[1]', '2021-01-06 11:27:27', '2021-01-06 19:14:16');
INSERT INTO `sub_goods` VALUES (27, 33, 0, 10, 'https://nfu-shop.oss-cn-beijing.aliyuncs.com/04c560a6-59f9-40d0-9566-36c77be36f93.jpg', 12, '[1]', '2021-01-06 11:35:37', '2021-01-06 15:38:48');

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
  `money` decimal(10, 0) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `phone`(`phone`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '1', '1', '2020-12-09 23:03:55', '2021-01-06 18:15:33', 'https://i1.hdslb.com/bfs/face/b6dad37ef0a68341b6e2f84de5e1f9ee02f8365e.jpg@128w_128h_1o.webp', 'Apple', NULL);

-- ----------------------------
-- Triggers structure for table goods
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_subGoods`;
delimiter ;;
CREATE TRIGGER `delete_subGoods` BEFORE DELETE ON `goods` FOR EACH ROW delete from sub_goods where sub_goods.goods_id = old.id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table user
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_address`;
delimiter ;;
CREATE TRIGGER `delete_address` BEFORE DELETE ON `user` FOR EACH ROW BEGIN
delete from address where address.user_id = old.id;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table user
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_order`;
delimiter ;;
CREATE TRIGGER `delete_order` BEFORE DELETE ON `user` FOR EACH ROW BEGIN
delete from `order` where user_id = old.id;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table user
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_cart`;
delimiter ;;
CREATE TRIGGER `delete_cart` BEFORE DELETE ON `user` FOR EACH ROW BEGIN
delete from cart where user_id = old.id;
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;

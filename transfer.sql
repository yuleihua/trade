/*
 Navicat Premium Data Transfer

 Source Server         : 172.16.1.172-dev
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : 172.16.1.172:3306
 Source Schema         : tradedb

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 31/08/2020 22:52:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for branch
-- ----------------------------
DROP TABLE IF EXISTS `branch`;
CREATE TABLE `branch`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `cid` bigint(20) NOT NULL DEFAULT 0,
  `short_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '简称',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名字',
  `bankid` bigint(20) NOT NULL DEFAULT 0 COMMENT '银行编号',
  `money_type` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户',
  `status` int(10) NOT NULL DEFAULT 0,
  `is_master` tinyint(1) NOT NULL DEFAULT 0,
  `created` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `branch_idx1`(`cid`, `account`) USING BTREE,
  UNIQUE INDEX `branch_idx2`(`cid`, `account`) USING BTREE,
  INDEX `branch_idx3`(`name`) USING BTREE,
  INDEX `branch_idx4`(`short_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of branch
-- ----------------------------
INSERT INTO `branch` VALUES (1, 1, 'test-com', 'yulei-test ', 888888, 'CNY', '6c84c4b8c2f1c1a15feb6967578049bb', 0, 1, '2020-08-27 23:24:31');
INSERT INTO `branch` VALUES (2, 2, 'one-com', 'some one', 666666, 'CNY', '6cbc0bf671915aa36970ade080f4c474', 0, 1, '2020-08-27 23:26:47');

-- ----------------------------
-- Table structure for customer
-- ----------------------------
DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名字',
  `nation` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '国家',
  `city` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址',
  `phone` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电话',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `remark` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `customer_idx1`(`name`, `phone`) USING BTREE,
  UNIQUE INDEX `customer_idx2`(`name`, `email`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of customer
-- ----------------------------
INSERT INTO `customer` VALUES (1, 'yulei', 'CN', 'Shanghai', 'pudong', '13181782128', 'test@aa.com', 'ray', '2020-08-27 21:45:20');
INSERT INTO `customer` VALUES (2, 'toone', 'CN', 'Beijing', 'fengtai', '15917828911', 'test@sss.com', 'to', '2020-08-27 23:25:15');

-- ----------------------------
-- Table structure for fund
-- ----------------------------
DROP TABLE IF EXISTS `fund`;
CREATE TABLE `fund`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `bid` bigint(20) NOT NULL DEFAULT 0,
  `money_type` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '源IP地址',
  `balance` bigint(20) NOT NULL DEFAULT 0,
  `freeze_balance` bigint(20) NOT NULL DEFAULT 0,
  `last_balance` bigint(20) NOT NULL DEFAULT 0,
  `created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `fund_idx1`(`bid`, `money_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of fund
-- ----------------------------
INSERT INTO `fund` VALUES (1, 1, 'CNY', 9999000000000, 0, 0, '2020-08-27 23:28:30');
INSERT INTO `fund` VALUES (2, 2, 'CNY', 150000, 0, 0, '2020-08-27 23:29:11');

-- ----------------------------
-- Table structure for trade
-- ----------------------------
DROP TABLE IF EXISTS `trade`;
CREATE TABLE `trade`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `trans_date` int(11) NOT NULL DEFAULT 0,
  `trans_time` int(11) NOT NULL DEFAULT 0,
  `uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '流水号',
  `from_uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '流水号',
  `to_uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '流水号',
  `from_bid` bigint(20) NOT NULL DEFAULT 0,
  `from_cid` bigint(20) NOT NULL DEFAULT 0,
  `to_bid` bigint(20) NOT NULL DEFAULT 0,
  `to_cid` bigint(20) NOT NULL DEFAULT 0,
  `is_delay` tinyint(1) NOT NULL DEFAULT 0,
  `is_large` tinyint(1) NOT NULL DEFAULT 0,
  `is_reject` tinyint(1) NOT NULL DEFAULT 0,
  `amt` bigint(20) NOT NULL DEFAULT 0,
  `fee` bigint(20) NOT NULL DEFAULT 0,
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `money_type` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '币种',
  `errcode` int(10) NOT NULL DEFAULT 0,
  `confirm_date` int(11) NOT NULL DEFAULT 0,
  `confirm_time` int(11) NOT NULL DEFAULT 0,
  `confirm_money_type` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `confirm_amt` bigint(20) NOT NULL DEFAULT 0,
  `confirm_opid` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `trade_idx1`(`uuid`) USING BTREE,
  UNIQUE INDEX `trade_idx2`(`uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_fee
-- ----------------------------
DROP TABLE IF EXISTS `trade_fee`;
CREATE TABLE `trade_fee`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `trans_date` int(11) NOT NULL DEFAULT 0,
  `trans_time` int(11) NOT NULL DEFAULT 0,
  `uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '流水号',
  `money_type` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '币种',
  `amt` bigint(20) NOT NULL DEFAULT 0,
  `fee` bigint(20) NOT NULL DEFAULT 0,
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `trade_fee_idx1`(`uuid`) USING BTREE,
  INDEX `trade_fee_idx2`(`trans_date`, `money_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for trade_risk
-- ----------------------------
DROP TABLE IF EXISTS `trade_risk`;
CREATE TABLE `trade_risk`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `trans_date` int(11) NOT NULL DEFAULT 0,
  `trans_time` int(11) NOT NULL DEFAULT 0,
  `level` int(11) NOT NULL DEFAULT 0,
  `uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '流水号',
  `amt` bigint(20) NOT NULL DEFAULT 0,
  `fee` bigint(20) NOT NULL DEFAULT 0,
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `trade_risk_idx1`(`uuid`) USING BTREE,
  INDEX `trade_risk_idx2`(`trans_date`, `level`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for trade_seq
-- ----------------------------
DROP TABLE IF EXISTS `trade_seq`;
CREATE TABLE `trade_seq`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `trans_date` int(11) NOT NULL DEFAULT 0,
  `trans_time` int(11) NOT NULL DEFAULT 0,
  `from_uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '流水号',
  `from_bid` bigint(20) NOT NULL DEFAULT 0,
  `from_cid` bigint(20) NOT NULL DEFAULT 0,
  `to_uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `to_bid` bigint(20) NOT NULL DEFAULT 0,
  `uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `errcode` int(4) UNSIGNED ZEROFILL NOT NULL DEFAULT 0000,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `trade_seq_idx1`(`from_uuid`) USING BTREE,
  INDEX `trade_seq_idx2`(`from_bid`, `trans_date`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

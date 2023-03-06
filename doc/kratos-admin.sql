/*
 Navicat Premium Data Transfer

 Source Server         : docker-mysql
 Source Server Type    : MySQL
 Source Server Version : 50741
 Source Host           : localhost:3307
 Source Schema         : kratos-admin

 Target Server Type    : MySQL
 Target Server Version : 50741
 File Encoding         : 65001

 Date: 06/03/2023 16:13:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_admin
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin`;
CREATE TABLE `sys_admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` char(36) NOT NULL DEFAULT '' COMMENT 'UID',
  `realname` varchar(100) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `mobile` varchar(15) NOT NULL DEFAULT '' COMMENT '电话号',
  `area_code` smallint(4) NOT NULL DEFAULT '86' COMMENT '区号',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'EMAIL',
  `weixin` varchar(30) NOT NULL DEFAULT '' COMMENT '微信号',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后一次登录时间',
  `operator` char(36) NOT NULL DEFAULT '' COMMENT '操作人',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uniq_uid` (`uid`) USING BTREE,
  UNIQUE KEY `uniq_mobile` (`mobile`,`area_code`,`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- ----------------------------
-- Table structure for sys_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin_role`;
CREATE TABLE `sys_admin_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` char(36) NOT NULL DEFAULT '' COMMENT 'UID',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_uid_role` (`uid`,`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员用户角色表';

-- ----------------------------
-- Table structure for sys_casbin_policy
-- ----------------------------
DROP TABLE IF EXISTS `sys_casbin_policy`;
CREATE TABLE `sys_casbin_policy` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` char(1) NOT NULL DEFAULT '' COMMENT '`p`或者`g`',
  `v0` varchar(40) NOT NULL DEFAULT '' COMMENT '角色ID或用户ID',
  `v1` varchar(16) NOT NULL DEFAULT '' COMMENT '菜单ID或角色ID',
  `v2` varchar(8) NOT NULL DEFAULT '' COMMENT 'act值 C,U,R,D 或 *',
  `v3` varchar(8) NOT NULL DEFAULT '' COMMENT 'deny或allow',
  `v4` varchar(8) NOT NULL DEFAULT '' COMMENT '保留字段',
  `v5` varchar(8) NOT NULL DEFAULT '' COMMENT '保留字段',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uniq_index` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='casbin policy表';

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '菜单标题',
  `i18n` varchar(200) NOT NULL COMMENT 'I18n 文案',
  `route` varchar(255) NOT NULL DEFAULT '' COMMENT '菜单路由',
  `api_url` varchar(255) NOT NULL DEFAULT '' COMMENT 'api路径',
  `method` char(7) NOT NULL DEFAULT '' COMMENT 'api请求方式',
  `icon` varchar(20) NOT NULL DEFAULT '' COMMENT '图标标识',
  `is_hide` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否隐藏',
  `jump_url` varchar(255) NOT NULL COMMENT '跳转地址',
  `jump_mode` tinyint(1) NOT NULL DEFAULT '0' COMMENT '跳转方式 0:当前页面跳转 1:新标签页面打开 2:内嵌页面打开',
  `operator` char(36) NOT NULL DEFAULT '' COMMENT '操作人',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(60) NOT NULL DEFAULT '' COMMENT '角色名',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '角色描述',
  `is_enable` tinyint(1) NOT NULL DEFAULT '0' COMMENT '启用状态 1:启用 0:未启用',
  `operator` char(36) NOT NULL DEFAULT '' COMMENT '操作人',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` char(36) NOT NULL DEFAULT '' COMMENT 'UID',
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `realname` varchar(100) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `mobile` varchar(15) NOT NULL DEFAULT '' COMMENT '电话号',
  `area_code` smallint(4) NOT NULL DEFAULT '86' COMMENT '区号',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'EMAIL',
  `weixin` varchar(30) NOT NULL DEFAULT '' COMMENT '微信号',
  `operator` char(36) NOT NULL DEFAULT '' COMMENT '操作人',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_uid` (`uid`) USING BTREE,
  UNIQUE KEY `uni_mobile` (`mobile`,`area_code`,`deleted_at`) USING BTREE,
  UNIQUE KEY `uni_username` (`username`,`deleted_at`) USING BTREE,
  UNIQUE KEY `uni_email` (`email`,`deleted_at`) USING BTREE,
  KEY `index_realname` (`realname`(3)) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

SET FOREIGN_KEY_CHECKS = 1;

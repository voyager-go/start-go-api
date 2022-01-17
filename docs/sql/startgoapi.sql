/*
 Navicat Premium Data Transfer

 Source Server         : mysql8
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : startgoapi

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 17/01/2022 15:43:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES (4, 'p', '1', '/api/sys_menu', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (3, 'p', '1', '/api/v1/delete1', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (2, 'p', '1', '/api/v1/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (5, 'p', '2', '/api/sys_menu', 'POST', NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `description` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `content` longtext COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `status` tinyint(1) NOT NULL COMMENT '是否启用 0:禁用 1:启用',
  `is_publish` tinyint(1) NOT NULL COMMENT '是否发布 0:否 1:是',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `status` (`status`) USING BTREE COMMENT '是否可用的筛选条件',
  KEY `is_publish` (`is_publish`) USING BTREE COMMENT '是否发布的筛选条件'
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='博文信息表';

-- ----------------------------
-- Records of post
-- ----------------------------
BEGIN;
INSERT INTO `post` VALUES (5, 'test', 'kkkkk', 'sdadsadsjdadk', 1, 1, '2011-11-11 02:00:00', NULL);
INSERT INTO `post` VALUES (6, 'foo hello', 'foo ', 'fooolslal', 1, 1, '2012-11-22 01:00:00', NULL);
INSERT INTO `post` VALUES (7, 'bar', 'bar', 'aoadjdaldm', 0, 0, '2022-01-04 14:37:25', NULL);
INSERT INTO `post` VALUES (8, 'hello', 'hello', 'dsajdakan', 1, 1, '2022-01-03 14:37:36', NULL);
INSERT INTO `post` VALUES (9, 'niu', 'nsdaj', 'daslkdasl', 0, 1, NULL, NULL);
INSERT INTO `post` VALUES (10, 'world hello ', 'sdsa', 'daads', 1, 1, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `path` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'api访问路径',
  `description` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api中文描述',
  `group` varchar(30) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'api所在分组',
  `method` tinyint(1) NOT NULL COMMENT '方法 1:创建POST 2:查看GET 3:更新PUT 4:删除DELETE',
  `is_use` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否可用 0禁用 1启用',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `path` (`path`,`method`,`group`) USING BTREE,
  KEY `is_use` (`is_use`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='api信息';

-- ----------------------------
-- Records of sys_api
-- ----------------------------
BEGIN;
INSERT INTO `sys_api` VALUES (1, '/user', '创建用户', 'api', 1, 0, '2022-01-12 16:33:57', '2022-01-12 16:33:57');
INSERT INTO `sys_api` VALUES (9, '/user', '更新用户', 'api', 3, 0, '2022-01-12 17:31:18', '2022-01-12 17:31:18');
INSERT INTO `sys_api` VALUES (10, '/api/sys_menu', '创建菜单', 'api', 1, 1, '2022-01-13 15:16:55', '2022-01-13 15:16:55');
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(30) COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `pid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父级ID',
  `is_use` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否可用 默认1表示启用 0表示禁用',
  `level` tinyint(1) NOT NULL DEFAULT '1' COMMENT '菜单级别 默认是1级菜单',
  `sort` tinyint(1) NOT NULL COMMENT '排序编号',
  `icon` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '图标ICON',
  `unique_key` varchar(30) COLLATE utf8mb4_general_ci NOT NULL COMMENT '唯一描述',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_key` (`unique_key`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE,
  KEY `is_use` (`is_use`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单表';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` VALUES (1, '用户管理', 0, 1, 1, 2, '', 'user-manager', '2022-01-10 19:01:15', '2022-01-10 19:01:15');
INSERT INTO `sys_menu` VALUES (2, '用户列表', 1, 1, 2, 1, '', 'user-list', '2022-01-10 19:01:45', '2022-01-10 19:01:45');
INSERT INTO `sys_menu` VALUES (3, '用户信息', 1, 1, 2, 2, '', 'user-info', '2022-01-10 19:02:08', '2022-01-10 19:02:08');
INSERT INTO `sys_menu` VALUES (4, '用户绑定信息', 1, 1, 2, 3, '', 'user-binding', '2022-01-10 19:02:38', '2022-01-10 19:02:38');
INSERT INTO `sys_menu` VALUES (5, '文章管理', 0, 1, 1, 1, '', 'ppost-manager', '2022-01-10 19:03:35', '2022-01-10 19:03:35');
INSERT INTO `sys_menu` VALUES (6, '文章信息', 5, 1, 2, 1, '', 'post-info', '2022-01-10 19:05:27', '2022-01-10 19:05:27');
INSERT INTO `sys_menu` VALUES (7, '文章阅读', 5, 1, 2, 14, '', 'post-show', '2022-01-10 19:28:37', '2022-01-10 19:28:37');
INSERT INTO `sys_menu` VALUES (9, '文章阅读2', 5, 1, 2, 14, '', 'post-show-2', '2022-01-10 19:29:54', '2022-01-10 19:29:54');
INSERT INTO `sys_menu` VALUES (11, '文章阅读3', 5, 1, 2, 14, '', 'post-show-3', '2022-01-10 19:31:01', '2022-01-10 19:31:01');
INSERT INTO `sys_menu` VALUES (15, '用户阅读4', 0, 1, 2, 14, '', 'user-show-4', '2022-01-10 19:36:24', '2022-01-10 19:36:24');
INSERT INTO `sys_menu` VALUES (16, '用户阅读5', 0, 0, 2, 14, '', 'user-show-5', '2022-01-10 19:41:06', '2022-01-10 19:41:06');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `pid` bigint unsigned NOT NULL COMMENT '父级ID',
  `is_use` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否可用 默认1表示启用 0表示禁用',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE,
  KEY `is_use` (`is_use`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` VALUES (1, '超级管理员', 0, 0, '2022-01-14 17:41:43', '2022-01-14 17:41:43');
INSERT INTO `sys_role` VALUES (2, '一级管理员', 1, 1, '2022-01-14 17:43:57', '2022-01-14 17:43:57');
INSERT INTO `sys_role` VALUES (3, '运营管理员', 1, 1, '2022-01-14 17:44:33', '2022-01-14 17:44:33');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `menu_id` bigint unsigned NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_id` (`role_id`,`menu_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单角色关联表';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_role_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_user`;
CREATE TABLE `sys_role_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` bigint unsigned NOT NULL COMMENT '角色编号',
  `user_id` bigint unsigned NOT NULL COMMENT '用户编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_id` (`role_id`,`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色与用户关联表';

-- ----------------------------
-- Records of sys_role_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_user` VALUES (1, 2, 1);
INSERT INTO `sys_role_user` VALUES (2, 3, 1);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `nickname` varchar(80) COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `phone` varchar(11) COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号',
  `password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `status` tinyint(1) NOT NULL COMMENT '是否启用 0:禁用 1:启用',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`) USING BTREE COMMENT '手机号唯一',
  KEY `status` (`status`) USING BTREE COMMENT '是否可用的筛选条件'
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, '', '15105191181', '$2a$10$Q51vH8mHSNHQimg9GeMH9O.HCSlz36iJpeDfjssCzMmrYFg1cnYCW', 1, '2022-01-06 16:30:43', '2022-01-06 16:41:06');
INSERT INTO `user` VALUES (2, '天天开心', '15105191182', '', 1, '2022-01-05 16:30:46', NULL);
INSERT INTO `user` VALUES (3, '今天是个好日子', '15105191183', '$2a$10$Q51vH8mHSNHQimg9GeMH9O.HCSlz36iJpeDfjssCzMmrYFg1cnYCW', 1, '2022-01-11 16:30:50', NULL);
INSERT INTO `user` VALUES (4, '李四', '15105191185', '$2a$10$DEeSpXz4d.yPhNYyfYU92OtNQUumM7IHyWxrCfYSzohDvw1WIagN6', 1, '2022-01-03 16:30:54', '2022-01-06 16:48:44');
INSERT INTO `user` VALUES (6, '张小三', '15105191192', '$2a$10$a.vt8cAZguPr/4PPFPZf.u7jykATJS6AG8lOGlNgU8tXTskcmvZ8y', 1, '2022-01-06 15:47:49', '2022-01-06 15:47:49');
INSERT INTO `user` VALUES (7, '张小三', '15105191191', '$2a$10$Ou88iY492z/My/teOmWo5uM3CL9k97IxiKDwlpoDhtZyLSWA6sW.W', 1, '2022-01-06 15:48:41', '2022-01-06 15:48:41');
INSERT INTO `user` VALUES (8, '张小三', '15105191193', '$2a$10$U0uUn/g7lp5Jc7AlWAMJlOEROuO55SJEYlQ2fq.9Ch88UnrX.VT6O', 1, '2022-01-07 09:42:47', '2022-01-07 09:42:47');
INSERT INTO `user` VALUES (9, '张小三', '15105191194', '$2a$10$z1zbuV06yZsN4tZqv/Auou5LX2RVoWzTTLJBY3SejTaMJJkHoAXCe', 1, '2022-01-07 09:43:52', '2022-01-07 09:43:52');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

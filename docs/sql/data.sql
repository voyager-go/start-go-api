CREATE DATABASE startgoapi;

use startgoapi;

CREATE TABLE `sys_user` (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `nickname` varchar(80) COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
    `phone` varchar(11) COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号',
    `password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
    `status` tinyint(1) NOT NULL COMMENT '是否启用 0:禁用 1:启用',
    `created_at` int unsigned DEFAULT NULL COMMENT '创建时间',
    `updated_at` int DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `phone` (`phone`) USING BTREE COMMENT '手机号唯一',
    KEY `status` (`status`) USING BTREE COMMENT '是否可用的筛选条件'
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';
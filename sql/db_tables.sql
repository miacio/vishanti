CREATE TABLE `user_account_info` (
  `id` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键uuid',
  `mobile` varchar(11) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '手机号',
  `email` varchar(256) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '邮箱',
  `account` varchar(32) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '账号',
  `password` varchar(32) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '密码',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `status` varchar(32) COLLATE utf8mb4_general_ci DEFAULT '1' COMMENT '账号状态: [USER_ACCESS_STATUS]',
  `lock_time` datetime DEFAULT NULL COMMENT '封号时间: 封号的到期时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户账号表';

CREATE TABLE `user_detailed_info` (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键uuid',
  `user_account_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '用户账号id',
  `vip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '1' COMMENT 'vip类别: [USER_VIP]',
  `head_pic_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '用户头像 - 文件id',
  `nick_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '用户昵称',
  `sex` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '0' COMMENT '用户性别: 0 未知 1 男 2 女',
  `birthday_year` int DEFAULT '0' COMMENT '用户生日-年',
  `birthday_month` int DEFAULT '0' COMMENT '用户生日-月',
  `birthday_day` int DEFAULT '0' COMMENT '用户生日-日',
  `profile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '个人简介',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';

CREATE TABLE `system_dictionary` (
  `id` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键uuid',
  `name` varchar(32) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '名称',
  `group` varchar(32) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '组名',
  `parent_group` varchar(32) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '上级组名',
  `describe` varchar(32) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '描述',
  `val` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '值',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `create_by` varchar(32) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '创建人id',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `update_by` varchar(32) DEFAULT '' COMMENT '修改人id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典表';

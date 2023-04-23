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

CREATE TABLE `system_file_info` (
  `id` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件id',
  `file_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件名',
  `object_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'minio object name',
  `region` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'minio region',
  `bucket` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'minio bucket',
  `file_size` bigint DEFAULT '0' COMMENT '文件大小',
  `file_md5` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件md5',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `create_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '上传者id',
  `used` int DEFAULT '0' COMMENT '是否使用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文件表';

CREATE TABLE `circles_info` (
  `id` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '圈子id',
  `logo` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '圈子logo id',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '圈子名称',
  `descirbe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '圈子描述',
  `create_time` datetime DEFAULT NULL COMMENT '圈子的创建时间',
  `create_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '圈子的创建者id',
  `owner` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '圈子的所有者id',
  `update_time` datetime DEFAULT NULL COMMENT '圈子的修改时间',
  `update_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '圈子的修改者id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='圈子主表';

CREATE TABLE `circles_server_info` (
  `id` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务信息id',
  `circles_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '圈子id',
  `server_url` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '圈子服务器地址',
  `public_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '公钥',
  `secret_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '私钥',
  `down` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '通信是否宕机 字典: [CIRCLES_SERVER_DOWN]',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `create_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '创建者id',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `update_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '修改者id',
  `used` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '是由启用 字典: [CIRCLES_SERVER_USED]',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='圈子三方服务信息';
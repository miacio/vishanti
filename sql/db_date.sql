-- 字典 用户状态
insert into system_dictionary (id, `name`, `group`, parent_group, `describe`, val, create_time, create_by) values 
(UPPER(REPLACE(UUID(),'-', '')), '正常', 'USER_ACCESS_STATUS', 'USER', '用户状态', '1', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '封号', 'USER_ACCESS_STATUS', 'USER', '用户状态', '2', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '注销中', 'USER_ACCESS_STATUS', 'USER', '用户状态', '3', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '注销', 'USER_ACCESS_STATUS', 'USER', '用户状态', '4', NOW(), '');

-- 字典 会员类别
insert into system_dictionary (id, `name`, `group`, parent_group, `describe`, val, create_time, create_by) values 
(UPPER(REPLACE(UUID(),'-', '')), '普通用户', 'USER_VIP', 'USER', '用户VIP类别', '1', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), 'VIP用户', 'USER_VIP', 'USER', '用户VIP类别', '2', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), 'SVIP用户', 'USER_VIP', 'USER', '用户VIP类别', '3', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '钻石VIP用户', 'USER_VIP', 'USER', '用户VIP类别', '4', NOW(), '');

--字典 会员对应圈子数
insert into system_dictionary (id, `name`, `group`, parent_group, `describe`, val, create_time, create_by) values 
(UPPER(REPLACE(UUID(),'-', '')), '1', 'USER_VIP_CIRCLES', 'CIRCLES', '会员等级对应圈子数', '1', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '2', 'USER_VIP_CIRCLES', 'CIRCLES', '会员等级对应圈子数', '2', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '3', 'USER_VIP_CIRCLES', 'CIRCLES', '会员等级对应圈子数', '3', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '4', 'USER_VIP_CIRCLES', 'CIRCLES', '会员等级对应圈子数', '-1', NOW(), '');

--字典 是否开启三方圈子服务
insert into system_dictionary (id, `name`, `group`, parent_group, `describe`, val, create_time, create_by) values 
(UPPER(REPLACE(UUID(),'-', '')), '启用', 'CIRCLES_SERVER_USED', 'CIRCLES', '圈子服务是否开启', '1', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '停用', 'CIRCLES_SERVER_USED', 'CIRCLES', '圈子服务是否开启', '2', NOW(), '');

--字典 当前用户所在圈子的用户状态
insert into system_dictionary (id, `name`, `group`, parent_group, `describe`, val, create_time, create_by) values 
(UPPER(REPLACE(UUID(),'-', '')), '邀请中', 'CIRCLES_SIGN_OUT', 'CIRCLES', '当前用户所在圈子的用户状态', '1', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '申请中', 'CIRCLES_SIGN_OUT', 'CIRCLES', '当前用户所在圈子的用户状态', '2', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '登记', 'CIRCLES_SIGN_OUT', 'CIRCLES', '当前用户所在圈子的用户状态', '3', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '退出', 'CIRCLES_SIGN_OUT', 'CIRCLES', '当前用户所在圈子的用户状态', '4', NOW(), '');


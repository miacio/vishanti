-- 字典 用户状态
insert into system_dictionary (id, `name`, `group`, parent_group, `describe`, val, create_time, create_by) values 
(UPPER(REPLACE(UUID(),'-', '')), '正常', 'USER_ACCESS_STATUS', '', '用户状态', '1', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '封号', 'USER_ACCESS_STATUS', '', '用户状态', '2', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '注销中', 'USER_ACCESS_STATUS', '', '用户状态', '3', NOW(), ''),
(UPPER(REPLACE(UUID(),'-', '')), '注销', 'USER_ACCESS_STATUS', '', '用户状态', '4', NOW(), '');

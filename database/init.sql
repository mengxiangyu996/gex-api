-- 用户表
create table if not exists byyfuncamera.`user` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `type` tinyint(3) default 0 comment '类型：0-用户；1-管理员',
    `username` varchar(128) comment '用户名',
    `nickname` varchar(128) comment '昵称',
    `gender` tinyint(3) comment '性别：0-女；1-男',
    `email` varchar(128) comment '邮箱',
    `phone` varchar(128) comment '电话',
    `password` varchar(256) comment '密码',
    `avatar` varchar(256) comment '头像',
    `wx_open_id` varchar(128) comment '微信openid',
    `wx_union_id` varchar(128) comment '微信unionid',
    `status` tinyint(3) default 1 comment '状态：0-禁用；1-启用'
) comment '用户表';

-- 角色表
create table if not exists byyfuncamera.`role` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `name` varchar(128) comment '角色名',
    `status` int(3) default 1 comment '状态：0-禁用；1-启用'
) comment '角色表';

-- 菜单表
create table if not exists byyfuncamera.`menu` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `parent_id` int(10) comment '父级id',
    `name` varchar(128) comment '菜单名',
    `type` tinyint(3) comment '类型：1-目录；2-菜单；3-按钮',
    `sort` int(10) comment '排序',
    `path` varchar(128) default '1' comment '权限地址',
    `component` varchar(128) comment '组件',
    `icon` varchar(128) comment '图标',
    `redirect` varchar(128) comment '重定向',
    `status` tinyint(3) default 1 comment '状态：0-禁用；1-启用'
) comment '菜单表';

-- 权限表
create table if not exists byyfuncamera.`permission` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `name` varchar(128) comment '权限名',
    `group_name` varchar(128) comment '组名',
    `path` varchar(128) default '1' comment '权限地址',
    `method` varchar(128) comment '请求方式',
    `status` tinyint(3) default 1 comment '状态：0-禁用；1-启用'
) comment '权限表';

-- 用户角色关系表
create table if not exists byyfuncamera.`user_role_relation` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `user_id` int(10) comment '用户id',
    `role_id` int(10) comment '角色id'
) comment '用户角色关系表';

-- 角色菜单关系表
create table if not exists byyfuncamera.`role_permission_relation` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `role_id` int(10) comment '角色id',
    `menu_id` int(10) comment '菜单id'
) comment '角色菜单关系表';

-- 角色权限关系表
create table if not exists byyfuncamera.`role_permission_relation` (
    `id` int(10) auto_increment comment 'id' primary key,
    `create_time` datetime comment '创建时间',
    `update_time` datetime comment '更新时间',
    `delete_time` datetime comment '删除时间',
    `role_id` int(10) comment '角色id',
    `permission_id` int(10) comment '权限id'
) comment '角色权限关系表';
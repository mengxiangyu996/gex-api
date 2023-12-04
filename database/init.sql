-- 用户表
create table if not exists breeze-api.`users`
(
`id` int null auto_increment comment 'id' primary key,
`create_time` datetime null comment '创建时间',
`update_time` datetime null comment '更新时间',
`delete_time` datetime null comment '删除时间',
`type` tinyint default 0 null comment '类型：0-用户；1-管理员',
`username` varchar(256) null comment '用户名',
`nickname` varchar(256) null comment '昵称',
`gender` tinyint null comment '性别：0-女；1-男',
`email` varchar(256) null comment '邮箱',
`phone` varchar(256) null comment '电话',
`password` varchar(256) null comment '密码',
`avatar` varchar(256) null comment '头像',
`wx_open_id` varchar(256) null comment '微信openid',
`wx_union_id` varchar(256) null comment '微信unionid',
`status` tinyint default 1 null comment '状态：0-禁用；1-启用'
) comment '用户表';

-- 角色表
create table if not exists breeze-api.`roles`
(
`id` int null auto_increment comment 'id' primary key,
`create_time` datetime null comment '创建时间',
`update_time` datetime null comment '更新时间',
`delete_time` datetime null comment '删除时间',
`name` varchar(256) null comment '角色名',
`status` int default 1 null comment '状态：0-禁用；1-启用'
) comment '角色表';

-- 菜单表
create table if not exists breeze.`menus`
(
`id` int null auto_increment comment 'id' primary key,
`create_time` datetime null comment '创建时间',
`update_time` datetime null comment '更新时间',
`delete_time` datetime null comment '删除时间',
`parent_id` int null comment '父级id',
`name` varchar(256) null comment '菜单名',
`type` tinyint null comment '类型：1-目录；2-菜单；3-按钮',
`sort` int null comment '排序',
`path` varchar(256) default '1' null comment '权限地址',
`component` varchar(256) null comment '组件',
`icon` varchar(256) null comment '图标',
`redirect` varchar(256) null comment '重定向',
`status` tinyint default 1 null comment '状态：0-禁用；1-启用'
) comment '菜单表';

-- 权限表
create table if not exists breeze.`permissions`
(
`id` int null auto_increment comment 'id' primary key,
`create_time` datetime null comment '创建时间',
`update_time` datetime null comment '更新时间',
`delete_time` datetime null comment '删除时间',
`name` varchar(256) null comment '权限名',
`path` varchar(256) default '1' null comment '权限地址',
`method` varchar(256) null comment '请求方式',
`status` tinyint default 1 null comment '状态：0-禁用；1-启用'
) comment '权限表';

-- 用户角色关系表
create table if not exists breeze-api.`user_role_relations`
(
`id` int null auto_increment comment 'id' primary key,
`create_time` datetime null comment '创建时间',
`update_time` datetime null comment '更新时间',
`delete_time` datetime null comment '删除时间',
`user_id` int null comment '用户id',
`role_id` int null comment '角色id'
) comment '用户角色关系表';

-- 角色菜单关系表
create table if not exists breeze-api.`role_permission_relations`
(
`id` int null auto_increment comment 'id' primary key,
`create_time` datetime null comment '创建时间',
`update_time` datetime null comment '更新时间',
`delete_time` datetime null comment '删除时间',
`role_id` int null comment '角色id',
`menu_id` int default 1 null comment '菜单id'
) comment '角色菜单关系表';

-- 角色权限关系表
create table if not exists breeze-api.`role_permission_relations`
(
`id` int null auto_increment comment 'id' primary key,
`create_time` datetime null comment '创建时间',
`update_time` datetime null comment '更新时间',
`delete_time` datetime null comment '删除时间',
`role_id` int null comment '角色id',
`permission_id` int default 1 null comment '权限id'
) comment '角色权限关系表';
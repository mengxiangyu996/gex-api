
-- 导出  表 atomgo.permission 结构
CREATE TABLE IF NOT EXISTS `permission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `parent_id` int DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `layout` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `keep_alive` tinyint DEFAULT NULL,
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `show` tinyint NOT NULL DEFAULT '1' COMMENT '是否展示在页面菜单',
  `enable` tinyint NOT NULL DEFAULT '1',
  `order` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `IDX_30e166e8c6359970755c5727a2` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- 正在导出表  atomgo.permission 的数据：~15 rows (大约)
INSERT INTO `permission` (`id`, `name`, `code`, `type`, `parent_id`, `path`, `redirect`, `icon`, `component`, `layout`, `keep_alive`, `method`, `description`, `show`, `enable`, `order`) VALUES
	(1, '资源管理', 'Resource_Mgt', 'MENU', 2, '/pms/resource', NULL, 'i-fe:list', '/src/views/pms/resource/index.vue', NULL, NULL, NULL, NULL, 1, 1, 1),
	(2, '系统管理', 'SysMgt', 'MENU', NULL, NULL, NULL, 'i-fe:grid', NULL, NULL, NULL, NULL, NULL, 1, 1, 2),
	(3, '角色管理', 'RoleMgt', 'MENU', 2, '/pms/role', NULL, 'i-fe:user-check', '/src/views/pms/role/index.vue', NULL, NULL, NULL, NULL, 1, 1, 2),
	(4, '用户管理', 'UserMgt', 'MENU', 2, '/pms/user', NULL, 'i-fe:user', '/src/views/pms/user/index.vue', NULL, 1, NULL, NULL, 1, 1, 3),
	(5, '分配用户', 'RoleUser', 'MENU', 3, '/pms/role/user/:roleId', NULL, 'i-fe:user-plus', '/src/views/pms/role/role-user.vue', 'full', NULL, NULL, NULL, 0, 1, 1),
	(6, '业务示例', 'Demo', 'MENU', NULL, NULL, NULL, 'i-fe:grid', NULL, NULL, NULL, NULL, NULL, 1, 1, 1),
	(7, '图片上传', 'ImgUpload', 'MENU', 6, '/demo/upload', NULL, 'i-fe:image', '/src/views/demo/upload/index.vue', '', 1, NULL, NULL, 1, 1, 2),
	(8, '个人资料', 'UserProfile', 'MENU', NULL, '/profile', NULL, 'i-fe:user', '/src/views/profile/index.vue', NULL, NULL, NULL, NULL, 0, 1, 99),
	(9, '基础功能', 'Base', 'MENU', NULL, '', NULL, 'i-fe:grid', NULL, '', NULL, NULL, NULL, 1, 1, 0),
	(10, '基础组件', 'BaseComponents', 'MENU', 9, '/base/components', NULL, 'i-me:awesome', '/src/views/base/index.vue', NULL, NULL, NULL, NULL, 1, 1, 1),
	(11, 'Unocss', 'Unocss', 'MENU', 9, '/base/unocss', NULL, 'i-me:awesome', '/src/views/base/unocss.vue', NULL, NULL, NULL, NULL, 1, 1, 2),
	(12, 'KeepAlive', 'KeepAlive', 'MENU', 9, '/base/keep-alive', NULL, 'i-me:awesome', '/src/views/base/keep-alive.vue', NULL, 1, NULL, NULL, 1, 1, 3),
	(13, '创建新用户', 'AddUser', 'BUTTON', 4, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 1, 1, 1),
	(14, '图标 Icon', 'Icon', 'MENU', 9, '/base/icon', NULL, 'i-fe:feather', '/src/views/base/unocss-icon.vue', '', NULL, NULL, NULL, 1, 1, 0),
	(15, 'MeModal', 'TestModal', 'MENU', 9, '/testModal', NULL, 'i-me:dialog', '/src/views/base/test-modal.vue', NULL, NULL, NULL, NULL, 1, 1, 5);

-- 导出  表 atomgo.profile 结构
CREATE TABLE IF NOT EXISTS `profile` (
  `id` int NOT NULL AUTO_INCREMENT,
  `gender` int DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_id` int NOT NULL,
  `nickname` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `IDX_a24972ebd73b106250713dcddd` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- 正在导出表  atomgo.profile 的数据：~1 rows (大约)
INSERT INTO `profile` (`id`, `gender`, `avatar`, `address`, `email`, `user_id`, `nickname`) VALUES
	(1, NULL, 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80', NULL, NULL, 1, 'Admin');

-- 导出  表 atomgo.role 结构
CREATE TABLE IF NOT EXISTS `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `enable` tinyint NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `IDX_ee999bb389d7ac0fd967172c41` (`code`) USING BTREE,
  UNIQUE KEY `IDX_ae4578dcaed5adff96595e6166` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- 正在导出表  atomgo.role 的数据：~2 rows (大约)
INSERT INTO `role` (`id`, `code`, `name`, `enable`) VALUES
	(1, 'SUPER_ADMIN', '超级管理员', 1),
	(2, 'ROLE_QA', '质检员', 1);

-- 导出  表 atomgo.role_permissions_permission 结构
CREATE TABLE IF NOT EXISTS `role_permissions_permission` (
  `role_id` int NOT NULL,
  `permission_id` int NOT NULL,
  PRIMARY KEY (`role_id`,`permission_id`) USING BTREE,
  KEY `IDX_b36cb2e04bc353ca4ede00d87b` (`role_id`) USING BTREE,
  KEY `IDX_bfbc9e263d4cea6d7a8c9eb3ad` (`permission_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- 正在导出表  atomgo.role_permissions_permission 的数据：~11 rows (大约)
INSERT INTO `role_permissions_permission` (`role_id`, `permission_id`) VALUES
	(2, 1),
	(2, 2),
	(2, 3),
	(2, 4),
	(2, 5),
	(2, 9),
	(2, 10),
	(2, 11),
	(2, 12),
	(2, 14),
	(2, 15);

-- 导出  表 atomgo.user 结构
CREATE TABLE IF NOT EXISTS `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `enable` tinyint NOT NULL DEFAULT '1',
  `create_time` datetime(6) NOT NULL DEFAULT (now(6)),
  `update_time` datetime(6) NOT NULL DEFAULT (now(6)) ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `IDX_78a916df40e02a9deb1c4b75ed` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- 正在导出表  atomgo.user 的数据：~1 rows (大约)
INSERT INTO `user` (`id`, `username`, `password`, `enable`, `create_time`, `update_time`) VALUES
	(1, 'admin', '$2a$10$FsAafxTTVVGXfIkJqvaiV.1vPfq4V9HW298McPldJgO829PR52a56', 1, '2023-11-18 16:18:59.150632', '2024-12-06 10:08:43.325324');

-- 导出  表 atomgo.user_roles_role 结构
CREATE TABLE IF NOT EXISTS `user_roles_role` (
  `user_id` int NOT NULL,
  `role_id` int NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`) USING BTREE,
  KEY `IDX_5f9286e6c25594c6b88c108db7` (`user_id`) USING BTREE,
  KEY `IDX_4be2f7adf862634f5f803d246b` (`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- 正在导出表  atomgo.user_roles_role 的数据：~2 rows (大约)
INSERT INTO `user_roles_role` (`user_id`, `role_id`) VALUES
	(1, 1),
	(1, 2);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;

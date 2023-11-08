-- --------------------------------------------------------
-- 主机:                           172.17.0.2
-- 服务器版本:                        8.0.31 - MySQL Community Server - GPL
-- 服务器操作系统:                      Linux
-- HeidiSQL 版本:                  12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 fiber-arco-admin 的数据库结构
CREATE DATABASE IF NOT EXISTS `fiber-arco-admin` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `fiber-arco-admin`;

-- 导出  表 fiber-arco-admin.system_config 结构
CREATE TABLE IF NOT EXISTS `system_config` (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `status` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'Enable',
  `sort` int NOT NULL DEFAULT '1000',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `updated_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `updated_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统_配置';

-- 正在导出表  fiber-arco-admin.system_config 的数据：~0 rows (大约)
DELETE FROM `system_config`;

-- 导出  表 fiber-arco-admin.system_dictionary 结构
CREATE TABLE IF NOT EXISTS `system_dictionary` (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `code` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `status` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'Enable',
  `sort` int NOT NULL DEFAULT '1000',
  `remark` varchar(191) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `updated_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  `updated_by` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统_字典';

-- 正在导出表  fiber-arco-admin.system_dictionary 的数据：~6 rows (大约)
DELETE FROM `system_dictionary`;
INSERT INTO `system_dictionary` (`id`, `name`, `code`, `status`, `sort`, `remark`, `created_at`, `updated_at`, `created_by`, `updated_by`) VALUES
	('ceakmm1h3kc1po7m334g', '系统_状态', 'system_status', 'Enable', 99, '保留', 1670728536, 1672387868, 'root', 'root'),
	('cect5ehh3kc09e0nsfmg', '员工_性别', 'staff_gender', 'Enable', 1000, '保留', 1671025338, 1672387873, 'root', 'root'),
	('cectfsph3kc09e0nsfn0', '员工_工作状态', 'staff_work_status', 'Enable', 1000, '保留', 1671026675, 1672387883, 'root', 'root'),
	('cen9puph3kc1sa2og61g', '菜单_类型', 'menu_type', 'Enable', 1000, '保留', 1672387835, 1672387877, 'root', 'root'),
	('cfucle9h3kc1071tiijg', '权限方法', 'menu_method', 'Enable', 1000, '保留', 1677511353, 1679559239, 'root', 'root');

-- 导出  表 fiber-arco-admin.system_dictionary_item 结构
CREATE TABLE IF NOT EXISTS `system_dictionary_item` (
  `id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `dictionary_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `label` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `value` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `color` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `status` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT 'Enable',
  `sort` int NOT NULL DEFAULT '1000',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `updated_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  `updated_by` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `dictionary_id_label` (`dictionary_id`,`label`),
  UNIQUE KEY `dictionary_id_value` (`dictionary_id`,`value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统_字典选项';

-- 正在导出表  fiber-arco-admin.system_dictionary_item 的数据：~18 rows (大约)
DELETE FROM `system_dictionary_item`;
INSERT INTO `system_dictionary_item` (`id`, `dictionary_id`, `label`, `value`, `color`, `status`, `sort`, `remark`, `created_at`, `updated_at`, `created_by`, `updated_by`) VALUES
	('ceaknvph3kc1po7m3350', 'ceakmm1h3kc1po7m334g', '正常', 'Enable', 'arcoblue', 'Enable', 99, '', 1670728703, 1671032020, 'root', 'root'),
	('ceako39h3kc1po7m335g', 'ceakmm1h3kc1po7m334g', '停用', 'Disable', 'red', 'Enable', 100, '', 1670728717, 1671032026, 'root', 'root'),
	('ced0jfph3kc07n1n6200', 'cect5ehh3kc09e0nsfmg', '女性', 'Female', 'orangered', 'Enable', 1000, '', 1671039423, 1671039423, 'root', 'root'),
	('ced0jjhh3kc07n1n620g', 'cect5ehh3kc09e0nsfmg', '男性', 'Male', 'arcoblue', 'Enable', 99, '', 1671039438, 1671039469, 'root', 'root'),
	('ced0jphh3kc07n1n6210', 'cect5ehh3kc09e0nsfmg', '未知', 'Unknown', 'purple', 'Enable', 1000, '', 1671039462, 1671039462, 'root', 'root'),
	('ced0kd9h3kc07n1n621g', 'cectfsph3kc09e0nsfn0', '在职', 'Work', 'arcoblue', 'Enable', 88, '', 1671039541, 1671039610, 'root', 'root'),
	('ced0kgph3kc07n1n6220', 'cectfsph3kc09e0nsfn0', '离职', 'Leave', 'red', 'Enable', 99, '', 1671039555, 1671039604, 'root', 'root'),
	('ced0krph3kc07n1n622g', 'cectfsph3kc09e0nsfn0', '停薪留职', 'Leave Without Pay', 'green', 'Enable', 1000, '', 1671039599, 1671039599, 'root', 'root'),
	('ced0lq1h3kc07n1n6230', 'cectfsph3kc09e0nsfn0', '退休', 'Retire', 'purple', 'Enable', 1000, '', 1671039720, 1671039729, 'root', 'root'),
	('cehgjehh3kc0f47dn6o0', 'cectfsph3kc09e0nsfn0', '实习生', 'Trainee', 'gold', 'Enable', 1000, '', 1671629242, 1671629242, 'root', 'root'),
	('cen9rj1h3kc1sa2og620', 'cen9puph3kc1sa2og61g', '目录', 'Catalog', 'gold', 'Enable', 1, '', 1672388044, 1672389873, 'root', 'root'),
	('cen9rn1h3kc1sa2og62g', 'cen9puph3kc1sa2og61g', '菜单', 'Menu', 'arcoblue', 'Enable', 10, '', 1672388060, 1672389869, 'root', 'root'),
	('cen9ruph3kc1sa2og630', 'cen9puph3kc1sa2og61g', '按钮', 'Button', 'lime', 'Enable', 1000, '', 1672388091, 1672388091, 'root', 'root'),
	('cfucll9h3kc1071tiik0', 'cfucle9h3kc1071tiijg', 'GET', 'GET', 'red', 'Enable', 10, '', 1677511381, 1677942762, 'root', 'root'),
	('cfuclp1h3kc1071tiikg', 'cfucle9h3kc1071tiijg', 'POST', 'POST', 'gold', 'Enable', 20, '', 1677511396, 1677942767, 'root', 'root'),
	('cfuclrhh3kc1071tiil0', 'cfucle9h3kc1071tiijg', 'PUT', 'PUT', 'orange', 'Enable', 30, '', 1677511406, 1677942774, 'root', 'root'),
	('cfuclvph3kc1071tiilg', 'cfucle9h3kc1071tiijg', 'DELETE', 'DELETE', 'lime', 'Enable', 1000, '', 1677511423, 1677511423, 'root', 'root'),
	('cg1lvohh3kc3452sk54g', 'cfucle9h3kc1071tiijg', 'PATCH', 'PATCH', 'blue', 'Enable', 40, '', 1677942754, 1677942779, 'root', 'root');

-- 导出  表 fiber-arco-admin.system_log 结构
CREATE TABLE IF NOT EXISTS `system_log` (
  `id` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `username` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `operation` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `method` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `params` varchar(191) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `time` bigint NOT NULL DEFAULT '0',
  `ip` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统_日志';

-- 正在导出表  fiber-arco-admin.system_log 的数据：~0 rows (大约)
DELETE FROM `system_log`;

-- 导出  表 fiber-arco-admin.system_login_log 结构
CREATE TABLE IF NOT EXISTS `system_login_log` (
  `id` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `username` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `status` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `ip` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统_登录日志\r\n';

-- 正在导出表  fiber-arco-admin.system_login_log 的数据：~0 rows (大约)
DELETE FROM `system_login_log`;

-- 导出  表 fiber-arco-admin.system_menu 结构
CREATE TABLE IF NOT EXISTS `system_menu` (
  `id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `parent_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `parent_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `icon` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `permission` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'Menu',
  `method` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `component` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `link` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `visible` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `redirect` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'Enable',
  `sort` int NOT NULL DEFAULT '1000',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `updated_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  `updated_by` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统_权限';

-- 正在导出表  fiber-arco-admin.system_menu 的数据：~44 rows (大约)
DELETE FROM `system_menu`;
INSERT INTO `system_menu` (`id`, `name`, `parent_id`, `parent_ids`, `icon`, `path`, `permission`, `type`, `method`, `component`, `link`, `visible`, `redirect`, `status`, `sort`, `remark`, `created_at`, `updated_at`, `created_by`, `updated_by`) VALUES
	('cenc851h3kc3195ujung', '系统管理', 'Root', '', 'setting', '/system', '', 'Catalog', '', '', '', '', '', 'Enable', 11, '', 1672397844, 1677686124, 'root', 'root'),
	('cencdp9h3kc3195ujuog', '菜单管理', 'cenc851h3kc3195ujung', '', '', '/system/menus', '', 'Menu', '', 'system/menu', '', '', '', 'Enable', 1000, '', 1672398565, 1672402629, 'root', 'root'),
	('cencmdph3kc2sv29o3v0', '用户管理', 'cenc851h3kc3195ujung', '', '', '/system/staffs', '', 'Menu', '', 'system/staff', '', '', '', 'Enable', 100, '', 1672399671, 1677511068, 'root', 'root'),
	('cencmoph3kc2sv29o3vg', '角色管理', 'cenc851h3kc3195ujung', '', '', '/system/roles', '', 'Menu', '', 'system/role', '', '', '', 'Enable', 100, '', 1672399715, 1672402637, 'root', 'root'),
	('cencn6hh3kc2sv29o400', '部门管理', 'cenc851h3kc3195ujung', '', '', '/system/organizations', '', 'Menu', '', 'system/organization', '', '', '', 'Enable', 1000, '', 1672399770, 1672402678, 'root', 'root'),
	('cencnd9h3kc2sv29o40g', '岗位管理', 'cenc851h3kc3195ujung', '', '', '/system/positions', '', 'Menu', '', 'system/position', '', '', '', 'Enable', 1000, '', 1672399797, 1672402830, 'root', 'root'),
	('cencnpph3kc2sv29o410', '字典管理', 'cenc851h3kc3195ujung', '', '', '/system/dictionaries', '', 'Menu', '', 'system/dictionary', '', '', '', 'Enable', 1000, '', 1672399847, 1672402958, 'root', 'root'),
	('cend531h3kc2sv29o440', '查看用户', 'cencmdph3kc2sv29o3v0', '', '', '/system/staffs', '', 'Button', 'GET', '', '', '', '', 'Enable', 1000, '', 1672401548, 1677511546, 'root', 'root'),
	('cend5v9h3kc2sv29o44g', '新建用户', 'cencmdph3kc2sv29o3v0', '', '', '/system/staffs', '', 'Button', 'POST', '', '', '', '', 'Enable', 1000, '', 1672401661, 1677511557, 'root', 'root'),
	('cend65ph3kc2sv29o450', '编辑用户', 'cencmdph3kc2sv29o3v0', '', '', '/system/staffs', '', 'Button', 'PUT', '', '', '', '', 'Enable', 1000, '', 1672401687, 1677511563, 'root', 'root'),
	('cend6a9h3kc2sv29o45g', '删除用户', 'cencmdph3kc2sv29o3v0', '', '', '/system/staffs', '', 'Button', 'DELETE', '', '', '', '', 'Enable', 1000, '', 1672401705, 1677511569, 'root', 'root'),
	('cend759h3kc2sv29o46g', '新建角色', 'cencmoph3kc2sv29o3vg', '', '', '/system/roles', '', 'Button', 'POST', '', '', '', '', 'Enable', 1000, '', 1672401813, 1677511593, 'root', 'root'),
	('cend7a1h3kc2sv29o470', '编辑角色', 'cencmoph3kc2sv29o3vg', '', '', '/system/roles', '', 'Button', 'PUT', '', '', '', '', 'Enable', 1000, '', 1672401832, 1677511611, 'root', 'root'),
	('cend7f9h3kc2sv29o47g', '查看角色', 'cencmoph3kc2sv29o3vg', '', '', '/system/roles', '', 'Button', 'GET', '', '', '', '', 'Enable', 100, '', 1672401853, 1677511598, 'root', 'root'),
	('cend7lhh3kc2sv29o480', '删除角色', 'cencmoph3kc2sv29o3vg', '', '', '/system/roles', '', 'Button', 'DELETE', '', '', '', '', 'Enable', 1000, '', 1672401878, 1677511622, 'root', 'root'),
	('cend97ph3kc2sv29o48g', '查看菜单树', 'cencdp9h3kc3195ujuog', '', '', '/system/menus/tree', '', 'Button', 'GET', '', '', '', '', 'Enable', 1000, '', 1672402079, 1672402079, 'root', 'root'),
	('cend9fph3kc2sv29o490', '新建菜单', 'cencdp9h3kc3195ujuog', '', '', '/system/menus', '', 'Button', 'POST', '', '', '', '', 'Enable', 1000, '', 1672402111, 1672402149, 'root', 'root'),
	('cendd71h3kc2sv29o49g', '编辑菜单', 'cencdp9h3kc3195ujuog', '', '', '/system/menus', '', 'Button', 'PUT', '', '', '', '', 'Enable', 1000, '', 1672402588, 1672402588, 'root', 'root'),
	('cendde9h3kc2sv29o4a0', '删除菜单', 'cencdp9h3kc3195ujuog', '', '', '/system/menus', '', 'Button', 'DELETE', '', '', '', '', 'Enable', 1000, '', 1672402617, 1672402617, 'root', 'root'),
	('cenddq9h3kc2sv29o4ag', '查看部门树', 'cencn6hh3kc2sv29o400', '', '', '/system/organizations/tree', '', 'Button', 'GET', '', '', '', '', 'Enable', 1000, '', 1672402665, 1672402665, 'root', 'root'),
	('cende2hh3kc2sv29o4b0', '新建部门', 'cencn6hh3kc2sv29o400', '', '', '/system/organizations', '', 'Button', 'POST', '', '', '', '', 'Enable', 1000, '', 1672402698, 1672402698, 'root', 'root'),
	('cende79h3kc2sv29o4bg', '编辑部门', 'cencn6hh3kc2sv29o400', '', '', '/system/organizations', '', 'Button', 'PUT', '', '', '', '', 'Enable', 1000, '', 1672402717, 1672402717, 'root', 'root'),
	('cendebhh3kc2sv29o4c0', '删除部门', 'cencn6hh3kc2sv29o400', '', '', '/system/organizations', '', 'Button', 'DELETE', '', '', '', '', 'Enable', 1000, '', 1672402734, 1672402734, 'root', 'root'),
	('cendeg9h3kc2sv29o4cg', '查看岗位', 'cencnd9h3kc2sv29o40g', '', '', '/system/positions', '', 'Button', 'GET', '', '', '', '', 'Enable', 1000, '', 1672402753, 1672402753, 'root', 'root'),
	('cendem1h3kc2sv29o4d0', '新建岗位', 'cencnd9h3kc2sv29o40g', '', '', '/system/positions', '', 'Button', 'POST', '', '', '', '', 'Enable', 1000, '', 1672402776, 1672402776, 'root', 'root'),
	('cendes1h3kc2sv29o4dg', '编辑岗位', 'cencnd9h3kc2sv29o40g', '', '', '/system/positions', '', 'Button', 'PUT', '', '', '', '', 'Enable', 1000, '', 1672402800, 1672402800, 'root', 'root'),
	('cendf1hh3kc2sv29o4e0', '删除岗位', 'cencnd9h3kc2sv29o40g', '', '', '/system/positions', '', 'Button', 'DELETE', '', '', '', '', 'Enable', 1000, '', 1672402822, 1672402822, 'root', 'root'),
	('cendfhhh3kc2sv29o4eg', '查看字典', 'cencnpph3kc2sv29o410', '', '', '/system/dictionaries', '', 'Button', 'GET', '', '', '', '', 'Enable', 1000, '', 1672402886, 1677570866, 'root', 'root'),
	('cendfnhh3kc2sv29o4f0', '新建字典', 'cencnpph3kc2sv29o410', '', '', '/system/dictionaries', '', 'Button', 'POST', '', '', '', '', 'Enable', 1000, '', 1672402910, 1677570881, 'root', 'root'),
	('cendft1h3kc2sv29o4fg', '编辑字典', 'cencnpph3kc2sv29o410', '', '', '/system/dictionaries', '', 'Button', 'PUT', '', '', '', '', 'Enable', 1000, '', 1672402932, 1677570886, 'root', 'root'),
	('cendg11h3kc2sv29o4g0', '删除字典', 'cencnpph3kc2sv29o410', '', '', '/system/dictionaries', '', 'Button', 'DELETE', '', '', '', '', 'Enable', 1000, '', 1672402948, 1677570891, 'root', 'root'),
	('cendh2ph3kc2sv29o4gg', '监控服务', 'Root', '', 'search', '/monitor', '', 'Catalog', '', '', '', '', '', 'Enable', 1000, '', 1672403083, 1677687767, 'root', 'root'),
	('cendhe9h3kc2sv29o4h0', '资源监控', 'cendh2ph3kc2sv29o4gg', '', '', 'server', '', 'Menu', '', '', '', '', '', 'Enable', 1000, '', 1672403129, 1672403129, 'root', 'root'),
	('cfur5r9h3kc2945kh7j0', '选项查看', 'cencnpph3kc2sv29o410', '', '', '/system/dictionary/items', '', 'Button', 'GET', '', '', '', '', 'Enable', 1000, '', 1677570797, 1677661241, 'root', 'root'),
	('cfur6rhh3kc2945kh7jg', '选项添加', 'cencnpph3kc2sv29o410', '', '', '/system/dictionary/items', '', 'Button', 'POST', '', '', '', '', 'Enable', 1000, '', 1677570926, 1677661246, 'root', 'root'),
	('cfur741h3kc2945kh7k0', '选项修改', 'cencnpph3kc2sv29o410', '', '', '/system/dictionary/items', '', 'Button', 'PUT', '', '', '', '', 'Enable', 1000, '', 1677570960, 1677661251, 'root', 'root'),
	('cfur77ph3kc2945kh7kg', '选项删除', 'cencnpph3kc2sv29o410', '', '', '/system/dictionary/items', '', 'Button', 'DELETE', '', '', '', '', 'Enable', 1000, '', 1677570975, 1677661254, 'root', 'root'),
	('cfvh8rhh3kc2n649idu0', '查看字典选项', 'cencnpph3kc2sv29o410', '', '', '/system/dictionaries/items', '', 'Button', 'GET', '', '', '', '', 'Enable', 10, '', 1677661294, 1677661315, 'root', 'root'),
	('cfvhb99h3kc3se4ve46g', '展示菜单', 'cencdp9h3kc3195ujuog', '', '', '/system/menus/tree', '', 'Button', 'GET', '', '', '', '', 'Enable', 10, '', 1677661605, 1677661615, 'root', 'root'),
	('cfvna6ph3kc3rr5a59kg', '控制面板', 'Root', '', 'dashboard', '/dashboard', '', 'Catalog', '', '', '', '', '', 'Enable', 10, '', 1677686043, 1680236530, 'root', 'root'),
	('cfvnagph3kc3rr5a59l0', '工作台', 'cfvna6ph3kc3rr5a59kg', '', '', '/dashboard/workplace', '', 'Menu', '', 'dashboard/workplace', '', '', '', 'Enable', 1000, '', 1677686083, 1679631094, 'root', 'root'),
	('cfvnanph3kc3rr5a59lg', '监控台', 'cfvna6ph3kc3rr5a59kg', '', '', '/dashboard/monitor', '', 'Menu', '', 'dashboard/monitor', '', '', '', 'Enable', 100, '', 1677686111, 1679631088, 'root', 'root'),
	('cgmpdh9h3kc0bp0acedg', '分配角色', 'cencmdph3kc2sv29o3v0', '', '', '/system/staffs/roles', '', 'Button', 'POST', '', '', '', '', 'Enable', 1000, '', 1680709317, 1680709317, 'root', 'root'),
	('cgmpoq9h3kc0bp0acee0', '分配权限', 'cencmoph3kc2sv29o3vg', '', '', '/system/roles/menus', '', 'Button', 'POST', '', '', '', '', 'Enable', 1000, '', 1680710761, 1680710761, 'root', 'root');

-- 导出  表 fiber-arco-admin.system_organization 结构
CREATE TABLE IF NOT EXISTS `system_organization` (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `code` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `parent_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `parent_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'Enable',
  `sort` int NOT NULL DEFAULT '100',
  `remark` varchar(191) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `updated_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  `updated_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统-组织';

-- 正在导出表  fiber-arco-admin.system_organization 的数据：~8 rows (大约)
DELETE FROM `system_organization`;
INSERT INTO `system_organization` (`id`, `name`, `code`, `parent_id`, `parent_ids`, `status`, `sort`, `remark`, `created_at`, `updated_at`, `created_by`, `updated_by`) VALUES
	('cehbbuhh3kc25r4sva20', '顺阳集团', 'SooYang Group', 'Root', 'Root', 'Enable', 1000, '', 1671607802, 1671607802, 'root', 'root'),
	('cehbc79h3kc25r4sva2g', '顺阳汽车', 'SooYang Auto', 'cehbbuhh3kc25r4sva20', 'Root,cehbbuhh3kc25r4sva20', 'Enable', 1000, '', 1671607837, 1671607837, 'root', 'root'),
	('cehbcn1h3kc25r4sva30', '顺阳金融', 'SooYang Finance', 'cehbbuhh3kc25r4sva20', 'Root,cehbbuhh3kc25r4sva20', 'Enable', 1000, '', 1671607900, 1671607900, 'root', 'root'),
	('cehbd81h3kc25r4sva3g', '顺阳物流', 'SooYang Logistics', 'cehbbuhh3kc25r4sva20', 'Root,cehbbuhh3kc25r4sva20', 'Enable', 1000, '', 1671607968, 1671607968, 'root', 'root'),
	('cehbdj9h3kc25r4sva40', '顺阳不动产', 'SooYang Real Estate', 'cehbbuhh3kc25r4sva20', 'Root,cehbbuhh3kc25r4sva20', 'Enable', 1000, '', 1671608013, 1671608013, 'root', 'root'),
	('cehbe2ph3kc25r4sva4g', '奇迹投资有限公司', 'Miracle Investments Limited', 'Root', 'Root', 'Enable', 1001, '', 1671608075, 1671612183, 'root', 'root'),
	('cehcb71h3kc1re45ldo0', '大营集团', 'DaYing Group', 'Root', 'Root', 'Enable', 1000, '', 1671611804, 1671611804, 'root', 'root'),
	('cehfao9h3kc44u7aj730', '大营汽车', 'DaYing Auto', 'cehcb71h3kc1re45ldo0', 'Root,cehcb71h3kc1re45ldo0', 'Enable', 1000, '', 1671624033, 1671624033, 'root', 'root');

-- 导出  表 fiber-arco-admin.system_position 结构
CREATE TABLE IF NOT EXISTS `system_position` (
  `id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `code` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `status` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT 'Enable',
  `sort` int NOT NULL DEFAULT '1000',
  `remark` varchar(191) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `updated_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `updated_by` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统_职位';

-- 正在导出表  fiber-arco-admin.system_position 的数据：~7 rows (大约)
DELETE FROM `system_position`;
INSERT INTO `system_position` (`id`, `name`, `code`, `status`, `sort`, `remark`, `created_at`, `updated_at`, `created_by`, `updated_by`) VALUES
	('cecut1ph3kc3ll6eqidg', '首席执行官', 'CEO', 'Enable', 11, '', 1671032455, 1671032709, 'root', 'root'),
	('cecutdhh3kc3ll6eqie0', '首席财务官', 'CFO', 'Enable', 22, '', 1671032502, 1671032713, 'root', 'root'),
	('cecutqph3kc3ll6eqieg', '人力资源经理', 'HR Manager', 'Enable', 1000, '', 1671032555, 1671032555, 'root', 'root'),
	('cecuu3ph3kc3ll6eqif0', '财务经理', 'Financial Manager', 'Enable', 1000, '', 1671032591, 1671032591, 'root', 'root'),
	('cecuuu9h3kc3ll6eqifg', '首席技术官', 'CTO', 'Enable', 33, '', 1671032697, 1671032716, 'root', 'root'),
	('cecuvg9h3kc3ll6eqig0', '首席运营官', 'COO', 'Enable', 44, '', 1671032769, 1671032769, 'root', 'root'),
	('ced0m81h3kc07n1n623g', '实习生', 'Trainee', 'Enable', 1000, '', 1671039776, 1671039776, 'root', 'root');

-- 导出  表 fiber-arco-admin.system_role 结构
CREATE TABLE IF NOT EXISTS `system_role` (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `code` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'Enable',
  `sort` int NOT NULL DEFAULT '1000',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `updated_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `updated_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统_角色';

-- 正在导出表  fiber-arco-admin.system_role 的数据：~5 rows (大约)
DELETE FROM `system_role`;
INSERT INTO `system_role` (`id`, `name`, `code`, `status`, `sort`, `remark`, `created_at`, `updated_at`, `created_by`, `updated_by`) VALUES
	('cdv3kghh3kc2ki34rjsg', '系统管理员', 'system_admin', 'Enable', 100, ' ', 1669216834, 1671351560, 'root', 'root'),
	('cefcqhph3kc3dt2urv6g', '普通用户', 'system_normal', 'Enable', 1000, '', 1671351623, 1671351815, 'root', 'root'),
	('cefcqn1h3kc3dt2urv70', '超级管理员', 'super_admin', 'Enable', 10, '', 1671351644, 1671351644, 'root', 'root'),
	('cefcr51h3kc3dt2urv7g', '系统测试员', 'system_test', 'Enable', 1000, '', 1671351700, 1671351700, 'root', 'root'),
	('cfublr1h3kc1nf41mee0', '空白角色', 'system_blank', 'Enable', 1000, '无权限', 1677507308, 1679554775, 'root', 'root');

-- 导出  表 fiber-arco-admin.system_role_menu 结构
CREATE TABLE IF NOT EXISTS `system_role_menu` (
  `role_id` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `menu_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`role_id`,`menu_id`),
  UNIQUE KEY `role_id_menu_id` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统_角色_权限_关联表';

-- 正在导出表  fiber-arco-admin.system_role_menu 的数据：~94 rows (大约)
DELETE FROM `system_role_menu`;
INSERT INTO `system_role_menu` (`role_id`, `menu_id`, `created_at`, `created_by`) VALUES
	('cdv3kghh3kc2ki34rjsg', 'cenc851h3kc3195ujung', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cencdp9h3kc3195ujuog', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cencmdph3kc2sv29o3v0', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cencmoph3kc2sv29o3vg', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cencn6hh3kc2sv29o400', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cencnd9h3kc2sv29o40g', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cend531h3kc2sv29o440', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cend5v9h3kc2sv29o44g', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cend65ph3kc2sv29o450', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cend6a9h3kc2sv29o45g', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cend759h3kc2sv29o46g', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cend7a1h3kc2sv29o470', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cend7f9h3kc2sv29o47g', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cend7lhh3kc2sv29o480', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cend97ph3kc2sv29o48g', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cend9fph3kc2sv29o490', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cendd71h3kc2sv29o49g', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cendde9h3kc2sv29o4a0', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cenddq9h3kc2sv29o4ag', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cende2hh3kc2sv29o4b0', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cende79h3kc2sv29o4bg', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cendebhh3kc2sv29o4c0', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cendeg9h3kc2sv29o4cg', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cendem1h3kc2sv29o4d0', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cendes1h3kc2sv29o4dg', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cendf1hh3kc2sv29o4e0', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cendfhhh3kc2sv29o4eg', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cfur5r9h3kc2945kh7j0', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cfvh8rhh3kc2n649idu0', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cfvhb99h3kc3se4ve46g', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cfvna6ph3kc3rr5a59kg', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cfvnagph3kc3rr5a59l0', 1677687694, 'Root'),
	('cdv3kghh3kc2ki34rjsg', 'cfvnanph3kc3rr5a59lg', 1677687694, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cend531h3kc2sv29o440', 1679591042, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cend7f9h3kc2sv29o47g', 1679591042, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cend97ph3kc2sv29o48g', 1679591042, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cenddq9h3kc2sv29o4ag', 1679591042, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cendeg9h3kc2sv29o4cg', 1679591042, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cendfhhh3kc2sv29o4eg', 1679591042, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cfur5r9h3kc2945kh7j0', 1679591042, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cfvh8rhh3kc2n649idu0', 1679591042, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cfvhb99h3kc3se4ve46g', 1679591042, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cfvna6ph3kc3rr5a59kg', 1679591042, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cfvnagph3kc3rr5a59l0', 1679591042, 'Root'),
	('cefcqhph3kc3dt2urv6g', 'cfvnanph3kc3rr5a59lg', 1679591042, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cenc851h3kc3195ujung', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cencdp9h3kc3195ujuog', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cencmdph3kc2sv29o3v0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cencmoph3kc2sv29o3vg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cencn6hh3kc2sv29o400', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cencnd9h3kc2sv29o40g', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cencnpph3kc2sv29o410', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cend531h3kc2sv29o440', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cend5v9h3kc2sv29o44g', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cend65ph3kc2sv29o450', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cend6a9h3kc2sv29o45g', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cend759h3kc2sv29o46g', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cend7a1h3kc2sv29o470', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cend7f9h3kc2sv29o47g', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cend7lhh3kc2sv29o480', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cend97ph3kc2sv29o48g', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cend9fph3kc2sv29o490', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendd71h3kc2sv29o49g', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendde9h3kc2sv29o4a0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cenddq9h3kc2sv29o4ag', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cende2hh3kc2sv29o4b0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cende79h3kc2sv29o4bg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendebhh3kc2sv29o4c0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendeg9h3kc2sv29o4cg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendem1h3kc2sv29o4d0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendes1h3kc2sv29o4dg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendf1hh3kc2sv29o4e0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendfhhh3kc2sv29o4eg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendfnhh3kc2sv29o4f0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendft1h3kc2sv29o4fg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendg11h3kc2sv29o4g0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendh2ph3kc2sv29o4gg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cendhe9h3kc2sv29o4h0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cfur5r9h3kc2945kh7j0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cfur6rhh3kc2945kh7jg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cfur741h3kc2945kh7k0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cfur77ph3kc2945kh7kg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cfvh8rhh3kc2n649idu0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cfvhb99h3kc3se4ve46g', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cfvna6ph3kc3rr5a59kg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cfvnagph3kc3rr5a59l0', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cfvnanph3kc3rr5a59lg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cgmpdh9h3kc0bp0acedg', 1680751791, 'Root'),
	('cefcqn1h3kc3dt2urv70', 'cgmpoq9h3kc0bp0acee0', 1680751791, 'Root'),
	('cefcr51h3kc3dt2urv7g', 'cfvh8rhh3kc2n649idu0', 1680751925, 'Root'),
	('cfublr1h3kc1nf41mee0', 'cend531h3kc2sv29o440', 1679627064, 'Root'),
	('cfublr1h3kc1nf41mee0', 'cfvna6ph3kc3rr5a59kg', 1679627064, 'Root'),
	('cfublr1h3kc1nf41mee0', 'cfvnagph3kc3rr5a59l0', 1679627064, 'Root'),
	('cfublr1h3kc1nf41mee0', 'cfvnanph3kc3rr5a59lg', 1679627064, 'Root');

-- 导出  表 fiber-arco-admin.system_role_organization 结构
CREATE TABLE IF NOT EXISTS `system_role_organization` (
  `role_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `organization_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`role_id`,`organization_id`),
  UNIQUE KEY `role_id_organization_id` (`role_id`,`organization_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统_角色-组织关系表';

-- 正在导出表  fiber-arco-admin.system_role_organization 的数据：~0 rows (大约)
DELETE FROM `system_role_organization`;

-- 导出  表 fiber-arco-admin.system_staff 结构
CREATE TABLE IF NOT EXISTS `system_staff` (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `username` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `password` varchar(191) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `name` varchar(32) COLLATE utf8mb4_bin DEFAULT '',
  `email` varchar(128) COLLATE utf8mb4_bin DEFAULT '',
  `mobile` varchar(32) COLLATE utf8mb4_bin DEFAULT '',
  `avatar` varchar(191) COLLATE utf8mb4_bin DEFAULT '',
  `gender` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT 'male',
  `organization_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `position_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `work_status` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `status` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT 'Enable',
  `sort` int NOT NULL DEFAULT '1000',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` bigint NOT NULL DEFAULT '0',
  `updated_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `updated_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `deleted_at` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统-员工';

-- 正在导出表  fiber-arco-admin.system_staff 的数据：~16 rows (大约)
DELETE FROM `system_staff`;
INSERT INTO `system_staff` (`id`, `username`, `password`, `name`, `email`, `mobile`, `avatar`, `gender`, `organization_id`, `position_id`, `work_status`, `status`, `sort`, `remark`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`) VALUES
	('Root', 'admin', '$2a$04$1OBhgKs/CTVc84qlq6Gp1.Q.o.DWqCRw.vQuAhNWOkOEvRns6X09i', '管理员', 'admin@qq.com', '12312312312', '', 'Male', 'cehbbuhh3kc25r4sva20', 'cecut1ph3kc3ll6eqidg', 'Work', 'Enable', 10, '', 1671631786, 1677482779, 'root', 'root', 0),
	('cehh39ph3kc3tc0sr0b0', 'chenyangji', '$2a$04$ctgU67/AnV9dHEeYCbQ49uFLXE2UzDRhs4b.QUvJ/SnyX3jeImerK', '陈养吉', 'chenyangji@sy.com', '18601234567', '', 'Male', 'cehbbuhh3kc25r4sva20', 'cecut1ph3kc3ll6eqidg', 'Work', 'Enable', 11, '', 1671631271, 1677422998, 'root', 'root', 0),
	('cehh4dph3kc0ss5e3dtg', 'chendaojun', '$2a$04$/Nea3miRQ79OnugDAtlTBui.gnW0ZvdDvkfBbSQHTmlvxTdcYctaa', '陈导俊', 'chendaojun@mr.com', '18601234568', '', 'Male', 'cehbe2ph3kc25r4sva4g', 'cecutdhh3kc3ll6eqie0', 'Work', 'Enable', 1000, '', 1671631415, 1671631415, 'root', 'root', 0),
	('cehh5b1h3kc0ss5e3du0', 'songhaoying', '$2a$04$e1ZkFKnmFDomNAu9rw8vZ.OTjPu6e0lk54ezI16m3lIseOdE631rG', '徐旻渶', 'songhaoying@soul.com', '18601234569', '', 'Female', 'cehbe2ph3kc25r4sva4g', 'ced0m81h3kc07n1n623g', 'Work', 'Enable', 1000, '', 1671631532, 1671631532, 'root', 'root', 0),
	('cehh5sph3kc0ss5e3dug', 'songrongji', '$2a$04$F2H3jEo8lY0aPTs0E7/Oee9rlfOOTheS8Ydua90WDri19QYuVPE7a', '陈荣基', 'songrongji@sy.com', '18601234570', '', 'Male', 'cehbdj9h3kc25r4sva40', 'cecuvg9h3kc3ll6eqig0', 'Work', 'Enable', 1000, '', 1671631603, 1671631603, 'root', 'root', 0),
	('cehh6b9h3kc0ss5e3dv0', 'chenxingjun', '$2a$04$6IloBdwtAWssKHKSDCNVDuft50UXKt4WsVV0IBcpzRI4uVXAJTv/.', '陈星俊', 'chenxingjun@sy.com', '18601234571', '', 'Male', 'cehbc79h3kc25r4sva2g', 'cecuu3ph3kc3ll6eqif0', 'Work', 'Enable', 1000, '', 1671631661, 1671631661, 'root', 'root', 0),
	('cehh6kph3kc0ss5e3dvg', 'chendongji', '$2a$04$w8FcvW48yP004Zpjg5Gg1eNaRQkygcZlkpV3NHgGDByF4zSetLx2C', '陈动基	', 'chendongji@sy.com', '18601234572', '', 'Male', 'cehbcn1h3kc25r4sva30', 'cecuvg9h3kc3ll6eqig0', 'Work', 'Enable', 1000, '', 1671631699, 1671631699, 'root', 'root', 0),
	('cehh741h3kc0ss5e3e0g', 'chenruijun', '$2a$04$dzimFCdaus1kJ3qhfxL/jO.rkrVK/7dvniSQF2sDL8OVggabIgYn.', '陈叡俊	', 'chenruijun@sy.com', '18601234573', '', 'Female', 'cehbd81h3kc25r4sva3g', 'cecuu3ph3kc3ll6eqif0', 'Work', 'Enable', 1000, '', 1671631760, 1671631760, 'root', 'root', 0),
	('cehh7ghh3kc0ss5e3e1g', '3222', '$2a$04$UgIbIx39ylbN6ZrEDoYZa.wZp/TPM4Y1ZpskfYqCgUHRGuM3rCDLO', '23232', '23232@qq.com', '22312312312', '', 'Male', 'cehbdj9h3kc25r4sva40', 'cecutqph3kc3ll6eqieg', 'Work', 'Enable', 1000, '', 1671631810, 1671631810, 'root', 'root', 1679505224),
	('cehh7lph3kc0ss5e3e20', 'werwers', '$2a$04$BoUzDTfuKSWtHZ4XJlEcIuYzlyqwajzxouLr6SrSwZN0x9j0rms26', '44343', '3423342@qq.comsdf', '11123332212', '', 'Male', 'cehbdj9h3kc25r4sva40', 'cecutqph3kc3ll6eqieg', 'Work', 'Enable', 1000, '', 1671631831, 1679504965, 'root', 'root', 0),
	('cehh7pph3kc0ss5e3e2g', '11111111', '$2a$04$YaCjsA1Zdntrzw1JV1rPUes8tUveNubBtpzIWAVd.ufGPByN4PSWy', '11111111', '11311@qq.com', '11111111111', '', 'Male', 'cehbdj9h3kc25r4sva40', 'cecuvg9h3kc3ll6eqig0', 'Work', 'Enable', 1000, '', 1671631847, 1679554242, 'root', 'root', 1679554253),
	('cejv2chh3kc33726f7s0', 'wushixuan', '$2a$04$51CMISRb/8g1PLAlNXKdpeci.uzyJiS1cpY366ZsnKrg4Apxze4KC', '吴世炫	', 'wushixuan@mirc.com', '18601234574', '', 'Male', 'cehbe2ph3kc25r4sva4g', 'cecutdhh3kc3ll6eqie0', 'Work', 'Enable', 1000, '', 1671950642, 1671950642, 'root', 'root', 0),
	('cejvos9h3kc33726f7sg', 'zhurongyi', '$2a$04$qLVM3.jg.0dcq69mgPyJ6Oy5tJbknyJZPq9m0pJtsykUflczvdXb.', '朱荣逸', 'zhurongyi@dying.com', '18601234575', '', 'Male', 'cehcb71h3kc1re45ldo0', 'cecut1ph3kc3ll6eqidg', 'Work', 'Enable', 1000, '', 1671953521, 1671953521, 'root', 'root', 0),
	('cftn229h3kc2is6h24fg', '23424324', '$2a$04$JufnDMnxykVCZz19pf5jr.4F5vDgdlovuDm/7h/qdB54z0xWycdCi', '234234324', '2234243@qq.com', '13883483444', '', 'Male', 'cehbc79h3kc25r4sva2g', 'cecutdhh3kc3ll6eqie0', 'Work', 'Enable', 1000, '', 1677422857, 1677422857, 'root', 'root', 0),
	('cfu9n2ph3kc1do5iq4og', '423423499', '$2a$04$dvg0JEVQovp6g.1bNliUs.8HaofILANh8DyzxTfi7ihz38TEpLLFa', '23423424', '23424@qq.com', '12312312328', '', 'Male', 'cehbbuhh3kc25r4sva20', 'cecutdhh3kc3ll6eqie0', 'Work', 'Enable', 1000, '', 1677499275, 1679497375, 'root', 'root', 0),
	('cgdhek9h3kc3o57154qg', '1231231', '$2a$04$h3USWnKDCjT8rF5/jJwAverk.EcsqrnOj2V0DyoqESfBq4UUte5/.', '123123', '12312@qq.com', '12312311111', '', 'Female', 'cehbcn1h3kc25r4sva30', 'cecutdhh3kc3ll6eqie0', 'Trainee', 'Disable', 1000, '', 1679497041, 1679498044, 'root', 'root', 0);

-- 导出  表 fiber-arco-admin.system_staff_role 结构
CREATE TABLE IF NOT EXISTS `system_staff_role` (
  `staff_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `role_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `created_at` bigint NOT NULL DEFAULT '0',
  `created_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`staff_id`,`role_id`),
  UNIQUE KEY `staff_id_role_id` (`staff_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统-员工_角色';

-- 正在导出表  fiber-arco-admin.system_staff_role 的数据：~9 rows (大约)
DELETE FROM `system_staff_role`;
INSERT INTO `system_staff_role` (`staff_id`, `role_id`, `created_at`, `created_by`) VALUES
	('Root', 'cefcqn1h3kc3dt2urv70', 1680709186, 'Root'),
	('cehh39ph3kc3tc0sr0b0', 'cdv3kghh3kc2ki34rjsg', 1677681442, 'Root'),
	('cehh4dph3kc0ss5e3dtg', 'cefcqhph3kc3dt2urv6g', 1679505127, 'Root'),
	('cehh5b1h3kc0ss5e3du0', 'cefcr51h3kc3dt2urv7g', 1677507448, 'Root'),
	('cehh5sph3kc0ss5e3dug', 'cefcqhph3kc3dt2urv6g', 1679505162, 'Root'),
	('cehh6b9h3kc0ss5e3dv0', 'cefcqhph3kc3dt2urv6g', 1679505149, 'Root'),
	('cehh6kph3kc0ss5e3dvg', 'cefcqhph3kc3dt2urv6g', 1679505109, 'Root'),
	('cehh741h3kc0ss5e3e0g', 'cefcqhph3kc3dt2urv6g', 1679505119, 'Root'),
	('cehh7ghh3kc0ss5e3e1g', 'cefcqhph3kc3dt2urv6g', 1679505166, 'Root');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;

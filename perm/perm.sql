
-- 创建role表
CREATE TABLE `tbl_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) DEFAULT 0,
  `role_name` varchar(255) NOT NULL,
  `memo` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

insert into tbl_role (role_id,role_name) values (100,'系统管理员');
insert into tbl_role (role_id,role_name) values (101,'经理');
insert into tbl_role (role_id,role_name) values (102,'员工');

select * from tbl_role;

-- 创建权限表
CREATE TABLE `tbl_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `perm_id` int(10) DEFAULT 0,
  `parent_id` int(10) DEFAULT 0,
  `name` varchar(255) DEFAULT '',
  `href` varchar(255) DEFAULT '',
  `type` tinyint(4) DEFAULT 0,
  `sort` int(10) DEFAULT 0,
  `bit_pos` int(10) DEFAULT 0,
  `bit_group` tinyint(4) DEFAULT 0,
  `permission` varchar(255) DEFAULT '',
  `memo` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

insert into tbl_permission (perm_id,`name`,bit_pos,bit_group,permission) values (1000,'查看统计权限',3,1,'00001000');
insert into tbl_permission (perm_id,`name`,bit_pos,bit_group,permission) values (1001,'查看姓名权限',2,1,'00000100');

select * from tbl_permission;

-- 创建role x 权限表
CREATE TABLE `tbl_role_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) DEFAULT 0,
  `permission_bit_group` tinyint(4) DEFAULT 0,
  `permission_bit_str` varchar(255) DEFAULT '',
  `memo` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

insert into tbl_role_permission (role_id,permission_bit_group,permission_bit_str) values (101,1,'00001100');
insert into tbl_role_permission (role_id,permission_bit_group,permission_bit_str) values (102,1,'00000100');

-- 根据角色名查看权限
SELECT
p.*
FROM
 tbl_role r
 LEFT JOIN tbl_role_permission rp ON r.role_id = rp.role_id
 LEFT JOIN tbl_permission p ON p.bit_group = rp.permission_bit_group
WHERE
CONV(rp.permission_bit_str,2,10) >> p.bit_pos & 1 = 1
and r.role_name = '经理';

SELECT
p.*
FROM
 tbl_role r
 LEFT JOIN tbl_role_permission rp ON r.role_id = rp.role_id
 LEFT JOIN tbl_permission p ON p.bit_group = rp.permission_bit_group
WHERE
CONV(rp.permission_bit_str,2,10) >> p.bit_pos & 1 = 1
and r.role_name = '员工';


-- select p.name,p.permission,p.bit_group,p.bit_pos from tbl_permission p;

-- +--------------------+------------+-----------+---------+
-- | name             | permission | bit_group | bit_pos |
-- +--------------------+------------+-----------+---------+
-- | 查看统计权限       | 00001000   |         1 |       4 |
-- | 查看姓名权限       | 00000100   |         1 |       3 |
-- +--------------------+------------+-----------+---------+


-- select rp.permission_bit_str,rp.permission_bit_group from tbl_role r 
--       LEFT JOIN tbl_role_permission rp 
--       ON r.role_id = rp.role_id where r.role_name = '经理';

-- +--------------------+----------------------+
-- | permission_bit_str | permission_bit_group |
-- +--------------------+----------------------+
-- | 00001100           |                    1 |
-- +--------------------+----------------------+


-- select p.name,p.permission from tbl_permission p 
--   inner join 
--     (select rp.permission_bit_str,rp.permission_bit_group from tbl_role r 
--       LEFT JOIN tbl_role_permission rp 
--       ON r.role_id = rp.role_id where r.role_name = '经理') as x 
--   ON CONV(x.permission_bit_str,2,10) >> p.bit_pos & 1 = 1;

-- select p.name,p.permission from tbl_permission p 
--   inner join 
--     (select rp.permission_bit_str,rp.permission_bit_group from tbl_role r 
--       LEFT JOIN tbl_role_permission rp 
--       ON r.role_id = rp.role_id where r.role_name = '员工') as x 
--   ON CONV(x.permission_bit_str,2,10) >> p.bit_pos & 1 = 1;

-- 根据权限名查角色
SELECT
r.*
FROM
tbl_permission p
LEFT JOIN tbl_role_permission rp ON p.bit_group = rp.permission_bit_group
LEFT JOIN tbl_role r ON r.role_id = rp.role_id
WHERE
CONV(rp.permission_bit_str,2,10) >> p.bit_pos & 1 = 1
and p.`name` = '查看统计权限';

SELECT
r.*
FROM
tbl_permission p
LEFT JOIN tbl_role_permission rp ON p.bit_group = rp.permission_bit_group
LEFT JOIN tbl_role r ON r.role_id = rp.role_id
WHERE
CONV(rp.permission_bit_str,2,10) >> p.bit_pos & 1 = 1
and p.`name` = '查看姓名权限';
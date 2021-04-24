-- 用户表
create table users(
    `id`               int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `created_at`       timestamp        NULL DEFAULT NULL COMMENT '元数据',
    `updated_at`       timestamp        NULL DEFAULT NULL COMMENT '元数据',
    `deleted_at`       timestamp        NULL DEFAULT NULL COMMENT '支持软删除',
    `user_name`        varchar(1024)    NOT NULL  COMMENT '用户名',
    `pass_word`        varchar(1024)    NOT NULL  COMMENT '密码',
    `emp_type`         int(10)          DEFAULT 0 COMMENT '员工类型',
    `auth_level`       int(10)          NOT NULL  COMMENT '权限级别',
    `status`           int(10)          DEFAULT 0 COMMENT '帐号状态',
    `department`       int(10)          NOT NULL  COMMENT '隶属部门',
    `is_root`          int(10)          DEFAULT 0 COMMENT '是否是root账号',
    primary key(`id`)
) Engine=InnoDB default charset=utf8;


-- 资料表
create table technique_datas(
    `id`               int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `title`            varchar(1024)    NOT NULL          COMMENT '资料名称',
    `created_at`       timestamp        NULL DEFAULT NULL COMMENT '元数据',
    `updated_at`       timestamp        NULL DEFAULT NULL COMMENT '元数据',
    `deleted_at`       timestamp        NULL DEFAULT NULL COMMENT '支持软删除',
    `add_user`         varchar(1024)    NOT NULL          COMMENT '添加人',
    `update_user`      varchar(2014)    NOT NULL          COMMENT '修后修改人',
    `type`             int(10)          NOT NULL          COMMENT '资料类型',
    `auth_level`       int(10)          NOT NULL          COMMENT '权限等级',
    `department`       int(10)          NOT NULL          COMMENT '所属部门',
    `introduction`     longtext        NOT NULL          COMMENT '资料介绍',
    `data_url`         varchar(1024)    NOT NULL          COMMENT '资料存储url',
    primary key (`id`),
    FULLTEXT(`title`, `introduction`)
)Engine=InnoDB default charset=utf8;


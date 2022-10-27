CREATE TABLE `user`
(
    `id`          INT(10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `email`       VARCHAR(255) NOT NULL UNIQUE COMMENT 'email',
    `password`    VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码',
    `desc`        VARCHAR(255) NOT NULL DEFAULT '',
    `status`      INT(2) NOT NULL DEFAULT 1,
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `admin`
(
    `id`          INT(10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_name`   VARCHAR(64)  NOT NULL UNIQUE COMMENT '用户名',
    `password`    VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码',
    `desc`        VARCHAR(255) NOT NULL DEFAULT '',
    `status`      INT(2) NOT NULL DEFAULT 1,
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员用户表';
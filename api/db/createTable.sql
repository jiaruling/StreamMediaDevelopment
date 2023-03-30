drop table if exists `user`;
create table user
(
    id  int unsigned auto_increment comment '主键id' primary key,
    username varchar(64) unique not null comment '用户名',
    pwd   varchar(128) not null comment '密码'
) comment '用户表' charset = utf8mb4;

drop table if exists `video`;
create table video
(
    id  varchar(64) not null comment '主键id' primary key,
    user_id int unsigned not null comment '用户id',
    name varchar(128) unique not null comment '名称',
    display_time varchar(128) not null comment '显示时间',
    create_time datetime comment '创建时间'
) comment '资源表' charset = utf8mb4;

drop table if exists `comments`;
create table comments
(
    id  varchar(64) not null comment '主键id' primary key,
    video_id varchar(64) not null comment '视频id',
    user_id int unsigned not null comment '用户id',
    content text comment  '评论内容',
    time datetime comment '评论时间'
) comment '评论表' charset = utf8mb4;

drop table if exists `sessions`;
create table sessions
(
    session_id varchar(64) NOT NULL PRIMARY KEY,
    TTL TINYTEXT comment '过期时间',
    username varchar(64)  comment '登录用户'
) comment 'sessions表' charset = utf8mb4;

drop table if exists `video_delete`;
create table video_delete
(
    id  varchar(64) not null comment '主键id' primary key
) comment '资源删除表' charset = utf8mb4;
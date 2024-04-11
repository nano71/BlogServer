create database if not exists blog;
use blog;
create table if not exists article
(
    id            int auto_increment
        primary key,
    title         varchar(255)                         not null comment '标题',
    content       text                                 null,
    description   text                                 not null comment 'description',
    update_time   datetime   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    create_time   datetime   default CURRENT_TIMESTAMP not null,
    tags          varchar(255)                         null,
    cover_image   varchar(255)                         null,
    read_count    int        default 0                 not null,
    comment_count int        default 0                 not null,
    markdown      text                                 not null,
    is_visible    tinyint(1) default 1                 not null
)
    comment '文章' charset = utf8mb4;


create table if not exists guestbook
(
    id          int auto_increment
        primary key,
    nickname    varchar(255)                         not null,
    face        varchar(128)                         not null,
    content     text                                 not null,
    url         varchar(255)                         null,
    create_time datetime   default CURRENT_TIMESTAMP not null,
    ip          varchar(15)                          not null comment 'ip地址',
    is_visible  tinyint(1) default 1                 not null
)
    comment '留言板' charset = utf8mb4;

create table if not exists history
(
    time    datetime default CURRENT_TIMESTAMP not null
        primary key,
    content json                               not null
)
    comment '历史记录' charset = utf8mb4;

create table if not exists tag
(
    id         int auto_increment,
    name       varchar(255)         not null
        primary key,
    content    text                 not null,
    is_visible tinyint(1) default 1 not null,
    constraint label_pk
        unique (id)
)
    comment '标签' charset = utf8mb4;

create table if not exists log
(
    id          int primary key auto_increment,
    ip          varchar(15)  not null comment '用户IP',
    create_time datetime     not null comment '时间',
    url         varchar(255) not null comment '访问网址',
    ua          varchar(255) not null comment 'UA标识',
    latency     varchar(255) not null comment '处理耗时',
    index (create_time),
    index (ip),
    index (url)
)
    comment '操作日志' charset = utf8mb4;



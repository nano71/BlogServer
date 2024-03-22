create database if not exists blog;
use blog;
create table if not exists article
(
    id            int auto_increment
        primary key,
    title         varchar(255)                       not null comment '标题',
    content       text                               null,
    description   varchar(255)                       not null comment 'description',
    update_time   datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    create_time   datetime default CURRENT_TIMESTAMP not null,
    tags          varchar(255)                       null,
    cover_image   varchar(255)                       null,
    read_count    int      default 0                 not null,
    comment_count int      default 0                 not null,
    markdown      text                               not null
)
    comment "文章表"
    charset = utf8mb4
    engine = InnoDB;

create table if not exists tag
(
    id      int auto_increment,
    name    varchar(255) not null
        primary key,
    content text         not null,
    constraint label_pk
        unique (id)
)
    comment "分类标签表"
    charset = utf8mb4
    engine = InnoDB;

create table if not exists guestbook
(
    id          int auto_increment
        primary key,
    nickname    varchar(255)                       not null,
    face        varchar(128)                       not null,
    content     text                               not null,
    url         varchar(255)                       null,
    create_time datetime default CURRENT_TIMESTAMP not null,
    ip          varchar(15)                        not null comment 'ip地址'
)
    comment '留言板表' charset = utf8mb4;

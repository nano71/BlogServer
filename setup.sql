create database blog;
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
    charset = utf8mb4
    engine = InnoDB;


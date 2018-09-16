create table users(
    username varchar(255),
    nickname varchar(255),
    male boolean,
    account_id varchar(255) not null,
    `desc` text,
    detail longtext,

    id int(10) unsigned not null auto_increment primary key,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null on update current_timestamp default current_timestamp,
    deleted_at timestamp null default null,

    unique users_account_id_unique (account_id),
    index users_username_index (username),
    index users_male_index (male)
)default charset=utf8;

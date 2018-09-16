create table user(
    username varchar(255),
    nickname varchar(255),
    male boolean,
    accountid varchar(255) not null,
    `desc` text,
    detail longtext,

    id int(10) unsigned not null auto_increment primary key,
    createdAt datetime not null default current_timestamp,
    updatedAt datetime not null on update current_timestamp default current_timestamp,

    unique user_accountid_unique (accountid),
    index user_username_index (username),
    index user_male_index (male)
)default charset=utf8;

create table user(
    username varchar(255),
    nickname varchar(255),
    male boolean,
    accountid varchar(255) not null,
    `desc` text,
    detail longtext,

    id int(20) auto_increment primary key,
    createdAt datetime,
    updatedAt datetime,

    unique user_accountid_unique (accountid),
    index user_username_index (username),
    index user_male_index (male)
)default charset=utf8;

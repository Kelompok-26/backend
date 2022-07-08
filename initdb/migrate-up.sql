create table admin
(
    id           bigint unsigned auto_increment
        primary key,
    phone_number longtext null,
    password     longtext null
);

insert admin (phone_number, password)
values ('08139820554', 'alta123');

create table user
(
    id             bigint unsigned auto_increment
        primary key,
    name           longtext    null,
    email          longtext    null,
    phone_number   longtext    null,
    password       longtext    null,
    date_of_birth  date        null,
    gender         tinytext    null,
    point          bigint      null,
    account_number longtext    null,
    created_at     datetime(3) null,
    updated_at     datetime(3) null,
    deleted_at     datetime(3) null
);

create index idx_user_deleted_at
    on user (deleted_at);

INSERT user (name, email, phone_number, password, date_of_birth, gender, point, account_number, created_at,
                          updated_at, deleted_at)
VALUES ('Galang Prang', 'galangprang@hotmail.com', '082587625551', 'galang!@#123', '2022-11-14', 'male', 200,
        '981238812', '2022-07-08 13:37:49.000', '2022-07-08 13:37:50.000', null);

create table product
(
    id            bigint unsigned auto_increment
        primary key,
    product_name  longtext    null,
    type_product  longtext    null,
    provider_name longtext    null,
    nominal       longtext    null,
    stock         bigint      null,
    created_at    datetime(3) null,
    updated_at    datetime(3) null,
    deleted_at    datetime(3) null
);

create index idx_product_deleted_at
    on product (deleted_at);

create table redeem
(
    id         bigint unsigned auto_increment
        primary key,
    user_id    bigint unsigned null,
    name       longtext        null,
    type       tinytext        null,
    nominal    longtext        null,
    point      bigint          null,
    status     mediumtext      null,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    constraint redeem_ibfk_1
        foreign key (user_id) references user (id)
);

create index user_id
    on redeem (user_id);

    create table transaction
(
    id         bigint unsigned auto_increment
        primary key,
    user_id    bigint unsigned null,
    product_id bigint unsigned null,
    name       longtext        null,
    type       tinytext        null,
    total      longtext        null,
    point      bigint          null,
    status     mediumtext      null,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    constraint transaction_ibfk_1
        foreign key (user_id) references user (id),
    constraint transaction_ibfk_2
        foreign key (product_id) references product (id)
);

create index product_id
    on transaction (product_id);

create index user_id
    on transaction (user_id);

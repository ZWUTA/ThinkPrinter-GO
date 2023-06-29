CREATE TABLE "users"
(
    id         integer
        constraint id_pk primary key,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    username   text
        constraint username_uk unique
                                 not null,
    password   text              not null,
    name       text,
    number     text,
    vip        numeric default 0 not null
)
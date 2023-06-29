CREATE TABLE "users"
(
    id         integer           not null
        primary key autoincrement,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    username   text              not null,
    password   text              not null,
    name       text,
    number     text,
    vip        numeric default 0 not null
)


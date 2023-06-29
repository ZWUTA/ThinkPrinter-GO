create table main.users
(
    uid       integer           not null
        constraint users_pk
            primary key autoincrement,
    sname     TEXT,
    snumber   TEXT,
    uname     TEXT              not null,
    upassword TEXT              not null,
    vip       integer default 0 not null
);
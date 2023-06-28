create table IF NOT EXISTS student(
sid int primary key,
sno varchar(12) unique,
sname varchar(12),
sex varchar(4),
sdept varchar(12),
csrq varchar(20)
);

create table IF NOT EXISTS course(
cid int primary key,
cno varchar(12) unique,
cname varchar(12),
cpno varchar(12),
ccredit Decimal(4,1)
);

create table IF NOT EXISTS sc(
sid int primary key,
sno varchar(20) not null,
cno varchar(20) not null,
grade Decimal(4,1) not null DEFAULT -1
);

create table IF NOT EXISTS jwbService(
rid int primary key,
sno varchar(20) not null,
rdate datetime not null,
rnote text not null,
roffice varchar(20) not null,
rstaff int not null);

create table IF NOT EXISTS users(
uid int primary key,
sno varchar(20) not null,
username  char(12),
password  varchar(255) ,
sid   char(12)  ,
sname   char(12),
balance  decimal(10,2) not null  default 1.00,
vip int not null default 1
);

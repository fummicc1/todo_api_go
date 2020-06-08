#!/bin/sh

SQL="
create database if not exists sample_todo_db;
create table todos (
    id varchar(100) primary key,
    title varchar(100) not null,
    content varchar(200),
    due datetime
);
"
mysql.server stop
mysql.server start
mysql -e "${SQL}"
go run main.go
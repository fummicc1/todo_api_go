#!/bin/sh

SQL="
use sample_todo_db;
drop table if exists todos;
create table todos (
    id int not null auto_increment primary key,
    title varchar(20) not null,
    content varchar(200),
    due datetime
);
"
mysql.server stop
mysql.server start
mysql -e "${SQL}"
create schema if not exists users;

create table if not exists users.user (
  user_id serial primary key,
  name varchar(255),
  age integer
);


insert into users.user (name, age) values ('jane doe', 10);

insert into users.user (name, age) values ('john doe', 20);

insert into users.user (name, age) values ('bob john', 30);

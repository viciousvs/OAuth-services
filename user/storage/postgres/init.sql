create extension if not exists "uuid-ossp";
create table if not exists users(
    uuid uuid primary key ,
    login varchar(255) not null unique ,
    password_hash varchar(255) not null check ( length(password_hash) > 8 )
);
insert into users(uuid, login, password_hash)
values
(uuid_generate_v4(), 'login 1', 'password1'),
(uuid_generate_v4(), 'login 2', 'password2'),
(uuid_generate_v4(), 'login 3', 'password3'),
(uuid_generate_v4(), 'login 4', 'password4'),
(uuid_generate_v4(), 'login 5', 'password5');
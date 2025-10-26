DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id serial primary key,
    name varchar(255) not null,
    nickname varchar(255) not null unique,
    email varchar(255) not null unique,
    password varchar(255) not null,
    createdin timestamp not null default NOW()
);
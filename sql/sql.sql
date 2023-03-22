/* USE mySQL */
CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

/* Aqui é PostgreSQL */
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE users(
    id serial primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(20) not null,
    createdOn TIMESTAMP default CURRENT_TIMESTAMP
);
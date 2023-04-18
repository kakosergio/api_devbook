/* USE mySQL */
CREATE DATABASE IF NOT EXISTS devbook;

USE devbook; /* no postgres é \c devbook */

/* Aqui é PostgreSQL */
DROP TABLE IF EXISTS publications;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;

CREATE TABLE users(
    id serial primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createdOn TIMESTAMP default CURRENT_TIMESTAMP
);

GRANT ALL ON ALL TABLES IN SCHEMA public TO golang;
GRANT ALL PRIVILEGES ON users TO golang;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO golang;

CREATE TABLE followers(
    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    follower_id int not null,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    PRIMARY KEY (user_id, follower_id)
);

CREATE TABLE publications(
    id serial primary key,
    title varchar(50) not null,
    body varchar(400) not null,

    author_id int not null,
    FOREIGN KEY (author_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    likes int default 0,
    createdOn TIMESTAMP default CURRENT_TIMESTAMP
);

GRANT ALL ON ALL TABLES IN SCHEMA public TO golang;
GRANT ALL PRIVILEGES ON publications TO golang;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO golang;
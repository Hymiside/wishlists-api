
-- +migrate Up
CREATE TABLE users (
    id TEXT PRIMARY KEY ,
    name TEXT NOT NULL,
    nickname TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL ,
    phone_number TEXT,
    image_url TEXT
);

CREATE TABLE wishes (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    price BIGINT NOT NULL,
    link TEXT NOT NULL ,
    image_url TEXT,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE favorites (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    wish_id TEXT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (wish_id) REFERENCES wishes(id) ON DELETE CASCADE
);

CREATE TABLE subscribes_users (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    user_id_sub TEXT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id_sub) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL,
    password TEXT NOT NULL
);

insert into admins (login, password) values ('admin', 'admin');

create unique index on users (id);
create unique index on users (nickname);

create index on wishes (id);
create index on wishes (user_id);


-- +migrate Down
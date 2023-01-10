
-- +migrate Up
create table users (
    id text unique primary key,
    name text not null,
    nickname text unique not null,
    email text unique not null,
    password_hash text not null,
    phone_number text default 'none',
    image_url text default 'none'
);

create table wishes (
    id text unique primary key,
    user_id text unique not null
        references users (id),
    title text not null,
    description text,
    price integer not null,
    link text not null,
    image text default 'none'
);

create table favorites (
    id serial unique primary key,
    user_id text not null
        references users (id),
    wish_id text not null
        references wishes (id)
);

create table subscribes_users (
    id serial unique not null,
    user_id text not null
        references users (id),
    user_id_sub text not null
        references users (id)
);

create unique index on users (id);
create unique index on users (nickname);
create unique index on users (email);
create unique index on users (phone_number);

create index on wishes (id);
create index on wishes (user_id);


-- +migrate Down

-- +migrate Up
create table users (
    id uuid unique primary key,
    name text not null,
    nickname text unique not null,
    email text unique not null,
    password_hash text not null,
    salt text not null,
    phone_number text,
    image_url text
);

create table wishes (
    id uuid unique primary key,
    user_id uuid unique not null
        references users (id),
    title text not null,
    description text,
    price integer not null,
    link text not null,
    image text
);

create table favorites (
    id uuid unique primary key,
    user_id uuid not null
        references users (id),
    wish_id uuid not null
        references wishes (id)
);

create table subscribes_users (
    id uuid unique not null,
    user_id uuid not null
        references users (id),
    user_id_sub uuid not null
        references users (id)
);

create unique index on users (id);
create unique index on users (nickname);
create unique index on users (email);
create unique index on users (phone_number);

create index on wishes (id);
create index on wishes (user_id);


-- +migrate Down
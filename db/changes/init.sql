CREATE SCHEMA IF NOT EXISTS idler;

CREATE TABLE "user"
(
    id            uuid    default gen_random_uuid() primary key,
    name          varchar(255)        not null,
    email         varchar(255) unique not null,
    password_hash varchar(255)        not null,
    registered_at timestamp           not null,
    visited_at    timestamp           not null,
    is_confirmed  boolean default false
);

CREATE TABLE conversation
(
    id           uuid default gen_random_uuid() primary key,
    name         varchar(255)                        not null,
    owner        uuid
        constraint fk_owner references "user" not null,
    participants jsonb
);
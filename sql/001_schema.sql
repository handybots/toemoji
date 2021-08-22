-- +goose Up

create table users (
    created_at    timestamp     not null default now(),
    id            bigint(20)    not null primary key
);

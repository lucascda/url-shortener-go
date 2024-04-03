CREATE TABLE IF NOT EXISTS users (
    id SERIAL primary key,
    name TEXT not null,
    email TEXT not null,
    password TEXT not null,
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
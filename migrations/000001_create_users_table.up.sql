CREATE TABLE IF NOT EXISTS users(
    id bigserial not null primary key,
    encrypted_password varchar not null,
    email varchar not null unique
);
-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL,
    email varchar UNIQUE,
    password varchar NOT NULL,
    refresh_token varchar,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    status varchar  NOT NULL DEFAULT 'active'
);

-- +goose Down
DROP TABLE users;
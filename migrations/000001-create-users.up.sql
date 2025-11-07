-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name   VARCHAR(100) NOT NULL,
    last_name    VARCHAR(100) NOT NULL,
    email        VARCHAR(150) UNIQUE NOT NULL,
    age INT NOT NULL,
    money REAL NOT NULL,
    description NOT NULL,
    password     TEXT NOT NULL,
    is_shop_owner BOOLEAN DEFAULT FALSE
);


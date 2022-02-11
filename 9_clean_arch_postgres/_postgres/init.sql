DROP TABLE IF EXISTS
    products, users
CASCADE;


DO $$ BEGIN
    CREATE TYPE role AS ENUM ('admin', 'user');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    email varchar(64) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    role role NOT NULL DEFAULT 'user'
);

CREATE TABLE IF NOT EXISTS products (
    id serial PRIMARY KEY,
    title varchar(64) NOT NULL,
    price int NOT NULL
);

INSERT INTO products (title, price) VALUES
('Ozon',	100),
('Wildberries',	200);

INSERT INTO users (email, password, role) VALUES
('user', 'f87d3cc032ff13dd4edccc776c76283fb2eaa2a4c8f87f013c637cba0703c85f', 'user'),
('admin', 'f87d3cc032ff13dd4edccc776c76283fb2eaa2a4c8f87f013c637cba0703c85f', 'admin');

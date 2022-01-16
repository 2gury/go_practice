DROP TABLE IF EXISTS
    products
CASCADE;

CREATE TABLE IF NOT EXISTS products (
    id serial PRIMARY KEY,
    title varchar(64) NOT NULL,
    price int NOT NULL
);

INSERT INTO products (title, price) VALUES
('Ozon',	100),
('Wildberries',	200);

DROP TABLE IF EXISTS
    products
CASCADE;
DROP SEQUENCE IF EXISTS items_id_seq;
CREATE SEQUENCE products_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE IF NOT EXISTS products (
    id integer DEFAULT nextval('products_id_seq') NOT NULL,
    title varchar(64) NOT NULL,
    price int NOT NULL
);

INSERT INTO products (title, price) VALUES
('Ozon',	100),
('Wildberries',	200);

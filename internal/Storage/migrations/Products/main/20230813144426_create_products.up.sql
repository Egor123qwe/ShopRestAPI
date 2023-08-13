CREATE TABLE products
(
    goods_id SERIAL PRIMARY KEY,
    type_id INTEGER NOT NULL,
    photos_id INTEGER NOT NULL,
    price FLOAT (16) NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    amount INTEGER NOT NULL
);
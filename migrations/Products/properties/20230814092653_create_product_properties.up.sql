CREATE TABLE properties
(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,

    color_id INTEGER NOT NULL,
    photos_id INTEGER NOT NULL,
    size_id INTEGER NOT NULL,
    amount INTEGER NOT NULL
);
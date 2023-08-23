CREATE TABLE types
(
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL
);
CREATE TABLE season
(
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL
);
CREATE TABLE styles
(
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL
);
CREATE TABLE country
(
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL
);

CREATE TABLE products
(
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL,
    price FLOAT (16) NOT NULL,
    description TEXT NOT NULL,
    print varchar NOT NULL,
    types_id INTEGER NOT NULL,
    FOREIGN KEY(types_id) REFERENCES types(id),
    style_id INTEGER NOT NULL,
    FOREIGN KEY(style_id) REFERENCES styles(id),
    season_id INTEGER NOT NULL,
    FOREIGN KEY(season_id) REFERENCES season(id),
    country_id INTEGER NOT NULL,
    FOREIGN KEY(country_id) REFERENCES country(id)
);


CREATE TABLE color
(
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL
);
CREATE TABLE sizes
(
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL
);


CREATE TABLE properties
(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    FOREIGN KEY(product_id) REFERENCES products(id),
    color_id INTEGER NOT NULL,
    FOREIGN KEY(color_id) REFERENCES color(id),
    photos_id INTEGER NOT NULL,
    size_id INTEGER NOT NULL,
    FOREIGN KEY(size_id) REFERENCES sizes(id),
    amount INTEGER NOT NULL
);

CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    email varchar NOT NULL unique,
    password varchar NOT NULL,
    role INTEGER NOT NULL
);

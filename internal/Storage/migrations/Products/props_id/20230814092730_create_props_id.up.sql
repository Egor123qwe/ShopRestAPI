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

CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    email varchar NOT NULL unique,
    password varchar NOT NULL,
    role INTEGER NOT NULL
);
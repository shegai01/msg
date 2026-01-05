CREATE TABLE IF not EXISTS users(
    id serial PRIMARY key,
    username VARCHAR(50) not NULL UNIQUE,
    email VARCHAR(100) not NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT TIMESTAMP
);
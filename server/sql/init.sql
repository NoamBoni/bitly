CREATE DATABASE bitly;
GRANT ALL PRIVILEGES ON DATABASE bitly TO postgres;

-- CREATE TABLE users (
--     id SERIAL PRIMARY KEY,
--     firstname VARCHAR(50),
--     lastname VARCHAR(50),
--     email VARCHAR(50),
--     password VARCHAR(100)
-- );

-- CREATE TABLE urls (
--     id SERIAL PRIMARY KEY,
--     original_url VARCHAR(100),
--     modified_url VARCHAR(100),
--     user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
-- );
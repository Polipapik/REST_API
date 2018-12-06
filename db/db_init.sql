CREATE DATABASE world;
\c world

CREATE TABLE countries(
    id          SERIAL          PRIMARY KEY,
    name        VARCHAR(256)    NOT NULL,
    population  BIGINT
);

INSERT INTO countries (name, population) VALUES
    ('Russiano', 10000),
    ('Ukraine', 7000000000),
    ('Keks', 200),
    ('JoJoFANS', 1000000);
    
-- SELECT * FROM countries;